package main

import "github.com/kujilabo/cocotola-1.23/cocotola-import/tatoeba"

func main() {
	if err := tatoeba.ImportTatoebaEngSentences(); err != nil {
		panic(err)
	}
}
