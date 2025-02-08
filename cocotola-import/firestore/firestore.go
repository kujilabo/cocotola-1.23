package firestore

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Iterator interface {
	Next(ctx context.Context) ([]string, error)
	HasNext() bool
}

type CSVReader struct {
	reader *csv.Reader
	words  []string
	eof    bool
	err    error
}

func NewCSVReader(reader io.Reader) Iterator {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = '\t'
	csvReader.LazyQuotes = true

	eof := false
	words, err := csvReader.Read()
	if err == io.EOF {
		eof = true
	}
	return &CSVReader{
		reader: csvReader,
		words:  words,
		eof:    eof,
		err:    err,
	}
}

func (r *CSVReader) Next(ctx context.Context) ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}
	words := r.words
	err := r.err
	r.words, r.err = r.reader.Read()
	if errors.Is(r.err, io.EOF) {
		r.eof = true
	}
	return words, err
}

func (r *CSVReader) HasNext() bool {
	return !r.eof
}

func ProcessLine(ctx context.Context, itr Iterator, fn func([]string) error) error {
	for itr.HasNext() {
		words, err := itr.Next(ctx)
		if err != nil {
			return err
		}
		if err := fn(words); err != nil {
			return err
		}
	}
	return nil
}

func Firebase(ctx context.Context) error {
	// Firebase Develop 管理者
	// Firebase Develop Admin
	sa := option.WithCredentialsFile("../keys/firestore-importer.serviceAccount.json")
	// sa := option.WithCredentialsFile("../keys/firestore-adminsdk.serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	filePath := path.Join("data", "english-word-problem-001.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	itr := NewCSVReader(file)
	// r := regexp.MustCompile(`<([^>]*>`)
	if err := ProcessLine(ctx, itr, func(words []string) error {
		if len(words) != 3 {
			return fmt.Errorf("invalid words: %v", words)
		}
		id := words[0]
		// englishText := words[1]
		// translationText := words[2]

		_, err = client.Collection("englishWordProblem").Doc(id).Set(ctx, map[string]interface{}{
			"id":         "sampleId",
			"userName":   "go master",
			"email":      "go.master@exsample.com",
			"technology": "bbb",
		})

		log.Println(words)
		return nil
	}); err != nil {
		return err
	}

	// _, err = client.Collection("users").Doc("aaa").Set(ctx, map[string]interface{}{
	// 	"id":         "sampleId",
	// 	"userName":   "go master",
	// 	"email":      "go.master@exsample.com",
	// 	"technology": "bbb",
	// })

	// // データ追加
	// docs, err := client.Collection("users").Documents(ctx).GetAll()
	// if err != nil {
	// 	return err
	// }
	// for _, doc := range docs {
	// 	log.Println(doc.Data())
	// }
	return nil
}
