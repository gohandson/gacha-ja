package gacha

import (
	"context"
	"runtime/trace"
)

type Play struct {
	Client   Client
	player   *Player
	resultCh chan *Card
	errCh    chan error
}

func NewPlay(p *Player) *Play {
	return &Play{
		player:   p,
		resultCh: make(chan *Card),
		errCh:    make(chan error),
	}
}

func (p *Play) client() Client {
	if p.Client != nil {
		return p.Client
	}
	return defaultClient
}

func (p *Play) Draw(ctx context.Context) (*Card, error) {
	region := trace.StartRegion(ctx, "Play.Draw")
	defer region.End()

	if err := p.player.draw(1); err != nil {
		return nil, err
	}

	card, err := p.draw()
	if err != nil {
		return nil, err
	}

	return card, nil
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
