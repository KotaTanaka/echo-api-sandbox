***echo-api-sandbox***

## About

Go (Echo) でのサーバーサイド開発の実験場  
Wi-Fi検索システムのバックエンドAPI

*[管理画面 - react-frontend-sandbox](https://github.com/KotaTanaka/react-frontend-sandbox)*

## Technology

* 言語 - `Go`
* フレームワーク - `Echo`
* データベース - `MySQL`
* ORM - `Gorm`
* 開発環境 - `Docker` `docker-compose`
* API定義書 - `OpenAPI` `ReDoc`

## Getting Started

* サービスの起動

```bash
# 初回
$ docker-compose up --build -d

# 2回目以降
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

## Package Architecture

```
find-wifi-backend
├── data
│   ├── admin
│   │   ├── area.go
│   │   ├── review.go
│   │   ├── service.go
│   │   └── shop.go
│   ├── client
│   │   ├── area.go
│   │   ├── review.go
│   │   └── shop.go
│   └── common.go
├── handler
│   ├── admin
│   │   ├── area.go
│   │   ├── review.go
│   │   ├── service.go
│   │   └── shop.go
│   ├── client
│   │   ├── area.go
│   │   ├── review.go
│   │   └── shop.go
│   └── hello.go
├── main.go
├── model
│   ├── area.go
│   ├── review.go
│   ├── service.go
│   └── shop.go
└── server
    ├── database.go
    └── router.go
```
