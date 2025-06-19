package schema

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"username":  &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"password":  &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.String},
		"profil":    &graphql.Field{Type: graphql.String},
	},
})

var ObatType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Obat",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"userId":    &graphql.Field{Type: graphql.Int},
		"nama_obat": &graphql.Field{Type: graphql.String},
		"dosis":     &graphql.Field{Type: graphql.String},
		"frekuensi": &graphql.Field{Type: graphql.String},
		"catatan":   &graphql.Field{Type: graphql.String},
		"tanggal":   &graphql.Field{Type: graphql.String},
		"waktu":     &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.String},
	},
})
