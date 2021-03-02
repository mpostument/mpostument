package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/mmcdole/gofeed"
)

type ReadmeData struct {
	Title string
	Link  string
}

func main() {
	template, err := template.ParseFiles("README.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://mpostument.com/index.xml")
	if err != nil {
		log.Fatalln(err)
	}

	postList := []ReadmeData{}
	for index, post := range feed.Items {
		if index == 7 {
			break
		}

		if strings.Contains(post.Link, "/projects/") {
			continue
		}

		readmeData := ReadmeData{
			Title: post.Title,
			Link:  post.Link,
		}
		postList = append(postList, readmeData)
	}

	f, err := os.Create("README.MD")
	if err != nil {
		log.Fatalln(err)
	}

	err = template.Execute(f, postList)
	if err != nil {
		log.Fatalln(err)
	}
}
