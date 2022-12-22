# history

```sh
# 版控初始化
git init -b main

# go專案初始化
go mod init go-ent-gqlgen

# 安裝ent,ORM工具
go get entgo.io/ent/cmd/ent

# 安裝entgql,ent的擴展和工具集合
go get entgo.io/contrib/entgql

# 建立Todo的Model
go run -mod=mod entgo.io/ent/cmd/ent init Todo

# 定義欄位,Fields(), Annotations()
open ./ent/schema/todo.go

# 定義CRUD邏輯
open ./graph/schema.resolvers.go

# 定義欄位產生相關檔案
go generate ./ent

# 安裝gqlgen,graphql工具
go get github.com/99designs/gqlgen

# 建立gqlgen的tool.go
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

# 檢查套件相依
go mod tidy

# gqlgen初始化
go run github.com/99designs/gqlgen init

# 測試可否執行server.go服務
go run server.go

# 刪除server.go,準備和ent合併
rm ./server.go

# 修改gqlgen.yml設定檔
open ./gqlgen.yml

# 修改resolver.go,加入ent.Client,NewSchema
open ./resolver.go

# 建立generate腳本
open ./generate.go

# 產生相關檔案
go generate .

# 建立main.go
touch main.go & open main.go

# 安裝mysql連線套件
go get github.com/go-sql-driver/mysql

# 檢查套件相依
go mod tidy

go run main.go
```
