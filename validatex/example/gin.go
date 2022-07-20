package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yahuian/gox/validatex"
)

type user struct {
	Name string `validate:"required"`
	Age  int    `validate:"gt=18"`
}

func main() {
	if err := validatex.Init(validatex.WithGin()); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		var u user
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": u})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}

// curl -X POST http://localhost:8080/ping -H 'content-type: application/json' -d '{ "name": "tom" }'
// {"msg":"Age must be greater than 18"}
