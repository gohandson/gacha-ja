// STEP05: ガチャチケットを導入しよう（ポインタ）

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type rarity string

const (
	rarityN  rarity = "N"
	rarityR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

func (r rarity) String() string {
	return string(r)
}

type card struct {
	rarity rarity // レア度
	name   string // 名前
}

func (c *card) String() string {
	// TODO: レア度:名前のように文字列を作る
	// 例："SR:ドラゴン"
	return c.rarity.String() + ":" + c.name
}

type player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

// プレイヤーが行えるガチャの回数
func (p *player) drawableNum() int {
	// TODO: ガチャが行える回数を返す
	// ガチャ券は1枚で1回、コインは10枚で1回ガチャが行える
	return p.tickets + p.coin/10
}

func (p *player) draw(n int) {

	if p.drawableNum() < n {
		fmt.Println("ガチャ券またはコインが不足しています")
		return
	}

	// TODO: ガチャ券で足りる場合はガチャ券だけ使う
	// ガチャ券から優先的に使う
	if p.tickets > n {
		p.tickets -= n
		return
	}

	p.tickets = 0
	p.coin -= n * 10 // 1回あたり10枚消費する
}

func main() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	p := player{tickets: 10, coin: 100}

	n := inputN(&p)
	results, summary := drawN(&p, n)

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *player) int {

	max := p.drawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		// TODO: nが0より大きくmax以下である場合はbreak
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}

func drawN(p *player, n int) ([]*card, map[rarity]int) {
	p.draw(n)

	results := make([]*card, n)
	summary := make(map[rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &card{rarity: rarityN, name: "スライム"}
	case num < 95:
		return &card{rarity: rarityR, name: "オーク"}
	case num < 99:
		// TODO: rarityフィールドがraritySRで
		// nameフィールドが"ドラゴン"であるcard構造体のポインタを返す
		return &card{rarity: raritySR, name: "ドラゴン"}
	default:
		return &card{rarity: rarityXR, name: "イフリート"}
	}
}
