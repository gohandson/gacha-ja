package gacha

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	Draw(dist Distribution) (*Card, error)
}

type Distribution map[Rarity]*Group

type Group struct {
	Ratio int64
	Cards []*Card
}

var defaultClient = &client{
	baseURL: "https://gohandson-gacha.uc.r.appspot.com/",
	timeout: 5 * time.Second,
}

type client struct {
	baseURL string
	timeout time.Duration
}

var _ Client = (*client)(nil)

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

func (cli *client) sendRequest(q string) ([]byte, error) {

	// TODO: context.WithTimeout関数を使ってタイムアウトを行うコンテキストを生成する
	// ベースとなるコンテキストはcontext.Background関数で生成する
	// タイムアウトの時間はcli.timeoutフィールドから取得する
	// 第1戻り値は変数ctx、第2戻り値はcancel変数に代入する

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cli.baseURL+"?delay=on&q="+q, nil)
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

	return body, nil
}

func (cli *client) drawRarity(dist Distribution) (Rarity, error) {
	rarities := make([]string, 0, len(dist))
	for rarity, group := range dist {
		rarities = append(rarities, fmt.Sprintf("%s:%d", rarity, group.Ratio))
	}
	q := strings.Join(rarities, ",")

	body, err := cli.sendRequest(q)
	if err != nil {
		return RarityUnkown, err
	}

	return Rarity(body), nil
}

func (cli *client) drawCard(cards []*Card) (*Card, error) {
	ids := make([]string, len(cards))
	cardMap := make(map[string]*Card, len(cards))
	for i := range cards {
		ids[i] = cards[i].ID + ":1"
		cardMap[cards[i].ID] = cards[i]
	}
	q := strings.Join(ids, ",")

	body, err := cli.sendRequest(q)
	if err != nil {
		return nil, err
	}

	id := string(body)
	return cardMap[id], nil
}
