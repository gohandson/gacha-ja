// STEP05: キャラクタを表す型を定義しよう（ユーザ定義型）

package main

import "fmt"

// TODO: string型をベースにしたrarity型を定義する
type rarity string

const (
	rarityN  rarity = "ノーマル"
	rarityR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

type card struct {
	rarity rarity // レア度
	// TODO: 文字列型のフィールドnameを設ける
	name string
}

func main() {

	// TODO: rarityフィールドがrarityNで
	// nameフィールドが"スライム"の変数slimeを定義する
	slime := card{rarity: rarityN, name: "スライム"}
	fmt.Println(slime)

	dragon := card{rarity: raritySR, name: "ドラゴン"}
	fmt.Println(dragon)
}
