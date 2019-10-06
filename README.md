# find-wifi-api
Find Wi-Fi バックエンド [Echo]

## 技術要素

* 言語 `Go`
* フレームワーク `Echo`
* データベース `MySQL`
* ORマッパー `Gorm`
* 仮想環境 `Docker` `docker-compose`
* API仕様書 `OpenAPI` `ReDoc`

## ローカル開発環境構築

* ソースコードのクローン

```
$ git clone git@github.com:KotaTanaka/find-wifi-api.git
$ cd find-wifi-api
```

* コンテナの起動 (docker-compose up)

```
$ ./start-docker.sh
```

* コンテナの停止 (docker-compose down)

```
$ ./stop-docker.sh
```

* アプリケーションサーバー起動

```
$ ./start-server.sh
```

→ http://localhost:1323 でサーバーが起動します。

* データベースログイン

```
$ ./mysql.sh
Enter password: password
mysql> use find_wifi_db;
```

* API仕様書の書き出し

```
$ npm i -g redoc-cli
$ redoc-cli bundle openapi.yml
```