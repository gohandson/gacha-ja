# STEP01: ボトルネックを見つけよう

## 新しく学ぶこと

* トレースの仕方

## Datastoreのエミュレーターのインストールの仕方

```
$ gcloud components install cloud-datastore-emulator
```

## Datastoreのエミュレーターの起動

```
$ gcloud beta emulators datastore start --project gohandson-gacha
```

## 動かし方

 Datastoreエミュレーターに接続するには環境変数`DATASTORE_EMULATOR_HOST`に`localhost:8081`を設定する必要がある。

```
$ go build -v -o step01
$ DATASTORE_EMULATOR_HOST=localhost:8081 ./step01
```

または

```
$ DATASTORE_EMULATOR_HOST=localhost:8081 go run .
```

## トレースデータの表示

プログラムを実行すると`trace.out`というトレース情報を記録したファイルができています。
次のコマンドを用いると結果をブラウザで見ることができます。

```
$ go tool trace trace.out
```

ブラウザで表示されたページにある`User-defined tasks` -> `draw handler`の`Count`列 -> `goroutine view`の順番で開くと次のような結果が表示されます。

<img src="trace.png" width="500px">
