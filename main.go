package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

var funcs = template.FuncMap{
	"increment": func(num int) int {
		return num + 1

	},
}

func main() {

	juan := &Person{"juan", 28, []string{"Leer", "Escribir", "programar"}}
	loadTemplate("index.html", juan)
}

func loadTemplate(fileName string, data interface{}) {
	t, err := template.New(fileName).Funcs(funcs).ParseFiles("templates/" + fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
