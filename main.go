package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	//r.SetFuncMap(template.FuncMap{
	//	"hSize": humanSize,
	//	"hAge":  humanAge,
	//})
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"Foo": "World",
		})
	})
	r.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/bar.tmpl", gin.H{
			"Bar": "World",
		})
	})
	r.Run(":9091")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		//t, err = t.New(name).Parse(string(h))
		t = template.Must(t.New(name).Funcs(template.FuncMap{
			"hSize": humanSize,
			"hAge":  humanAge,
		}).Parse(string(h)))
		//if err != nil {
		//	return nil, err
		//}
	}
	return t, nil
}

func humanSize(s string) string {
	return strconv.Itoa(rand.Intn(999)) + "mb"
}

func humanAge(s string) string {
	return strconv.Itoa(rand.Intn(150)) + "岁"
}
