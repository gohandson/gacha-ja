// STEP02: エラーを出力しよう

package main

import (
	"fmt"
	"os"

	"github.com/gohandson/gacha-ja/gacha"
)

func main() {
	p := gacha.NewPlayer(10, 100)

	n := inputN(p)
	results, summary := gacha.DrawN(p, n)

	saveResults(results)
	saveSummary(summary)
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

func saveResults(results []*gacha.Card) {
	f, err := os.Create("results.txt")
	if err != nil {
		// TODO: 標準エラー出力にerrを出力する

		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}
}

func saveSummary(summary map[gacha.Rarity]int) {
	f, err := os.Create("summary.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// TODO: returnする
	}

	defer func() {
		if err := f.Close(); err != nil {
			// TODO: 標準エラー出力にerrを出力する
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}
