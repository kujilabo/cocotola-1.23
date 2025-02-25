package import_firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"gorm.io/gorm"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/service"
)

type WordSentence struct {
	SentenceNumber int    `json:"sentenceNumber"`
	Lang2          string `json:"lang2"`
	Text           string `json:"text"`
	Author         string `json:"author"`
}
type NewWordSentencePair struct {
	Src WordSentence `json:"src"`
	Dst WordSentence `json:"dst"`
}

type NewWordSentencePairs struct {
	List map[string]NewWordSentencePair
}

func toSrcDstSentenceNumber(key string) (int, int, error) {
	keys := strings.Split(key, "-")
	if len(keys) != 2 {
		return 0, 0, rsliberrors.Errorf("key is invalid: %s", key)
	}

	srcSentenceNumber, err := strconv.Atoi(keys[0])
	if err != nil {
		return 0, 0, rsliberrors.Errorf("strconv.Atoi: %w", err)
	}
	dstSentenceNumber, err := strconv.Atoi(keys[1])
	if err != nil {
		return 0, 0, rsliberrors.Errorf("strconv.Atoi: %w", err)
	}

	return srcSentenceNumber, dstSentenceNumber, nil
}

func ImportNewWordSentencePairs(ctx context.Context, client *firestore.Client, db *gorm.DB, jsonFile io.Reader) error {
	logger := slog.Default()

	sentencePairs := NewWordSentencePairs{}
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&sentencePairs.List); err != nil {
		return rsliberrors.Errorf("json.Decode err: %w", err)
	}

	wordSentencePairRepository := gateway.NewWordSentencePairRepository(db)

	sentencePairsToAdd := make(map[string]NewWordSentencePair)
	for key, value := range sentencePairs.List {
		srcSentenceNumber, dstSentenceNumber, err := toSrcDstSentenceNumber(key)
		if err != nil {
			return rsliberrors.Errorf("toSrcDstSentenceNumber: %w", err)
		}

		contains, err := wordSentencePairRepository.ContainsWorkSentencePairBySentenceNumber(ctx, srcSentenceNumber, dstSentenceNumber)
		if err != nil {
			return rsliberrors.Errorf("ContainsWorkSentencePairBySentenceNumber: %w", err)
		}

		if !contains {
			sentencePairsToAdd[key] = value
		}
	}

	for _, value := range sentencePairsToAdd {
		wordSentencePairAddParameter := service.WordSentencePairAddParameter{
			WorkbookID: 1,
			Src: service.WordSentence{
				SentenceNumber: value.Src.SentenceNumber,
				Lang2:          value.Src.Lang2,
				Text:           value.Src.Text,
				Author:         value.Src.Author,
			},
			Dst: service.WordSentence{
				SentenceNumber: value.Dst.SentenceNumber,
				Lang2:          value.Dst.Lang2,
				Text:           value.Dst.Text,
				Author:         value.Dst.Author,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		wordSentencePairID, err := wordSentencePairRepository.AddWorkSentencePair(ctx, &wordSentencePairAddParameter)
		if err != nil {
			return rsliberrors.Errorf("Create: %w", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("ID: %+v", wordSentencePairID))

		doc, result, err := client.Collection("word_sentence_pair").Add(ctx, map[string]interface{}{
			"src": map[string]interface{}{
				"sentenceNumber": value.Src.SentenceNumber,
				"lang2":          value.Src.Lang2,
				"text":           value.Src.Text,
				"author":         value.Src.Author,
			},
			"dst": map[string]interface{}{
				"sentenceNumber": value.Dst.SentenceNumber,
				"lang2":          value.Dst.Lang2,
				"text":           value.Dst.Text,
				"author":         value.Dst.Author,
			},
			"createdAt": wordSentencePairAddParameter.CreatedAt,
			"updatedAt": wordSentencePairAddParameter.UpdatedAt,
		})
		if err != nil {
			return rsliberrors.Errorf("client.Collection.Adderr: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("doc: %v", doc))
		logger.InfoContext(ctx, fmt.Sprintf("result: %v", result))
		if err := wordSentencePairRepository.UpdateWordSentencePairDocumentID(ctx, wordSentencePairID, doc.ID); err != nil {
			return rsliberrors.Errorf("Update: %w", err)
		}
	}

	return nil
}
