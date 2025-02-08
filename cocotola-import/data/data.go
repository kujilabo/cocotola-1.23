package data

import (
	"embed"
	_ "embed"
)

//go:embed *.tsv
//go:embed *.csv
var Data embed.FS
