package gacha

type Play struct {
	Client  Client
	player  *Player
	results []*Card
	summary map[Rarity]int
	err     error
}

func NewPlay(p *Player) *Play {
	return &Play{
		player:  p,
		summary: make(map[Rarity]int),
	}
}

// TODO: Clientフィールドがnilでない場合はClientフィールドの値
// そうでない場合はdefaultClientを返す*Play型のメソッドclient

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
	if p.err != nil {
		return false
	}

	if err := p.player.draw(1); err != nil {
		p.err = err
		return false
	}

	card, err := p.draw()
	if err != nil {
		p.err = err
		return false
	}
	p.results = append(p.results, card)
	p.summary[card.Rarity]++

	return true
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
