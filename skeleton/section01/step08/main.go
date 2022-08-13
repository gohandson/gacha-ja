// STEP08: ガチャを選べるようにしよう

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

	var n int // ガチャを繰り返す回数
	fmt.Println("1: 単発ガチャ 2: 11連ガチャ")

	// TODO: LOOPというラベルのついた無限ループを作る
LOOP:
  for {
		fmt.Print(">")
		var kind int
		// TODO: 変数kindに入力した値を入れる
		fmt.Scanln(&kind)
		switch kind {
		case 1: // 単発ガチャ
			n = 1
			break LOOP
		case 2: // 11連ガチャ
			// TODO: 変数nに11を入れる
			n = 11
			break LOOP
		default:
			fmt.Println("もう一度入力してください")
		}
	}

	 /* TODO: 継続条件をiがn以下の場合とする */
	for i := 1; i <= n; i++ {

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
