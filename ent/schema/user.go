package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/vektah/gqlparser/v2/ast"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),

		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),

		field.String("password").
			NotEmpty().
			Sensitive().
			Annotations(
				entgql.OrderField("PASSWORD"),
			),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("建立日期").
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新日期").
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				// 該欄位取消加入註釋
				// entsql.WithComments(false),
			),
	}
}

// 可以使用Mixin加入時間欄位
// func (User) Mixin() []ent.Mixin {
//     return []ent.Mixin{
//         mixin.Time{},
//         // Or, mixin.CreateTime only for create_time
//         // and mixin.UpdateTime only for update_time.
//     }
// }

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M
		edge.To("todos", Todo.Type).
			Annotations(
				// 分頁
				entgql.RelayConnection(),
			),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// 指定表
		entsql.Annotation{Table: "users"},
		// 查詢
		entgql.QueryField(),
		// 分頁
		entgql.RelayConnection(),
		// 新增,修改
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		// 全部欄位加入註釋
		entsql.WithComments(true),
		// Directives功能,對所有填加@hasPermissions,不太好用
		// entgql.Directives(
		// 	HasPermissions([]string{"ADMIN", "OWNER"}),
		// ),
	}
}

// Directives功能,對所有填加@hasPermissions,不太好用
func HasPermissions(permissions []string) entgql.Directive {
	children := make(ast.ChildValueList, 0, len(permissions))
	for _, p := range permissions {
		children = append(children, &ast.ChildValue{
			Value: &ast.Value{
				Raw:  p,
				Kind: ast.StringValue,
			},
		})
	}
	return entgql.NewDirective(
		"hasPermissions",
		&ast.Argument{
			Name: "permissions",
			Value: &ast.Value{
				Children: children,
				Kind:     ast.ListValue,
			},
		},
	)
}
