package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/data"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/firestore"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/tatoeba"
)

func main() {
	ctx := context.Background()
	if err := import1(ctx); err != nil {
		log.Fatalf("fireabase error: %+v", err)
		panic(err)
	}
	// if err := import2(ctx); err != nil {
	// 	log.Fatalf("fireabase error: %+v", err)
	// 	panic(err)
	// }
}

func import1(ctx context.Context) error {
	x := len("もう彼女には言えないよ。そんなに単純なことではなくなってきたからね。")
	logger := slog.Default()
	logger.InfoContext(ctx, fmt.Sprintf("len: %d", x))
	cfg, err := config.LoadConfig("local")
	if err != nil {
		return err
	}

	if err := tatoeba.ImportTatoebaSentences(ctx, data.Data, cfg.DataSource.TatoebaDataSource.EngSentencesFile); err != nil {
		return err
	}
	if err := tatoeba.ImportTatoebaSentences(ctx, data.Data, cfg.DataSource.TatoebaDataSource.JpnSentencesFile); err != nil {
		return err
	}
	if err := tatoeba.ImportTatoebaLinks(ctx, data.Data, cfg.DataSource.TatoebaDataSource.LinksFile); err != nil {
		return err
	}
	return nil
}

func import2(ctx context.Context) error {
	if err := firestore.Firebase(ctx); err != nil {
		return err
	}
	return nil
}
