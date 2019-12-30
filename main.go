package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"hSize": humanSize,
		"hAge":  humanAge,
	})
	r.LoadHTMLGlob("./views/**/*")

	r.GET("/users/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/user.html", gin.H{
			"age": "1",
		})
	})

	r.GET("/goods/good", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/good.html", gin.H{
			"good": "2",
		})
	})
	r.Run("localhost:9091")

	//var size, age string
	//for i := 0; i < 100; i++ {
	//	size += humanSize("")
	//	age += humanAge("")
	//}
	//fmt.Println(age, size)

}

func humanSize(s string) string {
	return strconv.Itoa(rand.Intn(999)) + "mb"
}

func humanAge(s string) string {
	return strconv.Itoa(rand.Intn(150)) + "岁"
}
