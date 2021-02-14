// STEP02: HTTPサーバを作ってガチャの結果をブラウザで表示しよう

package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gohandson/gacha-ja/gacha"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	p := gacha.NewPlayer(10, 100)
	// ※本当はハンドラ間で競合が起きるのでNG
	play := gacha.NewPlay(p)

	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		if play.Draw() {
			fmt.Fprintln(w, play.Result())
		}

		if err := play.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "残り:", p.DrawableNum())
	})

	return http.ListenAndServe(":8080", nil)
}

func initialTickets() (int, error) {
	if flag.NArg() == 0 {
		return 0, errors.New("ガチャチケットの枚数を入力してください")
	}

	num, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		return 0, fmt.Errorf("ガチャチケット数のパース(%q):%w", flag.Arg(0), err)
	}

	return num, nil
}
