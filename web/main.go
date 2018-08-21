package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var users = map[string]map[string]string{
	"1": map[string]string{"id": "1", "name": "wxnacy"},
	"2": map[string]string{"id": "2", "name": "wxnacy2"},
}

func GetUsers(c *gin.Context) {
	// fmt.Println(c.Query("id"))
	var id = c.DefaultQuery("id", "0")
	fmt.Println(id)

	var res = []map[string]string{}

	if id != "0" {
		res = append(res, users[id])
	} else {

		for _, v := range users {
			res = append(res, v)
		}

	}

	c.JSON(http.StatusOK, res)
}

func Start() {
	r := gin.Default()

	r.GET("/user/:id", GetUserById)
	r.GET("/user", GetUsers)
	r.POST("/user", CreateUser)
	r.POST("/upload", Upload)

	r.Run(":6666") // listen and serve on 0.0.0.0:8080
}
