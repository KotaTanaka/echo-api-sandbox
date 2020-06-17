***find-wifi-backend***

## About

Find Wi-Fi バックエンド RESTful API

*[管理コンソールUI - find-wifi-console-app](https://github.com/KotaTanaka/find-wifi-console-app)*

## Technology

* 言語 `Go`
* フレームワーク `Echo`
* データベース `MySQL`
* ORマッパー `Gorm`
* 仮想環境 `Docker` `docker-compose`
* API仕様書 `OpenAPI` `ReDoc`

## Getting Started

* インストール

```bash
$ git clone git@github.com:KotaTanaka/find-wifi-backend.git
$ cd find-wifi-backend
```

* サービスの起動

```bash
$ docker-compose up -d
```

* アプリケーションの起動

```bash
$ ./start-server.sh
```

→ http://localhost:1323

* サービスの停止

```bash
$ docker-compose down
```

## Utility Commands

* データベースログイン

```bash
$ ./mysql.sh
Enter password: password
mysql> use find_wifi_db;
```

* データベース初期化

```bash
# DB削除
$ rm -rf docker/db/mysql_data

# サービス再起動(DB再生成)
$ docker-compose down && docker-compose up -d
```

* API定義書生成

```bash
# OpenAPIからReDocへの書き出し
$ ./redoc.sh
```

→ http://localhost:1323/doc

## Database

データベース名 `find_wifi_db`

| テーブル物理名 | 論理名 |
|:---|:---|
| `services` | Wi-Fiサービス |
| `shops` | Wi-Fi提供店舗 |
| `reviews` | 店舗レビュー |
| `areas` | エリアマスタ |
