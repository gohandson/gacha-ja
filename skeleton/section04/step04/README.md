# STEP04: gachaパッケージにバージョンを付けよう

## 新しく学ぶこと

* モジュールへのバージョンの付け方

## バージョンの付け方

* GitHub上でリポジトリを開く（ブラウザ）
* 右端にある`Releases`をクリックする
* `Create a new release`をクリックする
* タグバージョンに`v0.0.1`と書く
* Release Titleに`v0.0.1: First release`と書く
* 詳細に`First release`と書く
* `Publish release`ボタンを押す
* ハンズオンのSection 03のSTEP04のディレクトリに戻る
* `go get github.com/tenntenn/gacha@v0.0.1`を実行し`v0.0.1`を取得する
* `go.mod`を開いて`v0.0.1`が利用されていることを確認する


* ※ `github.com/tenntenn`は自身のGitHubアカウントに読み替えてください。
* ※ リリースタイトルや詳細は実際にはそのバージョンで更新された内容などを書きます

## 動かし方

```
$ go build -v -o step04
$ ./step04
```

または

```
$ go run .
```

