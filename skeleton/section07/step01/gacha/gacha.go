package gacha

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://gohandson-gacha.uc.r.appspot.com/"

type Play struct {
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

	return p.player.DrawableNum() > 0
}

func (p *Play) draw() (*Card, error) {
	q := "スライム:80,オーク:15,ドラゴン:4,イフリート:1"
	// TODO: GETメソッドのリクエストを生成する
	// URLはbaseURLの末尾に?q=と変数qの文字列を付加したもの
	// リクエストボディはnil

	if err != nil {
		return nil, fmt.Errorf("リクエスト作成:%w", err)
	}

	// TODO: デフォルトクライアントを使ってリクエストを送る
	// レスポンスは変数respで受け取る

	if err != nil {
		return nil, fmt.Errorf("APIリクエスト:%w", err)
	}

	// TODO: ボディをクローズする

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Bodyの読み込み:%w", err)
	}

	result := string(body)
	switch result {
	case "スライム":
		return &Card{Rarity: RarityN, Name: "スライム"}, nil
	case "オーク":
		return &Card{Rarity: RarityR, Name: "オーク"}, nil
	case "ドラゴン":
		return &Card{Rarity: RaritySR, Name: "ドラゴン"}, nil
	default:
		return &Card{Rarity: RarityXR, Name: "イフリート"}, nil
	}
}
