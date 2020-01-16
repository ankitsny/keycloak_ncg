package main

import (
	"fmt"

	"github.com/Nerzal/gocloak"
	controller "github.com/anksk333/keycloak_ncg/controllers"
	"github.com/gin-gonic/gin"
)

// User :
type User struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:":lastName"`
	Enabled   *bool   `json:"enabled"`
	UserName  *string `json:"userName"`
}

var userData = `{
	"firstName": "Ankit",
	"lastName": "Kumar",
	"enabled": true,
	"userName: "ankso"
}
`

// REALM :
const REALM = "master"

// AdminUserID :
const AdminUserID = "ankit"

// AdminPassword :
const AdminPassword = "admin"

// KeycloakHost :
const KeycloakHost = "http://127.0.0.1:8080"

// CreateUserByID :
func CreateUserByID(client gocloak.GoCloak, token *gocloak.JWT, id string) string {

	user := gocloak.User{
		FirstName: "Ankit",
		LastName:  "Kumar",
		Enabled:   true,
		Username:  id,
	}

	nUID, err1 := client.CreateUser(token.AccessToken, REALM, user)

	if err1 != nil {
		fmt.Println(err1.Error())
		panic(err1)
	}

	return *nUID
}

func main() {

	router := gin.Default()

	controller.Register(router)

	router.Run(":8001")

}
