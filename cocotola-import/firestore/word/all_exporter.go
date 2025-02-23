package word

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"gorm.io/gorm"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

func ExportAllWords(ctx context.Context, client *firestore.Client, latestExportedAt time.Time) error {
	// documents, err := client.Collection("word_sentence_pair").Where("updated_at", "<", latestExportedAt.Format(time.RFC3339)).Documents(ctx).GetAll()
	documents, err := client.Collection("word_sentence_pair").Where("updatedAt", ">", latestExportedAt).Documents(ctx).GetAll()
	if err != nil {
		return rsliberrors.Errorf("client.Collection.Documents.GetAll err: %w", err)
	}

	logger := slog.Default()
	logger.InfoContext(ctx, fmt.Sprintf("exportAllwords : %d", len(documents)))
	for _, doc := range documents {
		data := doc.Data()
		logger.InfoContext(ctx, fmt.Sprintf("data: %+v", data))
	}
	return nil
}

func exportAllWords(ctx context.Context, client *firestore.Client, db *gorm.DB) error {

	return nil
}

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

type WordSentencePairEntity struct {
	ID                int
	WorkbookID        int
	SrcSentenceNumber int
	SrcLang2          string
	SrcText           string
	SrcAuthor         string
	DstSentenceNumber int
	DstLang2          string
	DstText           string
	DstAuthor         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (u *WordSentencePairEntity) TableName() string {
	return "word_sentence_pair"
}
func ImportNewWordSentencePairs(ctx context.Context, client *firestore.Client, db *gorm.DB, jsonFile io.Reader) error {
	decoder := json.NewDecoder(jsonFile)

	sentencePairs := NewWordSentencePairs{}
	if err := decoder.Decode(&sentencePairs.List); err != nil {
		return rsliberrors.Errorf("json.Decode err: %w", err)
	}

	logger := slog.Default()
	logger.InfoContext(ctx, fmt.Sprintf("sentencePairs: %+v", sentencePairs))

	sentencePairsToAdd := make(map[string]NewWordSentencePair)
	for key, value := range sentencePairs.List {
		keys := strings.Split(key, "-")
		if len(keys) != 2 {
			return rsliberrors.Errorf("key is invalid: %s", key)
		}
		srcSentenceNumber := keys[0]
		dstSentenceNumber := keys[1]
		wordSentencePairEntity := WordSentencePairEntity{}
		if result := db.Where("src_sentence_number = ? and dst_sentence_number = ?", srcSentenceNumber, dstSentenceNumber).First(&wordSentencePairEntity); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				sentencePairsToAdd[key] = value
			} else {
				return rsliberrors.Errorf("Select: %w", result.Error)
			}
		}
	}

	for _, value := range sentencePairsToAdd {
		wordSentencePairEntity := WordSentencePairEntity{
			WorkbookID:        1,
			SrcSentenceNumber: value.Src.SentenceNumber,
			SrcLang2:          value.Src.Lang2,
			SrcText:           value.Src.Text,
			SrcAuthor:         value.Src.Author,
			DstSentenceNumber: value.Dst.SentenceNumber,
			DstLang2:          value.Dst.Lang2,
			DstText:           value.Dst.Text,
			DstAuthor:         value.Dst.Author,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}
		if result := db.Create(&wordSentencePairEntity); result.Error != nil {
			return rsliberrors.Errorf("Create: %w", result.Error)
		}

		logger.InfoContext(ctx, fmt.Sprintf("ID: %+v", wordSentencePairEntity.ID))

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
			"createdAt": wordSentencePairEntity.CreatedAt,
			"updatedAt": wordSentencePairEntity.UpdatedAt,
		})
		if err != nil {
			return rsliberrors.Errorf("client.Collection.Add err: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("doc: %v", doc))
		logger.InfoContext(ctx, fmt.Sprintf("result: %v", result))
		if result := db.Model(&WordSentencePairEntity{}).Where("id = ?",
			wordSentencePairEntity.ID).Update("document_id", doc.ID); result.Error != nil {
			return rsliberrors.Errorf("Update: %w", result.Error)
		}
	}

	return nil
}
