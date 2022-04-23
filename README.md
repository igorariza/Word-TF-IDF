# Challenge Word-TF-IDF

CLI tool in GO that receives a word and a text document and returns the TF-IDF weight of the word in the document.

## Implementation Guide

- Search word
```
//successful case
go run cmd/tfidf-word/main.go -word "enim"
go run cmd/tfidf-word/main.go -word "Nunc"
go run cmd/tfidf-word/main.go -word "scelerisque"

//case for all words
go run cmd/tfidf-word/main.go

//case word not found
go run cmd/tfidf-word/main.go -word "foo"

```

- URL-TF-IDF

```
//case receive url
go run cmd/tfidf-url/main.go

```
