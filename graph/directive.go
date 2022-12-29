package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

func HasPermission() func(context.Context, interface{}, graphql.Resolver, []string) (interface{}, error) {
	return func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
		permissions []string,
	) (res interface{}, err error) {
		// 權限
		log.Println(permissions)

		// 找出用戶並核對權限
		// 或寫相對應邏輯

		if false {
			return nil, fmt.Errorf("Access denied")
		}

		return next(ctx)
	}
}
