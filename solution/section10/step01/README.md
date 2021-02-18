# STEP01: ボトルネックを見つけよう

## 新しく学ぶこと

* 非公開な機能のテスト

## 動かし方

```
$ go build -v -o step01
$ ./step01
```

または

```
$ go run .
```

## トレースデータの表示

プログラムを実行すると`trace.out`というトレース情報を記録したファイルができています。
次のコマンドを用いると結果をブラウザで見ることができます。

```
$ go tool trace trace.out
```

ブラウザで表示されたページにある`User-defined tasks` -> `draw handler`の`Count`列 -> `goroutine view`の順番で開くと次のような結果が表示されます。

<img src="trace.png" width="500px">
