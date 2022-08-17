// STEP05: ガチャチケットを導入しよう（ポインタ）

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

type player struct {
	tickets int // ガチャ券の枚数
}

func main() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	// TODO: 10枚のガチャ券を持ったプレイヤーを作る
	p := player{tickets: 10}
	n := inputN(&p)
	results, summary := drawN(&p, n)

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *player) int {
	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		// TODO: nが0より大きくp.tickets以下である場合はbreak
		if n > 0 && n <= p.tickets {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", p.tickets)
	}
	return n
}

func drawN(p *player, n int) ([]card, map[rarity]int) {
	// TODO: p.ticketsをnだけ減らす
	p.tickets -= n
	results := make([]card, n)
	summary := make(map[rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return card{rarity: rarityN, name: "スライム"}
	case num < 95:
		return card{rarity: rarityR, name: "オーク"}
	case num < 99:
		return card{rarity: raritySR, name: "ドラゴン"}
	default:
		return card{rarity: rarityXR, name: "イフリート"}
	}
}
