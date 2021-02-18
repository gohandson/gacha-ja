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

func TestPlayer_draw(t *testing.T) {
	cases := map[string]struct {
		n               int
		tickets         int
		coin            int
		wantDrawableNum int
		wantErr         bool
	}{
		"zero-zero":      {1, 0, 0, 0, true},
		"plus-zero":      {1, 10, 0, 9, false},
		"plus-plus":      {1, 10, 10, 10, false},
		"zero-plus":      {1, 0, 10, 0, false},
		"zero-notenough": {1, 0, 1, 0, true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := gacha.NewPlayer(tt.tickets, tt.coin)

			err := gacha.ExportPlayerDraw(p, tt.n)
			switch {
			case !tt.wantErr && err != nil:
				t.Fatal("unexpected error", err)
			case tt.wantErr && err == nil:
				t.Fatal("expected error does not occur")
			}

			got := p.DrawableNum()
			if got != tt.wantDrawableNum {
				t.Errorf("want %d but got %d", tt.wantDrawableNum, got)
			}
		})
	}

}
