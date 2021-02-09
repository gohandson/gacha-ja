// STEP01: ファイルを分けよう

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	p := player{tickets: 10, coin: 100}

	n := inputN(&p)
	results, summary := drawN(&p, n)

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *player) int {

	max := p.drawableNum()
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
