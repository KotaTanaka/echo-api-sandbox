#!/bin/sh
# ------------------------------
# サーバーを起動するスクリプト
# ------------------------------
docker-compose exec app sh -c "go run app/main.go"
