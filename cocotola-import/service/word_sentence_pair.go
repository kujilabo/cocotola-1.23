package service

import (
	"context"
	"time"
)

type WordSentence struct {
	SentenceNumber int
	Lang2          string
	Text           string
	Author         string
}
type WordSentencePairAddParameter struct {
	WorkbookID int
	Src        WordSentence
	Dst        WordSentence
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type WordSentencePairRepository interface {
	ContainsWorkSentencePairBySentenceNumber(ctx context.Context, srcSentenceNumber, dstSentenceNumber int) (bool, error)
	AddWorkSentencePair(ctx context.Context, param *WordSentencePairAddParameter) (int, error)
	UpdateWordSentencePairDocumentID(ctx context.Context, workbookID int, documentID string) error
}
