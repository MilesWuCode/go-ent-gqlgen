# history

```sh
# 版控初始化
git init -b main
# go專案初始化
go mod init go-ent-gqlgen
# 安裝ent
go get entgo.io/ent/cmd/ent
# 建立Todo的Model
go run -mod=mod entgo.io/ent/cmd/ent init Todo
```
