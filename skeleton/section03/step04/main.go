// STEP04: ガチャの結果とレア度毎の集計結果を返す関数を定義しよう

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

	// 関数inputNを呼び出しその結果を変数nに代入
	n := inputN()

	// TODO: 関数drawNの引数に変数nを指定して呼び出す
	// 結果を変数resultsとsummaryに代入する

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN() int {
	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 {
			break
		}
		fmt.Println("もう一度入力してください")
	}
	return n
}

func drawN(n int) /* TODO: 戻り値として結果をレア度ごとの集計を返す */ {
	results := make([]card, n)
	summary := make(map[rarity]int)
	for i := 0; i < n; i++ {
		// TODO: 関数drawが返す値をresultsのi番目に代入する

		summary[results[i].rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() /* TODO: 戻り値の型をcardにする */ {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return card{rarity: rarityN, name:"スライム"}
	case num < 95:
		return card{rarity: rarityR, name:"オーク"}
	case num < 99:
		// TODO: rarityフィールドがraritySRで
		// nameフィールドが"ドラゴン"のcard型の値を返す
	default:
		return card{rarity: rarityXR, name:"イフリート"}
	}
}
