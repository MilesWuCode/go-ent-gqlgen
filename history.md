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
code ./ent/schema/todo.go

# 定義CRUD邏輯
code ./graph/schema.resolvers.go

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
code ./gqlgen.yml

# 修改resolver.go,加入ent.Client,NewSchema
code ./resolver.go

# 建立generate腳本
code ./generate.go

# 產生相關檔案
go generate .

# 建立main.go
touch ./main.go
code ./main.go

# 安裝mysql連線套件
go get github.com/go-sql-driver/mysql

# 檢查套件相依
go mod tidy

# 啓動服務
go run main.go

# 建立todo.graphqls,新增及修改
touch ./graph/todo.graphqls
code ./graph/todo.graphqls

# 產生相關檔案,./graph/todo.resolvers.go
go generate .

# 填加邏輯,查詢單筆/新增/修改
code ./graph/todo.resolvers.go

# 啓動服務,檢查:查詢單筆/新增/修改的結果
go run main.go
```
