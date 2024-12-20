package gateway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	testlibgateway "github.com/kujilabo/cocotola-1.23/redstart/testlib/gateway"
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
