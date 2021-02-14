// STEP04: 初期コインの枚数をプログラム引数で渡そう

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gohandson/gacha-ja/gacha"
)

func main() {
	tickets := initialTickets()
	p := gacha.NewPlayer(tickets, 100)

	n := inputN(p)
	results, summary := gacha.DrawN(p, n)

	saveResults(results)
	saveSummary(summary)
}

func initialTickets() int {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "ガチャチケットの枚数を入力してください")
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return num
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
		os.Exit(1)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}
