// STEP04: アタリとハズレを作ろう

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

	// 0または1のどちらかの値をランダムに発生させる
	num := rand.Intn(2)

	// TODO: 変数numが0のときに"アタリ"と表示し
	// そうでない場合は"ハズレ"と表示する
	if num == 0 {
		fmt.Println("アタリ")
	} else {
		fmt.Println("ハズレ")
	}
}
