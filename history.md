# history

```sh
# 版控初始化
git init -b main

# go專案初始化
go mod init go-ent-gqlgen

# 安裝ent,ORM工具
go get entgo.io/ent/cmd/ent@master

# 安裝entgql,ent的擴展和工具集合
go get entgo.io/contrib@master

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

# 填加邏輯,查詢單筆/新增/修改...
code ./graph/todo.resolvers.go

# 啓動服務,檢查:查詢單筆/新增/修改的結果
go run main.go

### 使用 versioned-migration ###

# 關閉自動migration功能,client.Schema.Create()
open ./main.go

# 修改entc.go,加入Features: []gen.Feature{gen.FeatureVersionedMigration},
open ./ent/entc.go

# 建立自定義生成
touch ./ent/migrate/main.go
open ./ent/migrate/main.go

# 建立空資料夾migrations,存放sql記錄檔
mkdir ./ent/migrate/migrations

# 執行ent/migrate/main.go,自動建立xxx_create_todos.sql記錄檔,atlas.sum校驗檔
go run -mod=mod ent/migrate/main.go create_todos

# 執行lint,用來驗證sql記錄檔與專案生成代碼是否有差異
atlas migrate lint \
 --dev-url="mysql://root:password@localhost:3306/test" \
  --dir="file://ent/migrate/migrations" \
  --latest=1

# 新增一個欄位updated_at
open ./ent/schema/todo.go

# 自動建立xxx_add_updated_at_to_todos.sql記錄檔
go run -mod=mod ent/migrate/main.go add_updated_at_to_todos

# 執行lint,顯示出updated_at欄位差異
atlas migrate lint \
 --dev-url="mysql://root:password@localhost:3306/test" \
  --dir="file://ent/migrate/migrations" \
  --latest=1

# 執行apply更新資料庫,因為是中途才用版控資料庫所以加入--baseline xxx
# xxx為./ent/migrate/migrations/20221223061143_create_todos.sql前面日期
# 之後的版控不需要--baseline
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url mysql://root:password@localhost:3306/go_ent_gqlgen \
  --baseline 20221223061143

### user 與 todo 關係 ###

# 建立user
go run -mod=mod entgo.io/ent/cmd/ent init User

# user欄位Fields()和關係Edges()
code ./ent/schema/user.go

# todo關係Edges()
code ./ent/schema/todo.go

# 產生相關檔案
go generate .

# 生成SQL
go run -mod=mod ent/migrate/main.go create_users

# 檢查SQL
atlas migrate lint \
 --dev-url="mysql://root:password@localhost:3306/test" \
  --dir="file://ent/migrate/migrations" \
  --latest=1

# 若有缺預設值,修改完後執行hash
atlas migrate hash \
  --dir "file://ent/migrate/migrations"

# 執行apply更新資料庫
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url mysql://root:password@localhost:3306/go_ent_gqlgen

# 修改autobind
code ./gqlgen.yml

# 建立user.graphqls
code ./graph/user.graphqls

# 產生user.graphqls相關檔案user.resolvers.go
go generate .

# 填加邏輯,查詢單筆/新增/修改...
code ./graph/user.resolvers.go

### 客製化欄位 ###

# 對User填加firstLetter: String! #gqlgen:resolver
code ./graph/user.graphqls

# 執行使user.resolvers.go產生FirstLetter()
go generate .

# 填加FirstLetter()邏輯
code ./graph/user.resolvers.go

### 查詢user裡的todos並使用分頁功能 ###

# 原本提供
# ./ent/schema/user.go的Edges()加入Annotations(entgql.RelayConnection())
code ./ent/schema/user.go
go generate .

# 自製分頁
# ./graph/user.graphqls填加todoPages欄位
code ./graph/user.graphqls
go generate .

# 填加user.resolvers.go裡的todoPages()邏輯
code ./graph/user.resolvers.go

### 填加 directive @hasPermissions 功能  ###

# 並填加刪除用戶的功能 deleteUser(id: ID!): Boolean! @hasPermissions(permissions: ["ADMIN"])
code ./graph/user.graphqls
go generate .
code ./graph/directive.go
code ./graph/resolver.go
go generate .
```
