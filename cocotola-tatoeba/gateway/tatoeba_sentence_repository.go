package gateway

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"gorm.io/gorm"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

const (
	shuffleBufferRate = 10
)

type tatoebaSentenceEntity struct {
	SentenceNumber int
	Lang3          string
	Text           string
	Author         string
	UpdatedAt      time.Time
}

type tatoebaSentencePairEntity struct {
	SrcSentenceNumber int
	SrcLang3          string
	SrcText           string
	SrcAuthor         string
	SrcUpdatedAt      time.Time
	DstSentenceNumber int
	DstLang3          string
	DstText           string
	DstAuthor         string
	DstUpdatedAt      time.Time
}

func (e *tatoebaSentenceEntity) toModel() (*service.TatoebaSentence, error) {
	lang3, err := libdomain.NewLang3(e.Lang3)
	if err != nil {
		return nil, rsliberrors.Errorf("failed to NewLang3. err: %w", err)
	}
	author := e.Author
	if author == "\\N" {
		author = ""
	}
	return service.NewTatoebaSentence(e.SentenceNumber, lang3, e.Text, author, e.UpdatedAt)
}

func (e *tatoebaSentencePairEntity) toModel() (*service.TatoebaSentencePair, error) {
	srcE := tatoebaSentenceEntity{
		SentenceNumber: e.SrcSentenceNumber,
		Lang3:          e.SrcLang3,
		Text:           e.SrcText,
		Author:         e.SrcAuthor,
		UpdatedAt:      e.SrcUpdatedAt,
	}
	srcM, err := srcE.toModel()
	if err != nil {
		return nil, err
	}

	dstE := tatoebaSentenceEntity{
		SentenceNumber: e.DstSentenceNumber,
		Lang3:          e.DstLang3,
		Text:           e.DstText,
		Author:         e.DstAuthor,
		UpdatedAt:      e.DstUpdatedAt,
	}
	dstM, err := dstE.toModel()
	if err != nil {
		return nil, err
	}

	return service.NewTatoebaSentencePair(srcM, dstM)
}

func (e *tatoebaSentenceEntity) TableName() string {
	return "tatoeba_sentence"
}

type tatoebaSentenceRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func newTatoebaSentenceRepository(db *gorm.DB) service.TatoebaSentenceRepository {
	if db == nil {
		panic(errors.New(""))
	}

	return &tatoebaSentenceRepository{
		db:     db,
		logger: slog.Default().With(slog.String(rsliblog.LoggerNameKey, "TatoebaSentenceAddParameterReader")),
	}
}

// func (r *tatoebaSentenceRepository) FindTatoebaSentences(ctx context.Context, param domain.TatoebaSentenceSearchCondition) (*domain.TatoebaSentenceSearchResult, error) {
// 	logger := log.FromContext(ctx)
// 	logger.Debug("tatoebaSentenceRepository.FindTatoebaSentences")
// 	limit := param.GetPageSize()
// 	offset := (param.GetPageNo() - 1) * param.GetPageSize()

// 	where := func() *gorm.DB {
// 		db := r.db.Where("lang3 = 'eng'")
// 		if param.GetKeyword() != "" {
// 			keyword := "%" + param.GetKeyword() + "%"
// 			db = db.Where("text like ?", keyword)
// 		}
// 		return db
// 	}

// 	entities := []tatoebaSentenceEntity{}
// 	if result := where().Limit(limit).Offset(offset).Find(&entities); result.Error != nil {
// 		return nil, result.Error
// 	}

// 	results := make([]domain.TatoebaSentence, 0)
// 	for _, e := range entities {
// 		m, err := e.toModel()
// 		if err != nil {
// 			return nil, err
// 		}
// 		results = append(results, m)
// 	}

// 	var count int64
// 	if result := where().Model(&azureTranslationEntity{}).Count(&count); result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &domain.TatoebaSentenceSearchResult{
// 		TotalCount: count,
// 		Results:    results,
// 	}, nil
// }
//SELECT *
// FROM development.tatoeba_sentence t1
// inner join development.tatoeba_link t2
// on t1.sentence_number= t2.`from`

// inner join development.tatoeba_sentence t3
// on t3.sentence_number= t2.`to`

// where t1.lang3='eng' and t3.lang3='jpn';

func (r *tatoebaSentenceRepository) whereFindTatoebaSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) *gorm.DB {
	db := r.db.Table("tatoeba_sentence AS T1").Select(
		// Src
		"T1.sentence_number AS src_sentence_number,"+
			"T1.lang3 AS src_lang3,"+
			"T1.text AS src_text,"+
			"T1.author AS src_author,"+
			"T1.updated_at AS src_updated_at,"+
			// Dst
			"T3.sentence_number AS dst_sentence_number,"+
			"T3.lang3 AS dst_lang3,"+
			"T3.text AS dst_text,"+
			"T3.author AS dst_author,"+
			"T3.updated_at AS dst_updated_at").
		Joins("INNER JOIN tatoeba_link AS T2 ON T1.sentence_number = T2.`src`").
		Joins("INNER JOIN tatoeba_sentence AS T3 ON T3.sentence_number = T2.`dst`").
		Where("T1.lang3 = ? AND T3.lang3 = ?", param.GetSrcLang2().ToLang3().String(), param.GetDstLang2().ToLang3().String())
	keywords := SplitString(param.GetKeyword(), ' ', '"')
	for _, keyword := range keywords {
		keyword1 := strings.ReplaceAll(keyword, "%", "\\%")
		keyword2 := "%" + keyword1 + "%"
		db = db.Where("T1.text like ?", keyword2)
	}
	return db
}

func (r *tatoebaSentenceRepository) CountTatoebaSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) (int, error) {
	var count int64 = 0
	if result := r.whereFindTatoebaSentencePairs(ctx, param).Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *tatoebaSentenceRepository) FindTatoebaSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) ([]*service.TatoebaSentencePair, error) {
	ctx, span := tracer.Start(ctx, "tatoebaSentenceRepository.FindTatoebaSentencePairs")
	defer span.End()

	// ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	// logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	r.logger.DebugContext(ctx, fmt.Sprintf("keyword: %s, random: %v", param.GetKeyword(), param.IsRandom()))
	if param.IsRandom() {
		return r.findTatoebaSentencesByRandom(ctx, param)
	}
	return r.findTatoebaSentences(ctx, param)
}

func (r *tatoebaSentenceRepository) findTatoebaSentences(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) ([]*service.TatoebaSentencePair, error) {
	// ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	// logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	r.logger.InfoContext(ctx, "tatoebaSentenceRepository.FindTatoebaSentences")
	limit := param.GetPageSize()
	offset := (param.GetPageNo() - 1) * param.GetPageSize()

	//db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})

	// 	FROM `sandbox` AS s
	// INNER JOIN (
	//   SELECT CEIL(RAND() * (SELECT MAX(`id`) FROM `sandbox`)) AS `id`
	// ) AS `tmp` ON s.id >= tmp.id
	// ORDER BY s.id

	r.logger.InfoContext(ctx, "tatoebaSentenceRepository.FindTatoebaSentences 222")
	entities := []tatoebaSentencePairEntity{}
	if result := r.whereFindTatoebaSentencePairs(ctx, param).Limit(limit).Offset(offset).Scan(&entities); result.Error != nil {
		return nil, result.Error
	}

	results := make([]*service.TatoebaSentencePair, len(entities))
	for i, e := range entities {
		m, err := e.toModel()
		if err != nil {
			return nil, err
		}
		results[i] = m
	}

	r.logger.InfoContext(ctx, "tatoebaSentenceRepository.FindTatoebaSentences 333")

	return results, nil
}

func SplitString(str string, space, quote rune) []string {
	quoted := false
	split := strings.FieldsFunc(str, func(r1 rune) bool {
		if r1 == quote {
			quoted = !quoted
		}
		return !quoted && r1 == space
	})
	for i := 0; i < len(split); i++ {
		split[i] = strings.Trim(split[i], string(quote))
	}
	return split
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (r *tatoebaSentenceRepository) findTatoebaSentencesByRandom(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) ([]*service.TatoebaSentencePair, error) {
	// ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	// logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	r.logger.Debug("tatoebaSentenceRepository.FindTatoebaSentences")
	limit := param.GetPageSize() * shuffleBufferRate
	offset := (param.GetPageNo() - 1) * param.GetPageSize()

	where := func() *gorm.DB {
		db := r.db.Table("tatoeba_sentence AS T1").Select(
			// Src
			"T1.sentence_number AS src_sentence_number," +
				"T1.lang3 AS src_lang3," +
				"T1.text AS src_text," +
				"T1.author AS src_author," +
				"T1.updated_at AS src_updated_at," +
				// Dst
				"T3.sentence_number AS dst_sentence_number," +
				"T3.lang3 AS dst_lang3," +
				"T3.text AS dst_text," +
				"T3.author AS dst_author," +
				"T3.updated_at AS dst_updated_at").
			Joins("INNER JOIN tatoeba_link AS T2 ON T1.sentence_number = T2.`src`").
			Joins("INNER JOIN tatoeba_sentence AS T3 ON T3.sentence_number = T2.`dst`").
			Joins("INNER JOIN (SELECT CEIL(RAND() * (SELECT MAX(`sentence_number`) FROM `tatoeba_sentence`)) AS `sentence_number`) AS `tmp` ON T1.sentence_number >= tmp.sentence_number").
			Where("T1.lang3 = 'eng' AND T3.lang3 = 'jpn'")
		if param.GetKeyword() != "" {
			keyword1 := strings.ReplaceAll(param.GetKeyword(), "%", "\\%")
			keyword2 := "%" + keyword1 + "%"
			db = db.Where("T1.text like ?", keyword2)
		}
		return db
	}

	entities := []tatoebaSentencePairEntity{}
	if result := where().Limit(limit).Offset(offset).Scan(&entities); result.Error != nil {
		return nil, result.Error
	}

	// rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(entities), func(i, j int) { entities[i], entities[j] = entities[j], entities[i] })

	r.logger.InfoContext(ctx, fmt.Sprintf("len(entities): %d", len(entities)))

	length := min(param.GetPageSize(), len(entities))
	results := make([]*service.TatoebaSentencePair, length)
	for i := 0; i < length; i++ {
		m, err := entities[i].toModel()
		if err != nil {
			return nil, err
		}
		results[i] = m
	}

	// var count int64 = 0
	// // if result := where().Count(&count); result.Error != nil {
	// // 	return nil, result.Error
	// // }

	return results, nil
}

func (r *tatoebaSentenceRepository) FindTatoebaSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (*service.TatoebaSentence, error) {
	entity := tatoebaSentenceEntity{}
	if result := r.db.Where("sentence_number = ?", sentenceNumber).
		First(&entity); result.Error != nil {
		return nil, result.Error
	}

	sentence, err := entity.toModel()
	if err != nil {
		return nil, err
	}

	return sentence, nil
}

func (r *tatoebaSentenceRepository) ContainsSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (bool, error) {
	entity := tatoebaSentenceEntity{}
	if result := r.db.Where("sentence_number = ?", sentenceNumber).
		First(&entity); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (r *tatoebaSentenceRepository) Add(ctx context.Context, param service.TatoebaSentenceAddParameterInterface) error {
	entity := tatoebaSentenceEntity{
		SentenceNumber: param.GetSentenceNumber(),
		Lang3:          param.GetLang3().String(),
		Text:           param.GetText(),
		Author:         param.GetAuthor(),
		UpdatedAt:      param.GetUpdatedAt(),
	}

	if result := r.db.Create(&entity); result.Error != nil {
		err := rslibgateway.ConvertDuplicatedError(result.Error, service.ErrTatoebaSentenceAlreadyExists)
		return rsliberrors.Errorf("failed to Add tatoebaSentence. err: %w", err)
	}

	return nil
}
