package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/firestore"
	"github.com/kujilabo/cocotola-1.23/cocotola-import/tatoeba"
)

func main() {
	ctx := context.Background()
	if err := import2(ctx); err != nil {
		log.Fatal(fmt.Sprintf("fireabase error: %+v", err))
		panic(err)
	}
}

func import1() error {
	if err := tatoeba.ImportTatoebaEngSentences(); err != nil {
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
