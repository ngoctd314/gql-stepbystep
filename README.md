# Introduction

A GraphQL service is created by defining types and fields on those types; then providing functions for each field on each type. 

## Who am I

- create tools.go then copy from my repository
- run go mod tidy
- create Makefile, copy then run 'make gqlinit'

- write first scheme
```gql
#  dir: graph/schema.graphqls
type Me {
	id: ID!
	name: String!
}

type Query {
	whoami: Me
}
```

- delete graph/schema.resolvers.go then run 'make gqlgen'
This command will generate resolvers base on your .graphqls file

- update: graph/schema.resolvers.go

```go
// Whoami is the resolver for the whoami field.
func (r *queryResolver) Whoami(ctx context.Context) (*model.Me, error) {
	me := &model.Me{
		ID:   "1",
		Name: "Gophers",
	}

	return me, nil
}
```

- cmd: go run server.go

- access http://localhost:8080

- insert

```txt
query whoami {
  whoami {
    id
	name    
  }
}
```

- result
![result](https://raw.githubusercontent.com/ngoctd314/gql-stepbystep/1-introduction/whoami.png)