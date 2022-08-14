// STEP02: ガチャを行う回数を入力する関数を定義しよう

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

	// TODO: 関数inputNを呼び出しその結果を変数nに代入
	n := inputN()
	for i := 1; i <= n; i++ {
		draw()
	}
}

func inputN() int {
	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 {
			break
		}
		fmt.Println("もう一度入力してください")
	}
	// TODO: nを返す
	return n
}

func draw() {
	// 0から99までの間で乱数を生成する
	num := rand.Intn(100)

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
