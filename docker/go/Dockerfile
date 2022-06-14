FROM golang:1.17.6-alpine

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app/server
RUN go build
# ホットリロード用のapps「air」の導入
RUN go install github.com/cosmtrek/air@latest
CMD air
