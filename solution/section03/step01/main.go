// STEP01: ガチャを行う関数を定義しよう

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type rarity string

const (
	rarityN  rarity = "ノーマル"
	rarityR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

type card struct {
	rarity rarity // レア度
	name   string // 名前
}

func main() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	draw()
}

func draw() {
	// 0から99までの間で乱数を生成する
	num := rand.Intn(100)

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
