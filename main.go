package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	r := gin.Default()

	//r.SetFuncMap(template.FuncMap{
	//	"hSize": humanSize,
	//	"hAge":  humanAge,
	//})
	//r.LoadHTMLGlob("./views/**/*")

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
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
}

func humanSize(s string) string {
	return strconv.Itoa(rand.Intn(999)) + "mb"
}

func humanAge(s string) string {
	return strconv.Itoa(rand.Intn(150)) + "岁"
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		//t, err = t.New(name).Parse(string(h))
		t = template.Must(t.New(name).
			Funcs(template.FuncMap{
				"hSize": humanSize,
				"hAge":  humanAge,
			}).
			Parse(string(h)))
		//if err != nil {
		//	return nil, err
		//}
	}
	return t, nil
}
