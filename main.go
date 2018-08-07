package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/google/go-github/github"
	"github.com/micnncim/medicine/config"
	"github.com/micnncim/medicine/gist"
	"github.com/micnncim/medicine/markdown"
	"github.com/micnncim/medicine/medium"
)

const version = "0.4.0"

var cnf config.Config

func init() {
	if err := cnf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	g, err := gist.New(cnf.GistConfig.Token)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) != 2 {
		log.Fatal(errors.New("invalid args"))
	}

	md, err := markdown.New(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if err := md.ParseTitle(); err != nil {
		log.Fatal(err)
	}
	if err := md.ParseSnippets(); err != nil {
		log.Fatal(err)
	}

	var urls []string
	var wg sync.WaitGroup
	for _, snippet := range md.Snippets {
		wg.Add(1)
		go func(s *markdown.Snippet) {
			defer wg.Done()
			files := map[github.GistFilename]github.GistFile{
				github.GistFilename(s.Filename): github.GistFile{
					Content: &s.Content,
				},
			}
			item, err := g.Create(context.Background(), files, "", true)
			if err != nil {
				log.Fatal(err)
			}
			urls = append(urls, *item.HTMLURL)
			fmt.Println(*item.HTMLURL)
		}(snippet)
		wg.Wait()
	}

	if err := md.Replace(urls...); err != nil {
		log.Fatal(err)
	}

	med := medium.New(cnf.MediumConfig.Token)
	if err := med.Publish(md); err != nil {
		log.Fatal(err)
	}
}
