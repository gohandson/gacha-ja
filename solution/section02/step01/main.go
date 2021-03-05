// STEP01: 11連ガチャの結果を記録しよう

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

	// TODO: 長さ11の文字列型の配列を変数resultとして定義する
	var result [11]string

	// TODO: 継続条件をiがresultの要素数より小さい場合にする
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
