package controller

import (
	"fmt"
	"strings"

	"github.com/anksk333/keycloak_ncg/keycloakservice"
	"github.com/anksk333/keycloak_ncg/models"
	"github.com/gin-gonic/gin"
)

// KeycloakController :
type KeycloakController struct {
	router gin.IRouter
}

// Register :
func Register(router gin.IRouter) {

	router.POST("/users/signup", func(c *gin.Context) {
		var user models.UserSignup
		if err := c.ShouldBindJSON(&user); err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"message": "Invalid Payload",
			})
		}

		userID, err := keycloakservice.Signup(user)

		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(201, gin.H{
			"userId": userID,
		})
	})

	router.GET("/users/me", func(c *gin.Context) {

		authToken := c.GetHeader("Authorization")
		t := strings.Split(authToken, " ")
		if len(t) != 2 {
			c.JSON(400, gin.H{
				"message": "Invalid auth token",
			})
			return
		}

		token := t[1]

		user, err := keycloakservice.Me(token)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, user)
	})

	router.POST("/users/login", func(c *gin.Context) {
		var login models.LoginModel
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := keycloakservice.Login(login)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, user)
	})

}
