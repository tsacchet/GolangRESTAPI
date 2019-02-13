package main

import (
	"github.com/gin-gonic/gin"
)

var (
	version string
	sha     string
)

type Checks struct {
	ID            string `gorm:"not null" form:"version" json:"version"`
	Description   string `gorm:"not null" form:"description" json:"description"`
	Lastcommitsha string `gorm:"not null" form:"lastcommitsha" json:"lastcommitsha"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/healthcheck", GetCheck)
	}

	r.Run(":8080")
}

func GetCheck(c *gin.Context) {

	var ver = version
	var lastsha = sha
	var checks = []Checks{
		Checks{ID: ver, Description: "pre-interview technical test", Lastcommitsha: lastsha},
	}

	c.JSON(200, checks)

	// curl -i http://localhost:8080/api/v1/healthcheck
}
