# STEP03: `gacha`パッケージ公開しよう

## 新しく学ぶこと

* パッケージ（モジュール）の公開の方法

## GitHub上のパッケージの公開方法

* (GitHubのアカウントを作る)
* GitHub上でリポジトリを作る(名前はgacha)
* `git clone github.com/tenntenn/gacha`のようにリポジトリをクローンする
* クローンしたディレクトリに移動する
* `go mod init github.com/tenntenn/gacha`のように`go.mod`ファイルを生成する
* ハンズオン資料のSection 04のSTEP03にあるgachaディレクトリ以下のファイルをクローンしてきたディレクトリ以下に移す
* `git add .`でクローンしたディレクトリ以下のファイルをすべてGitの管理下に置く
* `git commit -a`ですべてコミットする
* `git push`でリモートにプッシュする
* `main.go`の`gacha`パッケージのインポートパスを自分の公開したものに変更する

* ※ `github.com/tenntenn`は自身のGitHubアカウントに読み替えてください。

## 動かし方

```
$ go build -v -o step03
$ ./step03
```

または

```
$ go run .
```

