# `elem-ssg` Template

This template is designed for building a static site generator using Go, featuring `elem-go` for HTML templating and `goldmark` for Markdown processing.

To learn more about these two libraries, check out the following links:

- [elem-go](https://github.com/chasefleming/elem-go)
- [goldmark](https://github.com/yuin/goldmark)

For a deeper walkthrough of how to use this template, check out the [elem-ssg tutorial](https://dev.to/chasefleming/building-a-go-static-site-generator-using-elem-go-3fhh).

## Usage

To use this template, click the green "Use this template" button at the top of the page. This will create a new repository in your account with the contents of this template.

### Prerequisites

To use this template, you will need to have Go installed on your machine. You can download Go from the [official Go website](https://golang.org/dl/).

### Project Structure

```
your-project/
│
├── main.go             # Main Go file for the generator
├── posts/              # Directory for your Markdown posts
├── public/             # Directory for generated HTML files
└── go.mod              # Go module file
```

## Adding Content

To add content to your site, simply add Markdown files to the `posts/` directory. The generator will automatically convert these Markdown files to HTML files and place them in the `public/` directory.

To run the generator, use the following command:

```bash
go run main.go
```

## Viewing Your Site

To view your site, open the `public/index.html` file in your browser.

### Mac/Linux

```bash
open public/index.html
```

### Windows

```bash
start public/index.html
```
