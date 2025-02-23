package import_firestore

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"gorm.io/gorm"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/sqls"
)

func ImportLocalFilesToFirestore(ctx context.Context) error {
	logger := slog.Default()
	logger.InfoContext(ctx, "importLocalFilesToFirestore")

	sa := option.WithCredentialsFile("../keys/firestore-importer.serviceAccount.json")

	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "cocotola-1-23-develop-24-11-02"}, sa)
	if err != nil {
		return rsliberrors.Errorf("firebase.NewApp err: %w", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return rsliberrors.Errorf("app.Firestore err: %w", err)
	}
	defer client.Close()

	// init db
	_, db, sqlDB, err := rslibconfig.InitDB(ctx, &rslibconfig.DBConfig{
		DriverName: "sqlite3",
		SQLite3: &rslibgateway.SQLite3Config{
			File: "./data/firestore.db",
		},
		Migration: true,
	}, sqls.SQL)
	libdomain.CheckError(err)
	defer sqlDB.Close()

	var value int64
	if result := db.Table("export_history").Select("*").Count(&value); result.Error != nil {
		return rsliberrors.Errorf("Select: %w", result.Error)
	}

	return importLocalFilesToFirestore(ctx, client, db)
}

func importLocalFilesToFirestore(ctx context.Context, client *firestore.Client, db *gorm.DB) error {
	logger := slog.Default()

	exportHistoryRepository := gateway.NewExportHistoryRepository(db)
	latestExportedAt, err := exportHistoryRepository.GetLatestExportedAt()
	if err != nil {
		return err
	}

	logger.InfoContext(ctx, fmt.Sprintf("latestExportedAt: %v", latestExportedAt))

	if err := ExportAllWords(ctx, client, latestExportedAt); err != nil {
		return rsliberrors.Errorf("ExportAllWords: %w", err)
	}

	jsonFile, err := os.Open("fileName.json")
	if err != nil {
		return rsliberrors.Errorf("Open: %w", err)
	}
	if err := ImportNewWordSentencePairs(ctx, client, db, jsonFile); err != nil {
		return rsliberrors.Errorf("ImportNewWordSentencePairs: %w", err)
	}
	if err := exportHistoryRepository.Add(1, "success", time.Now()); err != nil {
		return rsliberrors.Errorf("Create: %w", err)
	}
	return nil
}
