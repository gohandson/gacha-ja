// STEP03: gachaパッケージ公開しよう

package main

import (
	"fmt"

	// TODO: インポートパスを公開したものに変更する
	"github.com/gohandson/gacha-ja/gacha"
)

func main() {
	p := gacha.NewPlayer(10, 100)

	n := inputN(p)
	results, summary := gacha.DrawN(p, n)

	fmt.Println(results)
	fmt.Println(summary)
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
