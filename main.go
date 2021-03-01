package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/mmcdole/gofeed"
)

type Profile struct {
	PostsList map[string]string
}

var postsList = make(map[string]string)

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

	for index, post := range feed.Items {
		if index == 10 {
			break
		}

		if strings.Contains(post.Link, "/projects/") {
			continue
		}

		postsList[post.Title] = post.Link
	}
	profile := Profile{
		PostsList: postsList,
	}

	f, err := os.Create("README.MD")
	if err != nil {
		log.Fatalln(err)
	}

	err = template.Execute(f, profile)
	if err != nil {
		log.Fatalln(err)
	}
}
