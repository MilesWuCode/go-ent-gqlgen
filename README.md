# go ent gqlgen

## gql 查詢範例

```gql
# 分頁查詢
{
  todos(first: 3, after: null) {
    totalCount
    edges {
      cursor
      node {
        id
        text
        status
        createdAt
      }
    }
    pageInfo {
      startCursor
      endCursor
      hasNextPage
      hasPreviousPage
    }
  }
}

# 單筆查詢
{
  todo(id: "3") {
    id
    text
    status
    createdAt
  }
}
# or
{
  node(id: "3") {
    ... on Todo {
      id
      text
      status
      createdAt
    }
  }
}

# 多筆查詢
{
  nodes(ids: ["3", "4", "999"]) {
    ... on Todo {
      id
      text
      status
      createdAt
    }
  }
}

# 新增
mutation {
  createTodo(input: { text: "abc123", status: IN_PROGRESS }) {
    id
    text
    createdAt
    status
  }
}

# 修改
mutation {
  updateTodo(id: "3", input: { text: "xyz", status: IN_PROGRESS }) {
    id
    text
    createdAt
    status
  }
}

# 刪除
mutation {
  deleteTodo(id: "6")
}
```
