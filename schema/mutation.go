package schema

import (
	"smartpill/resolver"
	"strings"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"profil":   &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.CreateUser(
					p.Args["username"].(string),
					p.Args["email"].(string),
					p.Args["password"].(string),
				), nil
			},
		},

		"updateUser": &graphql.Field{
			Type:        UserType,
			Description: "Update data user",
			Args: graphql.FieldConfigArgument{
				"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"username": &graphql.ArgumentConfig{Type: graphql.String},
				"email":    &graphql.ArgumentConfig{Type: graphql.String},
				"password": &graphql.ArgumentConfig{Type: graphql.String},
				"profil":   &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := uint(p.Args["id"].(int))

				username, _ := p.Args["username"].(string)
				email, _ := p.Args["email"].(string)
				password, _ := p.Args["password"].(string)

				var profil *string
				if p.Args["profil"] != nil {
					profilStr := p.Args["profil"].(string)
					profil = &profilStr
				}

				return resolver.UpdateUser(id, &username, &email, &password, profil)
			},
		},

		"createObat": &graphql.Field{
			Type: ObatType,
			Args: graphql.FieldConfigArgument{
				"userId":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"nama_obat": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"dosis":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"frekuensi": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"catatan":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"tanggal":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"waktu":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String)))},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				waktuRaw := p.Args["waktu"].([]interface{})
				var waktuList []string
				for _, w := range waktuRaw {
					if s, ok := w.(string); ok {
						waktuList = append(waktuList, s)
					}
				}

				// Simpan dalam bentuk string gabungan, misalnya "08:00,12:00"
				waktuStr := strings.Join(waktuList, ",")

				return resolver.CreateObat(
					uint(p.Args["userId"].(int)),
					p.Args["nama_obat"].(string),
					p.Args["dosis"].(string),
					p.Args["frekuensi"].(string),
					p.Args["catatan"].(string),
					p.Args["tanggal"].(string),
					waktuStr,
				), nil
			},
		},

		"updateObat": &graphql.Field{
			Type: ObatType,
			Args: graphql.FieldConfigArgument{
				"id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"nama_obat": &graphql.ArgumentConfig{Type: graphql.String},
				"dosis":     &graphql.ArgumentConfig{Type: graphql.String},
				"frekuensi": &graphql.ArgumentConfig{Type: graphql.String},
				"catatan":   &graphql.ArgumentConfig{Type: graphql.String},
				"tanggal":   &graphql.ArgumentConfig{Type: graphql.String},
				"waktu":     &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.UpdateObat(
					uint(p.Args["id"].(int)),
					p.Args,
				)
			},
		},

		"login": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				password := p.Args["password"].(string)

				return resolver.LoginUser(email, password)
			},
		},

		"resetPassword": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				password := p.Args["password"].(string)

				return resolver.ResetPassword(email, password)
			},
		},

		"deleteObat": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolver.DeleteObat(uint(p.Args["id"].(int)))
			},
		},
	},
})
