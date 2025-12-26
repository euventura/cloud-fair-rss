package services

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title   string
	Sources []Source
}

type Source struct {
	URL  string
	Star bool
}

type Entry struct {
	Star        bool
	Title       string
	Url         string
	Link        string
	Author      string
	Content     template.HTML
	Description string
	Date        time.Time
	Class       string
	ID          string
}

var artPath = "/template/article.html"
var hePath = "/template/headline.html"
var stPath = "/../static/"

func newFeed() *Feed {
	return &Feed{
		Title:   "RSS Zombie",
		Sources: []Source{},
	}

}

func (f *Feed) loadSources() {

	file, err := os.Open("../sources.txt")

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			star := false
			lastChar := line[len(line)-1:]
			if lastChar == "*" {
				line = line[:len(line)-1]
				star = true
			}
			source := Source{
				URL:  line,
				Star: star,
			}
			f.Sources = append(f.Sources, source)
		}
		return
	}

	r2, _ := NewR2()

	content, _ := r2.Get("sources.txt")

	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			star := false
			lastChar := line[len(line)-1:]
			if lastChar == "*" {
				line = line[:len(line)-1]
				star = true
			}
			source := Source{
				URL:  line,
				Star: star,
			}
			f.Sources = append(f.Sources, source)
		}

	}
}

func (f *Feed) fetch() {
	f.loadSources()

	for _, source := range f.Sources {
		fmt.Println("Source:", source.URL, "Starred:", source.Star)
		fp := gofeed.NewParser()
		fp.UserAgent = "CloudFair 0.1"

		feed, _ := fp.ParseURL(source.URL)
		f.process(feed, source.Star)
	}
}

func dir() string {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	return dir
}

func (f *Feed) process(gof *gofeed.Feed, star bool) {

	wPath := dir() + stPath + time.Now().Format("02012006") + "/"
	err := os.MkdirAll(wPath, os.ModePerm)

	var headline string
	var fiName string
	fmt.Println("Processing feed:", gof.Title)
	fmt.Println("Items: ", len(gof.Items))

	for _, item := range gof.Items {

		// if item.PublishedParsed.Format("02012006") != time.Now().Format("02012006") {
		// 	fmt.Println(item.PublishedParsed.Format("02/01/2006"))
		// 	continue
		// }

		//
		//Url := time.Now().Format("02-01-2006") + "/" + gof.Title + ".html"
		author := item.Authors[0].Name

		if author == "" {
			url, _ := url.Parse(item.Link)
			author = url.Hostname()
		}

		const regex = `<.*?>`
		r := regexp.MustCompile(regex)
		desc := r.ReplaceAllString(item.Description, "")
		words := strings.Fields(desc)

		class := slug.Make(author)

		data := Entry{
			Star:        star,
			Title:       item.Title,
			Link:        item.Link,
			Url:         item.Link,
			Author:      author,
			Content:     template.HTML(item.Content),
			Date:        *item.PublishedParsed,
			Description: strings.Join(words[0:min(38, len(words))], " "),
			Class:       class,
			ID:          slug.Make(item.Title),
		}
		fiName = slug.Make(item.Title) + ".html"
		fmt.Println(fiName)
		article := f.make(data, dir()+artPath)
		err = os.WriteFile(wPath+"/"+fiName, []byte(article), 0644)
		if err != nil {
			fmt.Println("Erro ao escrever arquivo:", err)
			return
		}

		headline += f.make(data, dir()+hePath)

	}

	hFi, err := os.OpenFile(wPath+"/index.html", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo de índice:", err)
		return
	}
	defer hFi.Close()

	if _, err := hFi.WriteString(headline); err != nil {
		fmt.Println("Erro ao escrever no headline de índice:", err)
		return
	}

}

func (f *Feed) make(data Entry, templatePath string) string {

	tmp, err := os.ReadFile(templatePath)

	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}

	t, err := template.New(templatePath).Parse(string(tmp))
	err = t.ExecuteTemplate(buf, templatePath, data)

	if err != nil {
		panic(err)
	}

	return buf.String()
}
