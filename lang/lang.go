package lang

import (
	"strings"
)

// New returns a new and initialized Lang.
func New(language string) Lang {
	l := Lang{language: strings.ToLower(language)} // case insensitivity

	// Add all lanaguages here, the keys should be named according to BCP47.
	// The keys must be in all lower case for normalized lookup.
	l.m = map[string]Term{
		"en": {
			Footnotes:    "Footnotes",
			Bibliography: "Bibliography",
			Index:        "Index",
		},
		"nl": {
			Footnotes:    "Voetnoten",
			Bibliography: "Bibliografie",
			Index:        "Index",
		},
		"de": {
			Footnotes:    "Fußnoten",
			Bibliography: "Literaturverzeichnis",
			Index:        "Index",
		},
		"ja": {
			Footnotes:    "脚注",
			Bibliography: "参考文献",
			Index:        "索引",
		},
		"zh-cn": {
			Footnotes:    "注释",
			Bibliography: "参考文献",
			Index:        "索引",
		},
		"zh-tw": {
			Footnotes:    "註釋",
			Bibliography: "參考文獻",
			Index:        "索引",
		},
	}

	return l
}

// Lang maps a language to the terms we use in the document. We use an 'int' as to use the parser.Flags
// to indicate which language we'are using.
type Lang struct {
	language string
	m        map[string]Term
}

// Term contains the specific terms for translation.
type Term struct {
	Footnotes    string
	Bibliography string
	Index        string
}

func (l Lang) Footnotes() string {
	t, ok := l.m[l.language]
	if !ok {
		return l.m["en"].Footnotes
	}
	return t.Footnotes
}

func (l Lang) Bibliography() string {
	t, ok := l.m[l.language]
	if !ok {
		return l.m["en"].Bibliography
	}
	return t.Bibliography
}

func (l Lang) Index() string {
	t, ok := l.m[l.language]
	if !ok {
		return l.m["en"].Index
	}
	return t.Index
}
