// STEP01: ガチャチケットが足りない場合にエラーを発生させよう

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gohandson/gacha-ja/gacha"
)

var (
	flagCoin int
)

func init() {
	flag.IntVar(&flagCoin, "coin", 0, "コインの初期枚数")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	tickets, err := initialTickets()
	// TODO: エラーが発生した場合はエラーをそのまま返す

	p := gacha.NewPlayer(tickets, flagCoin)

	n := inputN(p)
	// TODO: gacha.DrawN関数を呼び出す
	// 戻り値はresults, summary, errに代入する

	if err != nil {
		return err
	}

	if err := saveResults(results); err != nil {
		return err
	}

	if err := saveSummary(summary); err != nil {
		return err
	}

	return nil
}

func initialTickets() (int, error) {
	if flag.NArg() == 0 {
		// TODO: 0とエラーを返す
		// エラーは"ガチャチケットの枚数を入力してください"
	}

	num, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return 0, err
	}

	// TODO: numとエラーがないことを表すnilを返す
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
	f, err := os.Create("results.txt")
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}

	return nil
}

func saveSummary(summary map[gacha.Rarity]int) (rerr error) {
	f, err := os.Create("summary.txt")
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}

	return nil
}
