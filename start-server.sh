#!/bin/bash
# ------------------------------
# サーバーを起動するスクリプト
# ------------------------------
DBPATH=./docker/db/mysql_data/find_wifi_db
RUN="docker-compose exec app sh -c \"go run app/main.go\""

echo "========================="
echo "  Find Wi-Fi API Server"
echo "========================="

# 初回の場合DBが作られるのを待つ
if [ -d $DBPATH ]; then
  eval $RUN
  exit 0
else
  /bin/echo -n "Initialize database ... "
fi
while [ ! -d $DBPATH ]; do
  sleep 1
done
echo "done."
eval $RUN
