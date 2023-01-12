package mTest

import (
	"testing"

	"github.com/kamalshkeir/klog"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/sqlitedriver"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	// init gorm and korm
	var err error
	sqlitedriver.Use()
	gormDB, err = gorm.Open(sqlite.Open("benchgorm.sqlite"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if klog.CheckError(err) {
		return
	}
	err = gormDB.AutoMigrate(&TestTableGorm{})
	if klog.CheckError(err) {
		return
	}
	dest := []TestTableGorm{}
	err = gormDB.Find(&dest, &TestTableGorm{}).Error
	if err != nil || len(dest) == 0 {
		err := gormDB.Create(&TestTableGorm{
			Content: "test",
		}).Error
		if klog.CheckError(err) {
			return
		}
	}
	_ = korm.New(korm.SQLITE, "bench")
	// migrate table test_table from struct TestTable
	err = korm.AutoMigrate[TestTable]("test_table")
	if klog.CheckError(err) {
		return
	}
	t, _ := korm.Table("test_table").All()
	if len(t) == 0 {
		_, err := korm.Model[TestTable]().Insert(&TestTable{
			Content: "test",
		})
		klog.CheckError(err)
	}
	korm.DisableCache()
	// ^ the cheat for 2m performance
	// caveat:
	// - database cannot be more than half of RAM size, since the cache has no limit
	// - you cannot update record using another tool (eg. database IDEs) since it didn't evict the cache

	// init tarantool
	taran = &Tt.Adapter{Connection: connectTarantool(), Reconnect: connectTarantool}
	taran.MigrateTables(tables)
	row1 := TestTable2{Id: 1, Content: `test`}
	_, err = taran.Replace(TableTestTable2, row1.ToArray())
	L.PanicIf(err, `failed to insert data`)

	m.Run()
}
