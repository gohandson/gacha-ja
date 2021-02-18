# STEP04: ガチャAPIに繋がらない時の対策をしよう

## 新しく学ぶこと

* コンテキスト
* タイムアウト

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
$ go build -v -o step04
$ DATASTORE_EMULATOR_HOST=localhost:8081 ./step04
```

または

```
$ DATASTORE_EMULATOR_HOST=localhost:8081 go run .
```
