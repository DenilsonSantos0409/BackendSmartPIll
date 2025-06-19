package schema

import (
	"smartpill/resolver"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: graphql.NewList(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.GetAllUser(), nil
			},
		},
		"obat": &graphql.Field{
			Type: graphql.NewList(ObatType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.GetAllObat(), nil
			},
		},
		"obatByUser": &graphql.Field{
			Type: graphql.NewList(ObatType),
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.GetObatByUser(uint(p.Args["userId"].(int))), nil
			},
		},
	},
})
