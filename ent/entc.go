//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		// 專案資料夾開始
		// 設定schema位置
		entgql.WithSchemaPath("graph/schema.graphqls"),
		// 設定gqlgen.yml位罝
		entgql.WithConfigPath("gqlgen.yml"),
		// ?
		entgql.WithWhereInputs(true),
		// ?
		// entgql.WithNodeDescriptor(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}

	// 原本未使用versioned-migrations
	// if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
	// 	log.Fatalf("running ent codegen: %v", err)
	// }

	// 使用versioned-migrations
	if err := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
