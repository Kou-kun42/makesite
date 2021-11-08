package main

import "fmt"
import "html/template"
import "io/ioutil"
import "os"
import "flag"
import "strings"


type HTML struct {
	Name string
	Text string
}

func save(txt string) {
	file, err := ioutil.ReadFile(txt)
	if err != nil {
		panic(err)
	}
	name := strings.Split(txt, ".")[0]
	html, err := os.Create(fmt.Sprintf("%s.html", name))
	if err != nil {
		panic(err)
	}

	convert := HTML{name, string(file)}

	tem := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = tem.Execute(html, convert)
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
	var file string
	var dir string

	flag.StringVar(&file, "file", "", "txt file to be rendered")
	flag.StringVar(&dir, "dir", "", "dir to be used")
	flag.Parse()

	if file != "" {
		save(file)
	}

	if dir != "" {
		var files []string

		allFiles, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}

		for _, f := range allFiles {
			name := strings.Split(f.Name(), ".")
			if name[len(name)-1] == "txt" {
				files = append(files, f.Name())
			}
		}

		for _, txt := range files {
			save(txt)
		}
	}
	
}
