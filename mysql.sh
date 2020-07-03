#!/bin/sh
# ------------------------------
# コンテナ内のMySQLに入るスクリプト
# ------------------------------
docker-compose exec mysql bash -c "mysql -u root -p"
