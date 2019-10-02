#!/bin/sh
# ------------------------------
# コンテナ内のMySQLに入るスクリプト
# ------------------------------
docker exec -it find_wifi_mysql \
  bash -c "mysql -u root -p"
