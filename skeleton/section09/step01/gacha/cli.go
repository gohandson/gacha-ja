package gacha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client interface {
	// TODO: 名前がDrawで引数がDistribution、
	// 戻り値が*Cardとerrorのメソッド
}

type Distribution map[Rarity]*Group

type Group struct {
	Ratio int64
	Cards []*Card
}

var defaultClient = &client{
	baseURL: "https://gohandson-gacha.uc.r.appspot.com/",
}

type client struct {
	baseURL string
}

// TODO: *client型がClientインタフェースを実装しているかチェックする

func (cli *client) Draw(dist Distribution) (*Card, error) {
	rarity, err := cli.drawRarity(dist)
	if err != nil {
		return nil, err
	}

	card, err := cli.drawCard(dist[rarity].Cards)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (cli *client) drawRarity(dist Distribution) (Rarity, error) {
	rarities := make([]string, 0, len(dist))
	for rarity, group := range dist {
		rarities = append(rarities, fmt.Sprintf("%s:%d", rarity, group.Ratio))
	}
	q := strings.Join(rarities, ",")

	req, err := http.NewRequest(http.MethodGet, cli.baseURL+"?q="+q, nil)
	if err != nil {
		return RarityUnkown, fmt.Errorf("リクエスト作成:%w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return RarityUnkown, fmt.Errorf("APIリクエスト:%w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RarityUnkown, fmt.Errorf("Bodyの読み込み:%w", err)
	}

	return Rarity(body), nil
}

func (cli *client) drawCard(cards []*Card) (*Card, error) {
	ids := make([]string, len(cards))
	cardMap := make(map[string]*Card, len(cards))
	for i := range cards {
		ids[i] = cards[i].ID+":1"
		cardMap[cards[i].ID] = cards[i]
	}
	q := strings.Join(ids, ",")

	req, err := http.NewRequest(http.MethodGet, cli.baseURL+"?q="+q, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成:%w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("APIリクエスト:%w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Bodyの読み込み:%w", err)
	}

	id := string(body)
	return cardMap[id], nil

}
