package gacha_test

import (
	"errors"
	"testing"

	"github.com/gohandson/gacha-ja/gacha"
)

func TestPlay_Draw(t *testing.T) {
	cases := map[string]struct {
		tickets int
		cli     gacha.Client

		wantRet  bool
		wantCard string
		wantErr  bool
	}{
		"zero":  {0, cli(t, "card", false), false, "-", true},
		"one":   {1, cli(t, "card", false), true, "card", false},
		"error": {1, cli(t, "card", true), false, "-", true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := gacha.NewPlayer(tt.tickets, 0)
			play := gacha.NewPlay(p)
			play.Client = tt.cli
			got := play.Draw()
			switch {
			case !tt.wantErr && play.Err() != nil:
				t.Fatal("unexpected error", play.Err())
			case tt.wantErr && play.Err() == nil:
				t.Fatal("expected error does not occur")
			}

			if got != tt.wantRet {
				t.Errorf("want card %v but got %v", tt.wantCard, got)
			}

			gotCard := play.Result()
			if gotCard != nil && gotCard.Name != tt.wantCard {
				t.Errorf("want card %v but got %v", tt.wantCard, gotCard.Name)
			}
		})
	}

}

func cli(t *testing.T, cardName string, hasErr bool) gacha.Client {
	t.Helper()
	card := &gacha.Card{ID: "1", Rarity: gacha.RarityN, Name: cardName}
	if hasErr {
		return &mockClient{card: card, err: errors.New("error")}
	}
	return &mockClient{card: card}
}

type mockClient struct {
	card *gacha.Card
	err  error
}

var _ gacha.Client = (*mockClient)(nil)

func (cli *mockClient) Draw(_ gacha.Distribution) (*gacha.Card, error) {
	if cli.err != nil {
		return nil, cli.err
	}
	return cli.card, nil
}
