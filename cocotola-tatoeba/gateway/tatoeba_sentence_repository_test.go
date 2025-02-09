package gateway_test

import (
	"reflect"
	"testing"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/gateway"
	testlibgateway "github.com/kujilabo/cocotola-1.23/redstart/testlib/gateway"
	"github.com/stretchr/testify/assert"
)

func Test_tatoebaSentenceRepository_FindTatoebaSentences(t *testing.T) {
	// logrus.SetLevel(logrus.DebugLevel)

	for _, db := range testlibgateway.ListDB() {
		// fmt.Println(driverName)
		sqlDB, err := db.DB()
		assert.NoError(t, err)
		defer sqlDB.Close()
	}
}

func TestSplitString(t *testing.T) {
	type args struct {
		str   string
		space rune
		quote rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "split string",
			args: args{
				str:   " arg1 arg2 \"hello world\" ",
				space: ' ',
				quote: '"',
			},
			want: []string{"arg1", "arg2", "hello world"},
		},
		{
			name: "split string",
			args: args{
				str:   "",
				space: ' ',
				quote: '"',
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gateway.SplitString(tt.args.str, tt.args.space, tt.args.quote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
