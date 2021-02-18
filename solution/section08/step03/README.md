# STEP03: バージョンアップをしてみよう

## 新しく学ぶこと

* GAE上のアプリのバージョンアップのしかた

## デプロイの仕方

```
$ gcloud app deploy --project gohandson-gacha --version v2
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
$ go build -v -o step03
$ DATASTORE_EMULATOR_HOST=localhost:8081 ./step03
```

または

```
$ DATASTORE_EMULATOR_HOST=localhost:8081 go run .
```
