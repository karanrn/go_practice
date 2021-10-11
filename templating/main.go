package main

import (
	"os"
	"text/template"
)

type User struct{
	Name string
	Age int
	Country string
	Skills []string
}

func main(){
	u := User { Name: "John", Age: 30, Country: "USA", Skills: []string {"go", "Python", "C#"}}

	// Parse template from file
	ut := template.Must(template.ParseFiles("sample_template.txt")) 

	err := ut.Execute(os.Stdout, u)

	if err != nil {
		panic(err)
	}
}