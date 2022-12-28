package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty().
			Annotations(
				entgql.OrderField("TEXT"),
			),

		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS").
			Annotations(
				entgql.OrderField("STATUS"),
			),

		field.Int("user_id").
			Annotations(
				entgql.OrderField("USER_ID"),
			),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("todos").
			Unique().
			Field("user_id").
			Required().
			Annotations(entsql.Annotation{
				// 用戶刪除時一起刪除
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// 查詢
		entgql.QueryField(),
		// 分頁
		entgql.RelayConnection(),
		// 新增,修改
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
