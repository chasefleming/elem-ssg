package main

import (
	"bytes"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/yuin/goldmark"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func layout(title string, content elem.Node) string {
	htmlPage := elem.Html(nil,
		elem.Head(nil,
			elem.Title(nil, elem.Text(title)),
		),
		elem.Body(nil,
			elem.Header(nil, elem.H1(nil, elem.Text(title))),
			elem.Main(nil, content),
			elem.Footer(nil, elem.Text("Footer content here")),
		),
	)

	return htmlPage.Render()
}

func createHTMLPage(title, content string) string {
	htmlOutput := layout(title, elem.Raw(content))

	postFilename := title + ".html"
	filepath := filepath.Join("public", postFilename)
	os.WriteFile(filepath, []byte(htmlOutput), 0644)
	return postFilename
}

func markdownToHTML(content string) string {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(content), &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func readMarkdownPosts(dir string) []string {
	var posts []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			htmlContent := markdownToHTML(string(content))
			title := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			postFilename := createHTMLPage(title, htmlContent)

			posts = append(posts, postFilename)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func createIndexPage(postFilenames []string) {
	listItems := make([]elem.Node, len(postFilenames))
	for i, filename := range postFilenames {
		link := elem.A(attrs.Props{attrs.Href: "./" + filename}, elem.Text(filename))
		listItems[i] = elem.Li(nil, link)
	}

	indexContent := elem.Ul(nil, listItems...)
	htmlOutput := layout("Home", indexContent)

	filepath := filepath.Join("public", "index.html")
	os.WriteFile(filepath, []byte(htmlOutput), 0644)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755) // or 0700 if you need it to be private
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	createDirIfNotExist("posts")
	createDirIfNotExist("public")

	posts := readMarkdownPosts("posts")
	createIndexPage(posts)
}
