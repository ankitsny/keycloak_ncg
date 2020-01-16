package keycloakservice

import (
	"fmt"

	"github.com/Nerzal/gocloak"
	"github.com/anksk333/keycloak_ncg/models"
	"github.com/dgrijalva/jwt-go"
)

// AdminUserID :
const AdminUserID = "ankit"

// AdminPassword :
const AdminPassword = "admin"

// KeycloakHost :
const KeycloakHost = "http://127.0.0.1:8080"

// REALM :
const REALM = "master"

// ClientID :
var ClientID = "account"

// ClientSecret :
var ClientSecret = "11c6d961-9454-4984-a46c-2a90bc2b2ccb"

// Signup :
func Signup(userData models.UserSignup) (string, error) {
	client := gocloak.NewClient(KeycloakHost)

	token, err := client.LoginAdmin(AdminUserID, AdminPassword, REALM)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	user := gocloak.User{
		Username:  userData.UserName,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Enabled:   true,
	}

	uID, err := client.CreateUser(token.AccessToken, REALM, user)

	if err != nil {
		return "", err
	}

	if err := client.SetPassword(token.AccessToken, *uID, REALM, userData.Password, false); err != nil {
		fmt.Println("SetPassword", err)
		return "", err
	}
	return *uID, nil
}

// Login :
func Login(login models.LoginModel) (interface{}, error) {

	client := gocloak.NewClient(KeycloakHost)

	return client.Login(ClientID, ClientSecret, REALM, login.UserName, login.Password)
}

// Me :
func Me(accessToken string) (*jwt.MapClaims, error) {
	client := gocloak.NewClient(KeycloakHost)

	_, claims, err := client.DecodeAccessToken(accessToken, REALM)
	if err != nil {
		return nil, err
	}
	return claims, nil

}
