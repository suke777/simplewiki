package wiki

import (
	"io/ioutil"
	"strings"
	"text/template"
)

// PageData １ページのデータ
type PageData struct {
	Title string
	Body  []byte
}

func loadPageData(title string) (*PageData, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("db/" + filename)
	if err != nil {
		return nil, err
	}
	return &PageData{title, body}, nil
}

func savePageData(title string, body []byte) error {
	filename := title + ".txt"
	return ioutil.WriteFile("db/"+filename, body, 0600)
}

func nl2Br(text []byte) []byte {
	safe := template.HTMLEscapeString(string(text))
	replacedText := strings.Replace(safe, "\n", "<br>", -1)
	return []byte(replacedText)
}
