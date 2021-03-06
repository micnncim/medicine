package markdown

import (
	"errors"
	"io/ioutil"
	"strings"
)

type (
	Markdown struct {
		Title    string
		Filename string
		Content  string
		Snippets []*Snippet
	}

	Snippet struct {
		Filename string
		Content  string
	}
)

func New(filename string) (*Markdown, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Markdown{
		Filename: filename,
		Content:  string(data),
	}, nil
}

func (m *Markdown) ParseSnippets() error {
	content := strings.Split(m.Content, "\n")
	for i := range content {
		line := string(content[i])
		if !strings.HasPrefix(line, "```") {
			continue
		}
		filename := strings.Trim(line, "`")
		if filename == "" {
			continue
		}

		var n int
		for {
			if strings.HasPrefix(string(content[i+n+1]), "```") {
				break
			}
			n++
		}
		m.Snippets = append(m.Snippets, &Snippet{
			Filename: filename,
			Content:  strings.Join(content[i+1:i+n+1], "\n"),
		})
	}

	if len(m.Snippets) == 0 {
		return errors.New("code blocks not exist")
	}
	return nil
}

func (m *Markdown) ParseTitle() error {
	if m.Content == "" {
		return errors.New("file content not exist")
	}

	content := strings.Split(m.Content, "\n")
	header := content[0]
	if strings.HasPrefix(header, "# ") {
		m.Title = strings.Trim(header, "# ")
	} else {
		m.Title = header
	}

	return nil
}

func (m *Markdown) Replace(urls ...string) error {
	if len(m.Snippets) == 0 {
		return errors.New("code blocks not exist")
	}
	if len(urls) != len(m.Snippets) {
		return errors.New("the number of URLs not match that of code blocks")
	}

	for i, s := range m.Snippets {
		block := "```" + s.Filename + "\n" + s.Content + "\n```"
		m.Content = strings.Replace(m.Content, block, urls[i], 1)
	}

	return nil
}
