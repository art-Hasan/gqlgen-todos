# gqlgen-todos

Backend for todo's using graphql

### Resources

* [GraphQL](https://gqlgen.com/getting-started/) generator
  * See `ent/schema` for definitions
* [Ent](https://entgo.io/) (ORM)
  * See `graph/schema.graphqls`
* [Wire](https://github.com/google/wire) dependency management
  * See `cmd/wire_gen.go`

### Examples

```
graphql
mutation createTodo {
  createTodo(input:{text:"t", userId:2}) {
    user {
      id
      name
    }
    text
    done
  }
}

mutation createUser {
  createUser(input:{name:"Tom"}) {
    id
    name
  }
}

mutation deleteUser {
  deleteUser(input:{id:2}) {
    id
  }
}

query findTodos {
  	todos {
      page {
        id
        text
      	done
        user {
          id
          name
        }
      }
      info {
        total
      }
    }
}

query findUsers {
  users {
    page {
      id
      name
    }
    info {
      total
    }
  }
}
```
