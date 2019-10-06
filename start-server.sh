#!/bin/sh
# ------------------------------
# サーバーを起動するスクリプト
# ------------------------------
docker-compose exec app go run app/main.go
