// STEP03: レア度ごとに集計しよう

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

	// ガチャを引く回数
	var n int

	for {
		fmt.Print("何回引きますか？>")
		fmt.Scanln(&n)

		// nが0より大きい場合はforをbreakする
		if n > 0 {
			break
		}

		fmt.Println("もう一度入力してください")
	}

	// TODO: キーがstring型で値がint型のマップを定義する


	for i := 0; i < n; i++ {

		// 0から99までの間で乱数を生成する
		num := rand.Intn(100)

		// 変数numが0〜79のときは"ノーマル"、
		// 80〜94のときは"R"、95〜98のときは"SR"、
		// それ以外のときは"XR"と表示する
		switch {
		case num < 80:
			result["ノーマル"]++
		case num < 95:
			// resultのi番目に"R"を代入する
			result["R"]++
		case num < 99:
			result["SR"]++
		default:
			result["XR"]++
		}
	}

	fmt.Println(result)
}
