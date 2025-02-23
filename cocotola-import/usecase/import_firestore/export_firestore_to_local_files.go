package import_firestore

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"cloud.google.com/go/firestore"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

func ExportAllWords(ctx context.Context, client *firestore.Client, latestExportedAt time.Time) error {
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
