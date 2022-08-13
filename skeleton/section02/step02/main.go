// STEP02: 任意の回数のガチャの結果を記録しよう

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

		// TODO: nが0より大きい場合はforをbreakする
		if n > 0 {
			break
		}

		fmt.Println("もう一度入力してください")
	}

	// TODO: 長さnの文字列型のスライスを変数resultとして定義する
	result := make([]string, n)

	/* TODO: 継続条件をiがresultの要素数より小さい場合にする */
	for i := 0; i < len(result); i++ {

		// 0から99までの間で乱数を生成する
		num := rand.Intn(100)

		// 変数numが0〜79のときは"ノーマル"、
		// 80〜94のときは"R"、95〜98のときは"SR"、
		// それ以外のときは"XR"と表示する
		switch {
		case num < 80:
			result[i] = "ノーマル"
		case num < 95:
			// TODO: resultのi番目に"R"を代入する
			result[i] = "R"
		case num < 99:
			result[i] = "SR"
		default:
			result[i] = "XR"
		}
	}

	fmt.Println(result)
}
