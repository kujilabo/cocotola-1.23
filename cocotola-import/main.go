package main

import (
	"context"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/tatoeba"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/usecase/import_firestore"
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
		if err := import_firestore.ImportLocalFilesToFirestore(ctx); err != nil {
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
