// TODO: gachaパッケージとは別でgacha_testパッケージにする

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
		// TODO: コインが1回分に満たない場合のテスト
	}

	for name, tt := range cases {
		// TODO: ttをこのスコープで再定義しておく

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := gacha.NewPlayer(tt.tickets, tt.coin)
			got := p.DrawableNum()
			if got != tt.want {
				// TODO: 分かりやすいメッセージを出してテストを失敗させる
			}
		})
	}
}
