package main

import (
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/graph"
	"log"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// sqlite
	// _ "github.com/mattn/go-sqlite3"

	// mysql
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {
	// 使用sqlite在ram
	// client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")

	// 使用sqlite在data.db
	// client, err := ent.Open(dialect.SQLite, "file:data.db&_fk=1")

	// 使用mysql資料庫
	client, err := ent.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go_ent_gqlgen?parseTime=True")

	// 資料庫連線失敗
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	// 關閉連線
	defer client.Close()

	// 自動migration,有版本控制建議關閉
	// if err := client.Schema.Create(
	// 	context.Background(),
	// 	migrate.WithGlobalUniqueID(true),
	// ); err != nil {
	// 	log.Fatal("running schema migration", err)
	// }

	// 建立服務器並注入資料庫連線
	srv := handler.NewDefaultServer(graph.NewSchema(client))

	// 資料庫使用交易
	srv.Use(entgql.Transactioner{TxOpener: client})

	// 除錯追蹤,不用可關閉
	// srv.Use(&debug.Tracer{})

	// 原生http,GraphQL的Endpoint
	// http.Handle("/query", srv)

	// 原生http,GraphQL的Playground介面
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// 預設埠
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// 提示
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	// 開啓服務器並設定埠
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal("http server terminated", err)
	// }

	// 使用Gin
	r := gin.Default()

	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").
			ServeHTTP(c.Writer, c.Request)
	})

	r.Run()
}
