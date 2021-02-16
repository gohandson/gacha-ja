package gacha_test

import (
	"testing"

	"github.com/gohandson/gacha-ja/gacha"
)

func TestPlayer_DrawableNum(t *testing.T) {
	cases := map[string]struct {
		tickets int
		coin    int
		want    int
	}{
		"zero-zero":      {0, 0, 0},
		"plus-zero":      {10, 0, 10},
		"plus-plus":      {10, 10, 11},
		"zero-plus":      {0, 10, 1},
		"plus-notenough": {10, 1, 10},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := gacha.NewPlayer(tt.tickets, tt.coin)
			got := p.DrawableNum()
			if got != tt.want {
				t.Errorf("want %d but got %d", tt.want, got)
			}
		})
	}
}
