package main

import "fmt"
import "html/template"
import "io/ioutil"
import "os"

func main() {
	fmt.Println("Hello, world!")

	// Read in file
	file, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	// Render and print file
	tem := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = tem.Execute(os.Stdout, string(file))
	if err != nil {
		panic(err)
	}

	// Write to html template
	html, err := os.Create("first-post.html")
	err = tem.Execute(html, string(file))
	if err != nil {
		panic(err)
	}
}
