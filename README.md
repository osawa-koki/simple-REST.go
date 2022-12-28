# simple-REST.go

Go言語でRESTAPIを構築するサンプルプログラム。  
[gorilla](https://github.com/gorilla/mux)  

## 実行方法

```shell
# デバグ実行
go run main.go

# ビルド
go build -a -x -o bin main.go
```

Dockerfileから実行する場合には以下のコマンドを実行。  

```shell
docker build -t my-golang-app .
docker run -p 80:8080 -it --rm --name my-running-app my-golang-app
```

<http://localhost:8080>へアクセス。  

## デプロイ設定

| キー | バリュー |
| ---- | ---- |
| Name | simple-rest-go |
| Region | Oregon(US West) |
| Branch | main |
| Root Directory |  |
| Environment | Docker |
| Dockerfile Path | ./Dockerfile |
| Docker Build Context Directory |  |
| Docker Command |  |
