package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

type Play struct {
	player *Player
	results []*Card
	summary map[Rarity]int
	err error
}

func NewPlay(p *Player) *Play {
	return &Play{
		player: p,
		summary: make(map[Rarity]int),
	}
}

func (p *Play) Results() []*Card {
	return p.results
}

func (p *Play) Result() *Card {
	if len(p.results) == 0 {
		return nil
	}
	return p.results[len(p.results)-1]
}

func (p *Play) Summary() map[Rarity]int {
	return p.summary
}

func (p *Play) Err() error {
	return p.err
}

func (p *Play) Draw() bool {
	// TODO: エラーがすでに発生した場合はfalseを返す
	if p.err != nil {
		return false
	}
	if err := p.player.draw(1); err != nil {
		// TODO: エラーをフィールドに代入する
		p.err = err
		return false
	}

	card := p.draw()
	p.results = append(p.results, card)
	p.summary[card.Rarity]++

	return p.player.DrawableNum() > 0
}

func (p *Play) draw() *Card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &Card{Rarity: RarityN, Name: "スライム"}
	case num < 95:
		return &Card{Rarity: RarityR, Name: "オーク"}
	case num < 99:
		return &Card{Rarity: RaritySR, Name: "ドラゴン"}
	default:
		return &Card{Rarity: RarityXR, Name: "イフリート"}
	}
}
