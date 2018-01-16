# WorkStation cleaning up app

Arudino + Golang

## install
```
go get -u github.com/yaaaaashiki/livething
```

## Golang version

* 1.9 


## mysql version

* 5.7 

## 依存パッケージ管理 

[glide](https://github.com/Masterminds/glide) で管理します
新たなパッケージを追加時は

```
glide install
```

上記コマンドをターミナルで打ちます


すると glide.yml が更新され、vendor 以下にパッケージが書き込まれます


## GOPATH の設定

```
if [ -x "`which go`" ]; then
    export GOPATH=$HOME/go
    export PATH=$PATH::$GOPATH/bin
fi
```
上記を `bashrc` or `zshrc` に記述・反映させる


## 作業ディレクトリ

```
GOPATH/src/github.com/yaaaaashiki/livething
```
上記のディレクトリで作業します

`go get` でインストールすると上記ディレクトリに clone されているはず


## dbconfig.yml

`dbconfig.yml` の設定は各々の MySql の設定に依存する

```
datasource: root:@/livething_dev?parseTime=true&collation=utf8_general_ci&interpolateParams=true

datasource: root:@/livething_test?parseTime=true&collation=utf8_general_ci&interpolateParams=true
```

上記は各々修正してください


(ex: password を設定している場合)
```
datasource: root:password@tcp(localhost:3306)/livething_dev?parseTime=true&collation=utf8_general_ci&interpolateParams=true
```


## seed 入れ直し (develop)

```
make developenv 
```


## 起動

```
make run
```


## seed 入れ直し (test)

テスト実行前には必ず実行してください
```
make testenv
```


## テスト
```
make test
```
