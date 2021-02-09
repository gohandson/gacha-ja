package main

import "math/rand"

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
		return &card{rarity: raritySR, name: "ドラゴン"}
	default:
		return &card{rarity: rarityXR, name: "イフリート"}
	}
}
