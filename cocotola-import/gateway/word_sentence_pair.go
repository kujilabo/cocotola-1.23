package gateway

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/service"
)

type WordSentencePairEntity struct {
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

func (u *WordSentencePairEntity) TableName() string {
	return "word_sentence_pair"
}

type wordSentencePairRepository struct {
	db *gorm.DB
}

func NewWordSentencePairRepository(db *gorm.DB) service.WordSentencePairRepository {
	return &wordSentencePairRepository{
		db: db,
	}
}

func (r *wordSentencePairRepository) ContainsWorkSentencePairBySentenceNumber(ctx context.Context, srcSentenceNumber, dstSentenceNumber int) (bool, error) {
	wordSentencePairEntity := WordSentencePairEntity{}
	if result := r.db.Where("src_sentence_number = ? and dst_sentence_number = ?", srcSentenceNumber, dstSentenceNumber).First(&wordSentencePairEntity); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, rsliberrors.Errorf("Select: %w", result.Error)
	}

	return true, nil
}

func (r *wordSentencePairRepository) AddWorkSentencePair(ctx context.Context, param *service.WordSentencePairAddParameter) (int, error) {
	wordSentencePairEntity := WordSentencePairEntity{
		WorkbookID:        param.WorkbookID,
		DocumentID:        "",
		SrcSentenceNumber: param.Src.SentenceNumber,
		SrcLang2:          param.Src.Lang2,
		SrcText:           param.Src.Text,
		SrcAuthor:         param.Src.Author,
		DstSentenceNumber: param.Dst.SentenceNumber,
		DstLang2:          param.Dst.Lang2,
		DstText:           param.Dst.Text,
		DstAuthor:         param.Dst.Author,
		CreatedAt:         param.CreatedAt,
		UpdatedAt:         param.UpdatedAt,
	}
	if result := r.db.Create(&wordSentencePairEntity); result.Error != nil {
		return 0, rsliberrors.Errorf("Create: %w", result.Error)
	}
	return wordSentencePairEntity.ID, nil
}

// func (r *wordSentencePairRepository) FindWordSentencePairByID() (*domain.WordSentencePair, error) {
// 	r.db
// }

func (r *wordSentencePairRepository) UpdateWordSentencePairDocumentID(ctx context.Context, workbookID int, documentID string) error {
	if result := r.db.Model(&WordSentencePairEntity{}).Where("id = ?",
		workbookID).Update("document_id", documentID); result.Error != nil {
		return rsliberrors.Errorf("Update: %w", result.Error)
	}
	return nil
}
