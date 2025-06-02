package benchmark

import (
	"github.com/SevgiF/location-api/internal/adapters/test"
	"github.com/SevgiF/location-api/pkg/database/mysql"
)

func Clean() {
	var db = mysql.NewGormDBManager().DB
	test.TeardownTest(db)
}
