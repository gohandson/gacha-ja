// STEP04: 非公開な機能のテストをしよう

package main

import (
	"context"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/gohandson/gacha-ja/gacha"
	"google.golang.org/api/iterator"
)

var tmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head><title>ガチャ</title></head>
	<body>
		<form action="/draw">
			<label for="num">枚数</input>
			<input type="number" name="num" min="1" value="1">
			<input type="submit" value="ガチャを引く">
		</form>
		<h1>結果一覧</h1>
		<ol>{{range .}}
		<li>{{.}}</li>
		{{end}}</ol>
	</body>
</html>`))

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "gohandson-gacha")
	if err != nil {
		return fmt.Errorf("Datastoreのクライアント作成:%w", err)
	}

	p := gacha.NewPlayer(1000, 100)
	// ※本当はハンドラ間で競合が起きるのでNG
	play := gacha.NewPlay(p)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results, err := getResults(client, 100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		num, err := strconv.Atoi(r.FormValue("num"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < num; i++ {
			if !play.Draw() {
				break
			}

			if err := saveResult(client, play.Result()); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
		}

		if err := play.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)
	return http.ListenAndServe(addr, nil)
}

func saveResult(client *datastore.Client, card *gacha.Card) error {
	ctx := context.Background()
	key := datastore.IncompleteKey("YourGitHubAccount-Results", nil)
	_, err := client.Put(ctx, key, card)
	if err != nil {
		return fmt.Errorf("結果の保存:%w", err)
	}
	return nil
}

func getResults(client *datastore.Client, limit int) ([]*gacha.Card, error) {
	results := make([]*gacha.Card, 0, limit)
	q := datastore.NewQuery("YourGitHubAccount-Results") // クエリの作成
	q = q.Limit(cap(results))          // リミット
	for it := client.Run(context.Background(), q); ; {
		var card gacha.Card
		_, err := it.Next(&card)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("結果の取得:%w", err)
		}
		results = append(results, &card)

	}

	return results, nil
}
