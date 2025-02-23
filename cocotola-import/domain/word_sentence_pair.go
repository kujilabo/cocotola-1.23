package domain

import "time"

type WordSentencePair struct {
	ID                int
	WorkbookID        int
	DocumentID        string
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
