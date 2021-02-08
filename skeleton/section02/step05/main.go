// STEP05: キャラクタを表す型を定義しよう（ユーザ定義型）

package main

import "fmt"

type card struct {
	rarity string // レア度
	// TODO: 文字列型のフィールドnameを設ける
}

func main() {

	// TODO: rarityフィールドが"ノーマル"で
	// nameフィールドが"スライム"の変数slimeを定義する

	fmt.Println(slime)

	dragon := card{rarity: "SR", name: "ドラゴン"}
	fmt.Println(dragon)
}
