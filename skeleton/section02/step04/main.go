// STEP04: キャラクタを表す型を定義しよう（構造体）

package main

import "fmt"

func main() {
	var card struct {
		rarity string // レア度
		name   string // 名前
	}

	// TODO: 変数cardのrarityフィールドに"ノーマル"と代入


	// TODO: 変数cardのnameフィールドに"スライム"と代入


	fmt.Println(card)
}
