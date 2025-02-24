package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/kujilabo/cocotola-1.23/lib/domain"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
)

var ErrTatoebaSentenceAlreadyExists = errors.New("tatoebaSentence already exists")
var ErrTatoebaSentenceNotFound = errors.New("tatoebaSentence not found")

// type TatoebaSentence interface {
// 	GetSentenceNumber() int
// 	GetLang3() *libdomain.Lang3
// 	GetText() string
// 	GetAuthor() string
// 	GetUpdatedAt() time.Time
// }

type TatoebaSentence struct {
	SentenceNumber int
	Lang3          *libdomain.Lang3
	Text           string
	Author         string
	UpdatedAt      time.Time
}

func NewTatoebaSentence(sentenceNumber int, lang3 *libdomain.Lang3, text, author string, updatedAt time.Time) (*TatoebaSentence, error) {
	m := &TatoebaSentence{
		SentenceNumber: sentenceNumber,
		Lang3:          lang3,
		Text:           text,
		Author:         author,
		UpdatedAt:      updatedAt,
	}

	return m, rslibdomain.Validator.Struct(m)
}

// func (m *tatoebaSentence) GetSentenceNumber() int {
// 	return m.SentenceNumber
// }

// func (m *tatoebaSentence) GetLang3() *libdomain.Lang3 {
// 	return m.Lang3
// }

// func (m *tatoebaSentence) GetText() string {
// 	return m.Text
// }

// func (m *tatoebaSentence) GetAuthor() string {
// 	return m.Author
// }

// func (m *tatoebaSentence) GetUpdatedAt() time.Time {
// 	return m.UpdatedAt
// }

// type TatoebaSentencePair interface {
// 	GetSrc() TatoebaSentence
// 	GetDst() TatoebaSentence
// }

type TatoebaSentencePair struct {
	Src *TatoebaSentence
	Dst *TatoebaSentence
}

func NewTatoebaSentencePair(src, dst *TatoebaSentence) (*TatoebaSentencePair, error) {
	m := &TatoebaSentencePair{
		Src: src,
		Dst: dst,
	}

	return m, rslibdomain.Validator.Struct(m)
}

// func (m *tatoebaSentencePair) GetSrc() TatoebaSentence {
// 	return m.Src
// }

// func (m *tatoebaSentencePair) GetDst() TatoebaSentence {
// 	return m.Dst
// }

type TatoebaSentenceAddParameterInterface interface {
	GetSentenceNumber() int
	GetLang3() *libdomain.Lang3
	GetText() string
	GetAuthor() string
	GetUpdatedAt() time.Time
}

type TatoebaSentenceAddParameter struct {
	SentenceNumber int `validate:"required"`
	Lang3          *libdomain.Lang3
	Text           string `validate:"required"`
	Author         string `validate:"required"`
	UpdatedAt      time.Time
}

func NewTatoebaSentenceAddParameter(sentenceNumber int, lang3 *libdomain.Lang3, text, author string, updatedAt time.Time) (*TatoebaSentenceAddParameter, error) {
	m := &TatoebaSentenceAddParameter{
		SentenceNumber: sentenceNumber,
		Lang3:          lang3,
		Text:           text,
		Author:         author,
		UpdatedAt:      updatedAt,
	}

	return m, rslibdomain.Validator.Struct(m)
}

func (p *TatoebaSentenceAddParameter) GetSentenceNumber() int {
	return p.SentenceNumber
}

func (p *TatoebaSentenceAddParameter) GetLang3() *libdomain.Lang3 {
	return p.Lang3
}

func (p *TatoebaSentenceAddParameter) GetText() string {
	return p.Text
}

func (p *TatoebaSentenceAddParameter) GetAuthor() string {
	return p.Author
}

func (p *TatoebaSentenceAddParameter) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

type TatoebaSentenceSearchConditionInterface interface {
	GetPageNo() int
	GetPageSize() int
	GetSrcLang2() *domain.Lang2
	GetDstLang2() *domain.Lang2
	GetKeyword() string
	IsRandom() bool
	ToString() string
}

type TatoebaSentenceSearchCondition struct {
	PageNo   int           `validate:"required,gte=1"`
	PageSize int           `validate:"required,gte=1,lte=100"`
	SrcLang2 *domain.Lang2 `validate:"required"`
	DstLang2 *domain.Lang2 `validate:"required"`
	Keyword  string
	Random   bool
	JSON     string
}

func NewTatoebaSentenceSearchCondition(pageNo, pageSize int, srcLang2, dstLang2 *domain.Lang2, keyword string, random bool) (*TatoebaSentenceSearchCondition, error) {
	m := &TatoebaSentenceSearchCondition{
		PageNo:   pageNo,
		PageSize: pageSize,
		SrcLang2: srcLang2,
		DstLang2: dstLang2,
		Keyword:  keyword,
		Random:   random,
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	m.JSON = string(bytes)

	return m, rslibdomain.Validator.Struct(m)
}

func (c *TatoebaSentenceSearchCondition) GetPageNo() int {
	return c.PageNo
}

func (c *TatoebaSentenceSearchCondition) GetPageSize() int {
	return c.PageSize
}

func (c *TatoebaSentenceSearchCondition) GetKeyword() string {
	return c.Keyword
}

func (c *TatoebaSentenceSearchCondition) GetSrcLang2() *domain.Lang2 {
	return c.SrcLang2
}

func (c *TatoebaSentenceSearchCondition) GetDstLang2() *domain.Lang2 {
	return c.DstLang2
}

func (c *TatoebaSentenceSearchCondition) IsRandom() bool {
	return c.Random
}
func (c *TatoebaSentenceSearchCondition) ToString() string {
	return c.JSON
}

type TatoebaSentencePairSearchResult struct {
	TotalCount int
	Results    []*TatoebaSentencePair
}

func NewTatoebaSentencePairSearchResult(totalCount int, results []*TatoebaSentencePair) *TatoebaSentencePairSearchResult {
	return &TatoebaSentencePairSearchResult{
		TotalCount: totalCount,
		Results:    results,
	}
}

type TatoebaSentenceRepository interface {
	FindTatoebaSentencePairs(ctx context.Context, param TatoebaSentenceSearchConditionInterface) ([]*TatoebaSentencePair, error)
	CountTatoebaSentencePairs(ctx context.Context, param TatoebaSentenceSearchConditionInterface) (int, error)

	FindTatoebaSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (*TatoebaSentence, error)

	Add(ctx context.Context, param TatoebaSentenceAddParameterInterface) error

	ContainsSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (bool, error)
}
