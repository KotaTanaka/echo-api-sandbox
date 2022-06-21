***echo-api-sandbox***

## About

Go (Echo) でのサーバーサイド開発の素振り  
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
# データベースの起動
docker compose up mysql -d

# 環境変数の設定
cp .envrc.sample .envrc
direnv allow

# ローカルサーバー配信
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
docker compose up redoc -d
```

→ http://localhost:1324

- サービス停止（MySQL, ReDoc）

```sh
docker compose down
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
