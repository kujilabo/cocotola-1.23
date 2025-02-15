package gateway

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
	"time"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

const (
	bufSize         = 4096
	textLimitLength = 100
)

type TatoebaSentenceAddParameterReader struct {
	// reader *csv.Reader
	reader *bufio.Reader
	num    int
	logger *slog.Logger
}

// type wrdomainedReader struct {
// 	reader io.Reader
// }

// func (r *wrappedReader) Read(p []byte) (int, error) {
// 	n, err := r.reader.Read(p)
// 	if err != nil {
// 		return 0, err
// 	}
// 	s1 := string(p)
// 	s2 := strings.ReplaceAll(s1, "\"", "\"\"")
// 	return n + len(s2) - len(s1), nil
// }

func NewTatoebaSentenceAddParameterReader(reader io.Reader) *TatoebaSentenceAddParameterReader {
	bufReader := bufio.NewReaderSize(reader, bufSize)
	// wrappedReader:=

	// csvReader := csv.NewReader(reader)
	// csvReader.Comma = '\t'
	// csvReader.LazyQuotes = true

	return &TatoebaSentenceAddParameterReader{
		// reader: csvReader,
		reader: bufReader,
		num:    1,
		logger: slog.Default().With(slog.String(rsliblog.LoggerNameKey, "TatoebaSentenceAddParameterReader")),
	}
}

func (r *TatoebaSentenceAddParameterReader) Next(ctx context.Context) (*service.TatoebaSentenceAddParameter, error) {
	// ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	// logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	b, _, err := r.reader.ReadLine()
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		r.logger.InfoContext(ctx, "zero")
	}
	line := strings.Split(string(b), "\t")
	// var line []string
	// line, err := r.reader.Read()
	// if errors.Is(err, io.EOF) {
	// 	return nil, err
	// }

	// if err != nil {
	// 	// skip
	// 	logger.InfoContext(ctx, fmt.Sprintf("skip rowNumber: %d, line: %v", r.num, line))
	// 	r.num++
	// 	return nil, nil
	// }

	sentenceNumber, err := strconv.Atoi(line[0])
	if err != nil {
		return nil, rsliberrors.Errorf("failed to parse sentenceNumber. rowNumber: %d, value: %s, err: %w", r.num, line[0], err)
	}

	lang3, err := libdomain.NewLang3(line[1])
	if err != nil {
		return nil, rsliberrors.Errorf("failed to NewLang3. rowNumber: %d, value: %s, err: %w", r.num, line[1], err)
	}

	text := line[2]
	author := line[3]

	time1 := line[4]
	time2 := line[5]

	if len(([]rune(text))) > textLimitLength {
		// skip
		r.logger.DebugContext(ctx, fmt.Sprintf("skip long text. rowNumber: %d, text: %s, len: %d", r.num, text, len(text)))
		r.num++
		return nil, nil
	}

	//\N	2020-02-23 05:07:26
	updatedAt, err := r.getUpdatedTime(time1, time2)
	if err != nil {
		return nil, err
	}

	param, err := service.NewTatoebaSentenceAddParameter(sentenceNumber, lang3, text, author, updatedAt)
	if err != nil {
		return nil, rsliberrors.Errorf("failed to NewTatoebaSentenceAddParameter. rowNumber: %d, values: %v, err: %w", r.num, line, err)
	}

	r.num++
	return param, nil
}

func (r *TatoebaSentenceAddParameterReader) getUpdatedTime(time1, time2 string) (time.Time, error) {
	if r.isValidDatetime(time1) || r.isValidDatetime(time2) {
		var timeS string
		if r.isValidDatetime(time1) {
			timeS = time1
		} else if r.isValidDatetime(time2) {
			timeS = time2
		}

		timeTmp, err := time.Parse("2006-01-02 15:04:05", timeS)
		if err != nil {
			return time.Time{}, rsliberrors.Errorf("failed to Parse. rowNumber: %d, value: %s, err: %w", r.num, timeS, err)
		}
		return timeTmp, nil
	}
	return time.Now(), nil
}

func (r *TatoebaSentenceAddParameterReader) isValidDatetime(value string) bool {
	return value != "\\N" && value != "0000-00-00 00:00:00"
}
