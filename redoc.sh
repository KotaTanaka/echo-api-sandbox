#!/bin/sh
# ------------------------------
# openapi.yml からAPI仕様書を生成するスクリプト
# ------------------------------
docker-compose exec app sh -c "redoc-cli bundle app/openapi.yml -o app/redoc.html"

if [ $? = 0 ]; then
  echo "\n[API仕様書] http://localhost:1323/doc"
fi
