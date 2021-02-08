// STEP06: レア度ごとに出る確率を変えてみよう

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

	// 0から99までの間で乱数を生成する
	num := rand.Intn(100)

	// TODO: 変数numが0〜79のときは"ノーマル"、
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
