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
要はhelloworldです。


# デプロイ

別に開発とかしたくない、
`gozero`コマンドだけ使いたい、
という最も普通のケース。

``` sh
go get -u github.com/heiwa4126/gozero/cmd/gozero
```
レポジトリパスの後ろに`cmd/gozero`がついてるのがミソ

または
``` sh
go get -u github.com/heiwa4126/gozero/...
```
でもOK

これで`$GOPATH`の最初のパスの、
src/github.com/heiwa4126/gozeroにソースがpullされて
bin/以下にgozeroがビルドされる。


アンインストールは

``` bash
go clean -i github.com/heiwa4126/gozero/cmd/gozero
# or
go clean -i github.com/heiwa4126/gozero/...
```

ただ、`$GOPATH/src`以下は消えないみたいなので
それは手で消してください。

[clean - The Go Programming Language](https://golang.org/pkg/cmd/go/internal/clean/)
見てもソース消すやつはないみたい。

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
- 最初のパスは`go get`で入れるパス。
- 2番めはgoのパス(Ubuntuで`sudo snap install --classic go`したので)。
- 3番め(`$MYGOPATH`)は自分の開発用パス。`gcd`エリアスで移動できる。


あとは
[ghq](https://github.com/motemen/ghq)をgo getして
`git config --global ghq.root "$MYGOPATH"`
しておく。


この3つGOPATHを使う利点は

- わけがわからなくなったら`rm -rf ~/.go`すれば人のコードはみんな消せる
- 自分のコードと人のコードが混ざらない

ですが

- 自分のコードを`go get`して最初のパスに入れた場合、それを消さないと3番目のパスで開発できない

という欠点があります。



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

もちろんあなたが開発者(または共同開発者)なら
`https://github.com/heiwa4126/gozero`
のかわりに
`git@github.com:heiwa4126/gozero.git`
を使うこと。


## ビルド

*注意*:
go getでgozeroを取得しているなら、
↑を参照してアンインストールしておくこと。
(もしくは3つもGOPATHを使うのを辞める)


```sh
cd $MYGOPATH/src/github.com/heiwa4126/gozero
```
して

```sh
go run ./cmd/gozero
# or
go build ./cmd/gozero ; ./gozero
```
でローカル実行、ローカルビルド


```sh
go install
```
で、gozero.aがこのGOPATHのpkgにインストールされる。

また、
```sh
cd cmd/gozero
go install
# or
go install ./cmd/gozero
```
で、コマンドgozeroがこのGOPATHのbin(つまり`$MYGOPATH/bin`)に
インストールされる。


## TODO

Releaseの練習をすること。

`cmd`は慣習。go本体にcmd/があるので、
```
go build cmd/gozero
```
(`./cmd/`ではなく)すると、結構たいへんなことになる。

「ビルドされるバイナリ名がディレクトリと同じになる」のを利用するだけだったら
`cmd/gozero/`ではなく`gozero/`にmainパッケージを置いたらいいのでは。
