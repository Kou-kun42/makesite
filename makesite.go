package main

import "fmt"
import "html/template"
import "io/ioutil"
import "os"
import "flag"
import "strings"


func save(txt *string) {
	file, err := ioutil.ReadFile(*txt)
	if err != nil {
		panic(err)
	}
	name := strings.Split(*txt, ".")[0]
	html, err := os.Create(fmt.Sprintf("%s.html", name))
	if err != nil {
		panic(err)
	}
	tem := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = tem.Execute(html, file)
	if err != nil {
		panic(err)
	}
}

func main() {
	// fmt.Println("Hello, world!")

	// // Read in file
	// file, err := ioutil.ReadFile("first-post.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// // Render and print file
	// tem := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// err = tem.Execute(os.Stdout, string(file))
	// if err != nil {
	// 	panic(err)
	// }

	// // Write to html template
	// html, err := os.Create("first-post.html")
	// err = tem.Execute(html, string(file))
	// if err != nil {
	// 	panic(err)
	// }

	// Add flag to file
	file := flag.String("file", "", "txt file to be rendered")
	flag.Parse()
	save(file)
}
