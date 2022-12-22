# go ent gqlgen

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
  todos(first: 3, after: null) {
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
      endCursor
      hasNextPage
      startCursor
      hasPreviousPage
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
```
