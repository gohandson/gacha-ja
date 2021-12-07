// STEP07: 11連ガチャを作ってみよう

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

	// TODO: 変数iを1から11まで増やしていき処理を10回繰り返す
	for i := 1; i <= 11; i += 1 {

		// 0から99までの間で乱数を生成する
		num := rand.Intn(100)

		fmt.Printf("%d回目 ", i)

		// 変数numが0〜79のときは"ノーマル"、
		// 80〜94のときは"R"、95〜98のときは"SR"、
		// それ以外のときは"XR"と表示する
		switch {
		case num < 80:
			fmt.Println("ノーマル")
		case num < 95:
			fmt.Println("R")
		case num < 99:
			fmt.Println("SR")
		default:
			fmt.Println("XR")
		}
	}
}
