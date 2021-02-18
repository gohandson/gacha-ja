# STEP02: Google Cloud Datastoreにガチャ結果を保存してみよう

## 新しく学ぶこと

* Google Cloud Datastore

## デプロイの仕方

```
$ gcloud app deploy --project gohandson-gacha --version v1
```

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
$ go build -v -o step02
$ DATASTORE_EMULATOR_HOST=localhost:8081 ./step02
```

または

```
$ DATASTORE_EMULATOR_HOST=localhost:8081 go run .
```

