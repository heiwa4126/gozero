# gozero

go 1.11以前の非モジュール系のプロジェクトの練習。

実行すると
```
$ gozero
version=1.0.0
世界の皆さん、こんにちは!
{English:Apple Japanese:りんご}
42
```
のような文字列が出力される。


# インストール

別に開発とかしたくない、`gozero`コマンドだけ使いたい、という場合。

``` bash
go get -u github.com/heiwa4126/gozero/cmd/gozero
```

レポジトリパスの後ろに`cmd/gozero`がついてるのがミソ

アンインストールは
``` bash
go clean -i github.com/heiwa4126/gozero/cmd/gozero
```


# 開発

## 準備

~/.profileでGOPATHを設定しておく。自分でいろいろ考えたのはこんな感じ

```
export MYGOPATH="$HOME/works/go_project"
export GOPATH="$HOME/.go:/snap/go/current:$MYGOPATH"
export PATH="$HOME/.go/bin:$PATH:$MYGOPATH/bin"
alias gcd='cd $MYGOPATH/src/github.com/heiwa4126'
```

GOPATHの
最初のは`go get`で入れるパス。
2番めはgoのパス(Ubuntuで`sudo snap install --classic go`したので)。
3番め(`$MYGOPATH`)は自分の開発用パス。`gcd`エリアスで移動できる。

あとは
[ghq](https://github.com/motemen/ghq)をgo getして
`git config --global ghq.root "$MYGOPATH"`
しておく。

## ソースの取得

``` bash
mkdir $MYGOPATH/src/github.com/heiwa4126
cd $MYGOPATH/src/github.com/heiwa4126
git clone https://github.com/heiwa4126/gozero
```
または

``` bash
ghq get github.com/heiwa4126/gozero
```

もちろんあなたが開発者、または共同開発者なら
```
https://github.com/heiwa4126/gozero
```
のかわりに
```
git@github.com:heiwa4126/gozero.git
```
を使うこと。



## ビルド

``` bash
cd $MYGOPATH/src/github.com/heiwa4126/gozero
```
して

``` bash
go install
```
で、gozero.aがこのGOPATHのpkgにインストールされる。

また、
``` bash
cd cmd/gozero
go install
```
で、コマンドgozeroがこのGOPATHのbin(つまり`$MYGOPATH/bin`)に
インストールされる。
