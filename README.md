# find-wifi-api
Find Wi-Fi バックエンド [Echo]

## 開発環境構築

* direnv のインストール

```
$ brew install direnv
```

* `~/.bashrc` に下記を追記

```
eval "$(direnv hook bash)"
```

* 上記変更を反映

```
$ source ~/.bashrc
```

* プロジェクトルートに下記内容の `.envrc` を作成する。

```
export GOPATH=$(pwd)
```

* ソースコードのクローン・依存関係のインストール

*※ 本環境構築手順に則る場合、プロジェクトの展開場所はどこでもOKです。*

```
$ git clone git@github.com:KotaTanaka/find-wifi-api.git
$ cd src/find-wifi-api
$ dep ensure
```

* アプリケーションの起動

```
$ go run main.go
```