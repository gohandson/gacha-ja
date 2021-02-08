// STEP05: レア度を作ってみよう

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

	// 0から3までの間で乱数を生成する
	num := rand.Intn(4)

	// TODO: 変数numが0のときは"ノーマル"、
	// 1のときは"R"、2のときは"SR"、
	// 3のときは"XR"と表示する
	switch num {
	case 0:
		fmt.Println("ノーマル")
	case 1:
		fmt.Println("R")
	case 2:
		fmt.Println("SR")
	case 3:
		fmt.Println("XR")
	}
}
