# history

```sh
# 版控初始化
git init -b main
# go專案初始化
go mod init go-ent-gqlgen
# 安裝ent
go get entgo.io/ent/cmd/ent
# 安裝entgql
go get entgo.io/contrib/entgql
# 建立Todo的Model
go run -mod=mod entgo.io/ent/cmd/ent init Todo
# 定義欄位,Fields()
open ent/schema/todo.go
# 定義欄位產生相關文件
go generate ./ent
```
