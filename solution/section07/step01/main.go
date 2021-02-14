// STEP01: ガチャAPIを使ってみよう

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/gohandson/gacha-ja/gacha"
)

var (
	flagCoin    int
	flagResults string
	flagSummary string
)

var (
	regexpResults = regexp.MustCompile(`^results.*\.txt$`)
	regexpSummary = regexp.MustCompile(`^summary.*\.txt$`)
)

func init() {
	flag.IntVar(&flagCoin, "coin", 0, "コインの初期枚数")
	flag.StringVar(&flagResults, "results", "results.txt", "結果ファイルの名前")
	flag.StringVar(&flagSummary, "summary", "summary.txt", "集計ファイルの名前")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	if !regexpResults.MatchString(flagResults) {
		return fmt.Errorf("結果ファイル名が不正(%s)", flagResults)
	}

	if !regexpSummary.MatchString(flagSummary) {
		return fmt.Errorf("集計ファイル名が不正(%s)", flagSummary)
	}

	tickets, err := initialTickets()
	if err != nil {
		return err
	}
	p := gacha.NewPlayer(tickets, flagCoin)
	play := gacha.NewPlay(p)

	n := inputN(p)
	for play.Draw() {
		if n <= 0 {
			break
		}
		fmt.Println(play.Result())
		n--
	}

	if err := play.Err(); err != nil {
		return fmt.Errorf("ガチャを%d回引く:%w", n, err)
	}

	if err := saveResults(play.Results()); err != nil {
		return err
	}

	if err := saveSummary(play.Summary()); err != nil {
		return err
	}

	return nil
}

func initialTickets() (int, error) {
	if flag.NArg() == 0 {
		return 0, errors.New("ガチャチケットの枚数を入力してください")
	}

	num, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return 0, fmt.Errorf("ガチャチケット数のパース(%q):%w", flag.Arg(0), err)
	}

	return num, nil
}

func inputN(p *gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}

func saveResults(results []*gacha.Card) (rerr error) {
	f, err := os.Create(flagResults)
	if err != nil {
		return fmt.Errorf("%sの作成:%w", flagResults, err)
	}

	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = fmt.Errorf("%sのクローズ:%w", flagResults, err)
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}

	return nil
}

func saveSummary(summary map[gacha.Rarity]int) (rerr error) {
	f, err := os.Create(flagSummary)
	if err != nil {
		return fmt.Errorf("%sの作成:%w", flagSummary, err)
	}

	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = fmt.Errorf("%sのクローズ:%w", flagSummary, err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}

	return nil
}
