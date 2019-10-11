#!/bin/sh
# ------------------------------
# コンテナを起動するスクリプト
# ------------------------------
DBPATH=./docker/db/mysql_data/find_wifi_db

# コンテナの起動
docker-compose up -d

# 初回の場合DBが作られるのを待つ
if [ -d $DBPATH ]; then exit 0
else /bin/echo -n "Initialize database ... "; fi
while [ ! -d $DBPATH ]; do sleep 1; done; echo "done!"