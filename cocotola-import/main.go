package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/firestore/word"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/sqls"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/tatoeba"
	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	"google.golang.org/api/option"
	"gorm.io/gorm"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

func main() {
	ctx := context.Background()
	rsliberrors.UseXerrorsErrorf()
	if false {
		if err := import1(ctx); err != nil {
			libdomain.CheckError(err)
		}
	}
	if true {
		if err := import2(ctx); err != nil {
			libdomain.CheckError(err)
		}
	}
}

func import1(ctx context.Context) error {
	cfg, err := config.LoadConfig("local")
	if err != nil {
		return err
	}

	if err := tatoeba.ImportTatoebaSentences(ctx, "data", cfg.DataSource.TatoebaDataSource.EngSentencesFile); err != nil {
		return err
	}
	if err := tatoeba.ImportTatoebaSentences(ctx, "data", cfg.DataSource.TatoebaDataSource.JpnSentencesFile); err != nil {
		return err
	}
	if err := tatoeba.ImportTatoebaLinks(ctx, "data", cfg.DataSource.TatoebaDataSource.LinksFile); err != nil {
		return err
	}
	return nil
}

type ExportHistorEntity struct {
	ID         int
	WorkbookID int
	Status     string
	ExportedAt time.Time
}

func (u *ExportHistorEntity) TableName() string {
	return "export_history"
}
func import2(ctx context.Context) error {
	fmt.Println("import2")

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

	latestExportedAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	if value != 0 {
		fmt.Println("import4")
		exportHistorEntities := []ExportHistorEntity{}
		if result := db.Where("workbook_id = ?", 1).First(&exportHistorEntities).Order("exported_at desc"); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				latestExportedAt = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return rsliberrors.Errorf("Select: %w", result.Error)
			}
		} else {
			latestExportedAt = exportHistorEntities[0].ExportedAt
		}
	}

	logger := slog.Default()
	logger.InfoContext(ctx, fmt.Sprintf("latestExportedAt: %v", latestExportedAt))

	if err := word.ExportAllWords(ctx, client, latestExportedAt); err != nil {
		return rsliberrors.Errorf("ExportAllWords: %w", err)
	}

	jsonFile, err := os.Open("fileName.json")
	if err != nil {
		return rsliberrors.Errorf("Open: %w", err)
	}
	if err := word.ImportNewWordSentencePairs(ctx, client, db, jsonFile); err != nil {
		return rsliberrors.Errorf("ImportNewWordSentencePairs: %w", err)
	}

	if result := db.Create(&ExportHistorEntity{
		WorkbookID: 1,
		Status:     "success",
		ExportedAt: time.Now(),
	}); result.Error != nil {
		return rsliberrors.Errorf("Create: %w", result.Error)
	}
	return nil
}
