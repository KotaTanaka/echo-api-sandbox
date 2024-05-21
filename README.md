***echo-api-sandbox***

## About

Go (Echo) でのサーバーサイドDDDの素振り  
Wi-Fi検索システムのバックエンドAPI

*[管理画面 - react-frontend-sandbox](https://github.com/KotaTanaka/react-frontend-sandbox)*

## Tech Stack

- 言語 - `Go`
- フレームワーク - `Echo`
- データベース - `MySQL`
- ORM - `Gorm`
- API定義書 - `OpenAPI` `ReDoc`

## Requirements

- Go 1.18
- Docker, docker-compose
- direnv

## Getting Started

- アプリケーションの起動

```sh
# ミドルウェアの起動
make up

# 環境変数の設定
cp .envrc.sample .envrc
direnv allow

# サーバー起動
make run
```

→ http://localhost:1323


## Utility Commands

- データベースログイン

```sh
./mysql.sh
Enter password: password
mysql> use find_wifi_db;
```

- API定義書生成（ReDocUI配信）

```sh
# `make up` していれば不要
docker compose up redoc -d
```

→ http://localhost:1324

- サービス停止（DB, ReDoc）

```sh
make down
```

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
src
├── application
│   ├── dto
│   │   ├── admin
│   │   │   └── {xxx}.go
│   │   ├── client
│   │   │   └── {xxx}.go
│   │   ├── common.go
│   │   └── error.go
│   └── usecase
│       ├── admin
│       │   └── {xxx}.go
│       └── client
│           └── {xxx}.go
├── domain
│   ├── model
│   │   └── {xxx}.go
│   └── repository
│       └── {xxx}.go
├── go.mod
├── go.sum
├── handler
│   ├── admin
│   │   └── {xxx}.go
│   └── client
│       └── {xxx}.go
├── infrastructure
│   └── gorm.go
├── lib
│   └── validator.go
├── main.go
├── registry
│   ├── admin.go
│   └── client.go
└── router
    ├── admin.go
    └── client.go
```
