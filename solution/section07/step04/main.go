// STEP04: ガチャの結果をデータベースに保存しよう

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gohandson/gacha-ja/gacha"
	"github.com/tenntenn/sqlite"
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

	db, err := sql.Open(sqlite.DriverName, "results.db")
	if err != nil {
		return fmt.Errorf("データベースのOpen:%w", err)
	}

	if err := createTable(db); err != nil {
		return err
	}

	p := gacha.NewPlayer(10, 100)
	// ※本当はハンドラ間で競合が起きるのでNG
	play := gacha.NewPlay(p)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results, err := getResults(db, 100)
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

			if err := saveResult(db, play.Result()); err != nil {
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

	return http.ListenAndServe(":8080", nil)
}

func createTable(db *sql.DB) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS results(
		id        INTEGER PRIMARY KEY,
		rarity	  TEXT NOT NULL,
		name      TEXT NOT NULL
	);`

	_, err := db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("テーブル作成:%w", err)
	}

	return nil
}

func saveResult(db *sql.DB, card *gacha.Card) error {
	const sqlStr = `INSERT INTO results(rarity, name) VALUES (?,?);`
	_, err := db.Exec(sqlStr, card.Rarity.String(), card.Name)
	if err != nil {
		return err
	}
	return nil
}

func getResults(db *sql.DB, limit int) ([]*gacha.Card, error) {
	const sqlStr = `SELECT rarity, name FROM results LIMIT ?`
	rows, err := db.Query(sqlStr, limit)
	if err != nil {
		return nil, fmt.Errorf("%qの実行:%w", sqlStr, err)
	}
	defer rows.Close()

	var results []*gacha.Card
	for rows.Next() {
		var card gacha.Card
		err := rows.Scan(&card.Rarity, &card.Name)
		if err != nil {
			return nil, fmt.Errorf("Scan:%w", err)
		}
		results = append(results, &card)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("結果の取得:%w", err)
	}

	return results, nil
}
