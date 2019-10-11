# find-wifi-api
Find Wi-Fi バックエンド [Echo]

## 技術要素

* 言語 `Go`
* フレームワーク `Echo`
* データベース `MySQL`
* ORマッパー `Gorm`
* 仮想環境 `Docker` `docker-compose`
* API仕様書 `OpenAPI` `ReDoc`

## ローカル開発環境構築

* ソースコードのクローン

```
$ git clone git@github.com:KotaTanaka/find-wifi-api.git
$ cd find-wifi-api
```

* コンテナの起動 (docker-compose up)

```
$ ./start-docker.sh
```

* コンテナの停止 (docker-compose down)

```
$ ./stop-docker.sh
```

* アプリケーションサーバー起動

```
$ ./start-server.sh
```

→ http://localhost:1323 でサーバーが起動します。

* データベースログイン

```
$ ./mysql.sh
Enter password: password
mysql> use find_wifi_db;
```

* データベース初期化

```
$ rm -rf docker/db/mysql_data
$ ./stop-docker.sh && ./start-docker.sh
```

* API仕様書の書き出し

```
$ npm i -g redoc-cli
$ redoc-cli bundle openapi.yml
```

→ コンテナ&サーバー再起動後 http://localhost:1323/doc で確認できます。

## テーブル定義

* Wi-Fiサービステーブル

```
+------------+------------------+------+-----+
| Field      | Type             | Null | Key |
+------------+------------------+------+-----+
| id         | int(10) unsigned | NO   | PRI | auto_increment
| created_at | timestamp        | YES  |     |
| updated_at | timestamp        | YES  |     |
| deleted_at | timestamp        | YES  | MUL |
| wifi_name  | varchar(255)     | YES  |     |
| link       | varchar(255)     | YES  |     |
+------------+------------------+------+-----+
```

* Wi-Fi提供店舗テーブル

```
+---------------+------------------+------+-----+
| Field         | Type             | Null | Key |
+---------------+------------------+------+-----+
| id            | int(10) unsigned | NO   | PRI | auto_increment
| created_at    | timestamp        | YES  |     |
| updated_at    | timestamp        | YES  |     |
| deleted_at    | timestamp        | YES  | MUL |
| service_id    | int(10) unsigned | YES  | MUL |
| ss_id         | varchar(255)     | YES  |     |
| shop_name     | varchar(255)     | YES  |     |
| description   | varchar(255)     | YES  |     |
| address       | varchar(255)     | YES  |     |
| shop_type     | varchar(255)     | YES  |     |
| opening_hours | varchar(255)     | YES  |     |
| seats_num     | int(11)          | YES  |     |
| has_power     | tinyint(1)       | YES  |     |
+---------------+------------------+------+-----+
```

* 店舗レビューテーブル

```
+----------------+------------------+------+-----+
| Field          | Type             | Null | Key |
+----------------+------------------+------+-----+
| id             | int(10) unsigned | NO   | PRI | auto_increment
| created_at     | timestamp        | YES  |     |
| updated_at     | timestamp        | YES  |     |
| deleted_at     | timestamp        | YES  | MUL |
| shop_id        | int(10) unsigned | YES  | MUL |
| comment        | varchar(1000)    | YES  |     |
| evaluation     | int(11)          | YES  |     |
| puplish_status | tinyint(1)       | YES  |     |
+----------------+------------------+------+-----+
```