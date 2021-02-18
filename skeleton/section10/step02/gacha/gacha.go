package gacha

import (
	"context"
	"runtime/trace"
)

type Play struct {
	Client  Client
	player  *Player
	resultCh chan *Card
	errCh     chan error
}

func NewPlay(p *Player) *Play {
	return &Play{
		player:  p,
		resultCh: make(chan *Card),
		// TODO: error型のチャネルをmake関数で生成する
		// errChフィールドに入れる

	}
}

func (p *Play) client() Client {
	if p.Client != nil {
		return p.Client
	}
	return defaultClient
}

func (p *Play) Result() /* TODO: 受信専用として返す */ {
	return p.resultCh
}

func (p *Play) Err() <-chan error {
	return p.errCh
}

// ※今のままだとゴールーチンリークを起こす
func (p *Play) Draw(ctx context.Context) {
	region := trace.StartRegion(ctx, "Play.Draw")
	defer region.End()

	if err := p.player.draw(1); err != nil {
		p.errCh <- err
		return
	}

	card, err := p.draw()
	if err != nil {
		// TODO: エラーをp.errChへ送信する

		return
	}
	// TODO: cardをp.resultChへ送信する

}

func (p *Play) draw() (*Card, error) {
	dist := Distribution{
		RarityN:  &Group{Ratio: 80, Cards: []*Card{&Card{ID: "n-1", Rarity: RarityN, Name: "スライム"}}},
		RarityR:  &Group{Ratio: 15, Cards: []*Card{&Card{ID: "r-1", Rarity: RarityR, Name: "オーク"}}},
		RaritySR: &Group{Ratio: 4, Cards: []*Card{&Card{ID: "sr-1", Rarity: RaritySR, Name: "ドラゴン"}}},
		RarityXR: &Group{Ratio: 1, Cards: []*Card{&Card{ID: "xr-1", Rarity: RarityXR, Name: "イフリート"}}},
	}
	return p.client().Draw(dist)
}
