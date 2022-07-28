# Queries and Mutations

## Fields

At its simplest, GraphQL is about asking for specific fields on objects. 

```gql
{
    hero {
        name
    }
}
```

- replace graph/schema.graphqls

```gql
type Character {
	id: ID!
	name: String!
}

type Query {
	hero(id: ID!): Character
}

input newHero {
	id: ID!
	name: String!
}

type Mutation {
	createHero(hero: newHero!): Character
}
```

- cmd: rm graph/schema.resolvers.go && make gqlgen

- replace graph/resolver.go

```go
type Resolver struct {
	hero map[string]*model.Character
}
```

- update graph/schema.resolvers.go

```go
// CreateHero is the resolver for the createHero field.
func (r *mutationResolver) CreateHero(ctx context.Context, hero model.NewHero) (*model.Character, error) {
	newHero := &model.Character{
		ID:   hero.ID,
		Name: hero.Name,
	}
	r.Storage[hero.ID] = newHero
	return newHero, nil
}

// Hero is the resolver for the hero field.
func (r *queryResolver) Hero(ctx context.Context, id string) (*model.Character, error) {
	return r.Storage[id], nil
}
```

- update server.go

```go
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Storage: make(map[string]*model.Character),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

- run: go run server.go

![result]()