FROM golang:1.18.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
# COPY . ./app
# WORKDIR /
# VOLUME ["./app"]
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
FROM alpine:3.8 AS Deploy
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

# => [internal] load build definition from Dockerfile                                                                                                                                       0.0s
#  => => transferring dockerfile: 1.59kB                                                                                                                                                     0.0s
#  => [internal] load .dockerignore                                                                                                                                                          0.0s
#  => => transferring context: 2B                                                                                                                                                            0.0s
#  => [internal] load metadata for docker.io/library/golang:1.18.0-alpine                                                                                                                    2.1s
#  => [internal] load metadata for docker.io/library/alpine:3.8                                                                                                                              2.7s
#  => [build 1/5] FROM docker.io/library/golang:1.18.0-alpine@sha256:a2ca4f4c0828b1b426a3153b068bf32a21868911c57a9fc4dccdc5fbb6553b35                                                        0.0s
#  => [deploy 1/5] FROM docker.io/library/alpine:3.8@sha256:2bb501e6173d9d006e56de5bce2720eb06396803300fe1687b58a7ff32bf4c14                                                                 1.4s
#  => => resolve docker.io/library/alpine:3.8@sha256:2bb501e6173d9d006e56de5bce2720eb06396803300fe1687b58a7ff32bf4c14                                                                        0.0s
#  => => sha256:2bb501e6173d9d006e56de5bce2720eb06396803300fe1687b58a7ff32bf4c14 1.41kB / 1.41kB                                                                                             0.0s
#  => => sha256:e802987f152d7826cf929ad4999fb3bb956ce7a30966aeb46c749f9120eaf22c 528B / 528B                                                                                                 0.0s
#  => => sha256:b22edbe95d11980cbd47579189e8e6382dbb39629ca636f2a3b2741fae23bf1d 1.51kB / 1.51kB                                                                                             0.0s
#  => => sha256:788aef77d06ba65af872cf0c2da5b49362f6c05a5c8d1f8ceb4fd8b070e56876 2.10MB / 2.10MB                                                                                             0.8s
#  => => extracting sha256:788aef77d06ba65af872cf0c2da5b49362f6c05a5c8d1f8ceb4fd8b070e56876                                                                                                  0.4s
#  => [internal] load build context                                                                                                                                                          0.1s
#  => => transferring context: 34.48kB                                                                                                                                                       0.1s
#  => CACHED [build 2/5] WORKDIR /app                                                                                                                                                        0.0s
#  => CACHED [build 3/5] RUN go install github.com/joho/godotenv/cmd/godotenv@latest                                                                                                         0.0s
#  => [build 4/5] COPY . /go/src/github.com/ussii39/go_rest_api                                                                                                                              1.7s
#  => [build 5/5] RUN cd /go/src/github.com/ussii39/go_rest_api && go build -o bin/todolist main.go                                                                                          3.1s
#  => [deploy 2/5] COPY --from=build /go/src/github.com/ussii39/go_rest_api/bin/todolist /usr/local/bin/                                                                                     0.0s 
#  => [deploy 3/5] COPY  ./ /go/src/                                                                                                                                                         0.6s 
#  => [deploy 4/5] COPY /start.sh /start.sh                                                                                                                                                  0.0s 
#  => [deploy 5/5] RUN chmod 744 /start.sh


# 本番環境起動 docker build -t test_go_deploy:latest --target Deploy ./
# cd /usr/local/bin/todolist

FROM golang:1.18.2 as dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]

# docker build -t test_go:latest ./

# => [internal] load build definition from Dockerfile                                                                                                                                       0.0s
#  => => transferring dockerfile: 1.70kB                                                                                                                                                     0.0s
#  => [internal] load .dockerignore                                                                                                                                                          0.0s
#  => => transferring context: 2B                                                                                                                                                            0.0s
#  => [internal] load metadata for docker.io/library/golang:1.18.2                                                                                                                           2.5s
#  => [internal] load build context                                                                                                                                                          0.1s
#  => => transferring context: 34.59kB                                                                                                                                                       0.1s
#  => [dev 1/6] FROM docker.io/library/golang:1.18.2@sha256:04fab5aaf4fc18c40379924674491d988af3d9e97487472e674d0b5fd837dfac                                                                 0.0s
#  => CACHED [dev 2/6] WORKDIR /app                                                                                                                                                          0.0s
#  => CACHED [dev 3/6] COPY go.mod go.sum ./                                                                                                                                                 0.0s
#  => CACHED [dev 4/6] RUN go mod download                                                                                                                                                   0.0s
#  => [dev 5/6] COPY . .                                                                                                                                                                     0.9s
#  => [dev 6/6] RUN go install github.com/cosmtrek/air@latest                                                                                                                                4.5s
#  => exporting to image                                                                                                                                                                     0.4s
#  => => exporting layers                                                                                                                                                                    0.3s
#  => => writing image sha256:cc2d084c77c18449660d553db7d76ea6b14164fa53d4682a87af9044cfa588e1                                                                                               0.0s 
#  => => naming to docker.io/library/test_go:latest
