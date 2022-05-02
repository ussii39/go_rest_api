FROM golang:1.16.2-alpine3.13 AS build

WORKDIR /
COPY . /go/src/github.com/ussii39/go_rest_api
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/google/uuid \
  && go get github.com/gorilla/mux
# main.goの実行ディレクトリ- に移動して、実行ファイルを作成する。
#  go buildをmainパッケージで実行すると/binに実行可能なバイナリファイルが作られる。
# go build [コンパイル後のファイル名] [コンパイルしたいファイル名]
RUN cd /go/src/github.com/ussii39/go_rest_api && go build -o bin/todolist main.go

# アプリケーションの開発用ビルドの依存とランタイムの依存を分離
FROM alpine:3.8
COPY --from=build /go/src/github.com/ussii39/go_rest_api/bin/todolist /usr/local/bin/
# Docker run 時にコマンドの引数として実行
CMD ["todolist"]