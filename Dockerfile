FROM golang:1.18.0-alpine AS build

WORKDIR /
RUN go install github.com/cosmtrek/air@latest

# go.modに追加 +mainなどでpackage import + install + downloadで読み込める
RUN go install github.com/joho/godotenv/cmd/godotenv@latest
# RUN go mod download github.com/joho/godotenv

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)

COPY . /go/src/github.com/ussii39/go_rest_api
# CMD ["air", "-c", ".air.toml"]

# main.goの実行ディレクトリ- に移動して、実行ファイルを作成する。
#  go buildをmainパッケージで実行すると/binに実行可能なバイナリファイルが作られる。
# go build [コンパイル後のファイル名] [コンパイルしたいファイル名]
RUN cd /go/src/github.com/ussii39/go_rest_api && go build -o bin/todolist main.go

# アプリケーションの開発用ビルドの依存とランタイムの依存を分離
FROM alpine:3.8
COPY --from=build /go/src/github.com/ussii39/go_rest_api/bin/todolist /usr/local/bin/
# Docker run 時にコマンドの引数として実行
# CMD ["todolist"]

# RUN go install github.com/cosmtrek/air@latest
COPY  ./ /go/src/
# WORKDIR /go/src
COPY /start.sh /start.sh
RUN chmod 744 /start.sh
# 作業用ディレクトリを指定

CMD ["todolist"]


# CMD [ "go" "run" "./main.go" ]