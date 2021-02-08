# ガチャを作ろう

## ハンズオンのやりかた

`skeleton`ディレクトリ以下に問題があり、10個のセクションに分けられています。さらにセクションはステップごとに分けられており、それぞれのステップを進めていくことで、GoでWebアプリが作れるようになっています。

各セクションの各ステップにはREADMEが用意されていますので、まずは`README`を読みます。
`README`には、そのステップを理解するための解説が書かれています。

`README`を読んだら、ソースコードを開き`TODO`コメントが書かれている箇所をコメントに従って修正して行きます。
`TODO`コメントをすべて修正し終わったら、`README`に書かれた実行例に従ってプログラムをコンパイルして実行します。

途中でわからなくなった場合は、`solution`ディレクトリ以下に解答例を用意していますので、そちらをご覧ください。

`macOS`の動作結果をもとに解説しています。
`Windows`の方は、パスの区切り文字やコマンド等を適宜読み替えてください。

## 目次

### [Section 01: 基本構文と標準パッケージ](./skeleton/section01)

学ぶこと：標準パッケージ、変数、制御構文（if、switch、for）

* STEP01: Hello, 世界
* STEP02: 変数を使ってみよう
* STEP03: ランダムな数字の表示
* STEP04: アタリとハズレを作ろう
* STEP05: レア度を作ってみよう
* STEP06: 11連ガチャを作ってみよう
* STEP07: ガチャを選べるようにしよう

### [Section 02: 型](./skeleton/section02)

学ぶこと：コンポジット型（配列、スライス、マップ）、構造体、ユーザ定義型

* STEP01: 11連ガチャの結果を記録しよう
* STEP02: 任意の回数のガチャの結果を記録しよう
* STEP03: レア度ごとに集計しよう
* STEP04: キャラクタを表す型を定義しよう（構造体）
* STEP05: キャラクタを表す型を定義しよう（ユーザ定義型）

### [Section 03: 関数とメソッド](./skeleton/section03)

学ぶこと：関数、ポインタ、メソッド

* STEP01: キャラクタ情報を入力する関数を定義しよう
* STEP02: キャラクタ情報を複数回入力する関数を定義しよう
* STEP03: ガチャチケットを導入しよう（ポインタ） 
* STEP04: ガチャチケットを導入しよう（メソッド） 

### [Section 04: パッケージとモジュール](./skeleton/section04)

学ぶこと：パッケージ、GOPATH、Go Modules

* STEP01: gachaパッケージを作ろう
* STEP02: gachaパッケージ公開しよう
* STEP03: gachaパッケージをライブラリとして使ってみよう
* STEP04: gachaパッケージをバージョンアップしよう

### Section 05: コマンドラインツール
### Section 06: エラー処理
### Section 07: HTTPクライアントとサーバ / データベース
### Section 08: Webアプリの公開
### Section 09: 抽象化とテスト
### Section 10: 並行処理

## ソースコードの取得

```
$ go env GOPATH
$ cd ↑のディレクトリに移動
$ mkdir -p src/github.com/gohandson/
$ cd src/github.com/gohandson
$ git clone https://github.com/gohandson/gacha-ja
$ cd gacha-ja
```

## ソースコードの編集

`skeleton`ディレクトリ以下のソースコードを編集する際にはセクションごとにブランチを作って作業するとよいでしょう。
以下の例は、Section 01を編集するための`fix-section01`ブランチを作成しています。

```
$ git checkout -b fix-section01
```

作業にひと区切りがついたら以下のように作業内容をコミットしてください。

```
$ git add 編集したファイル
$ git commit -m "変更の概要"
```

## ソースコードのアップデート

ハンズオン資料が更新された場合は以下のように更新してください。
なお、編集中のものがある場合はコミットしておきましょう。

```
$ git fetch -p
$ git merge origin/main 
```

アップデートの内容によっては編集中の内容とコンフリクトを起こす可能性があります。

## ライセンス

<a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">
	<img width="200" src="by-nc.eu.png">
</a>