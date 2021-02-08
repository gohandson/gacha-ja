// STEP03: ランダムな数字の表示

package main

import (
	// TODO: fmtパッケージをインポートする

	"math/rand"
	"time"
)

func main() {

	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	// TODO: 0から9までの間で乱数を作り変数numに代入する

	// 変数numを表示する
	fmt.Println(num)
}
