package mTest

import (
	"testing"

	"github.com/kamalshkeir/klog"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/sqlitedriver"
	"github.com/sourcegraph/conc/pool"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestTable struct {
	Id      uint   `korm:"pk"`
	Content string `korm:"size:50"`
}

type TestTableGorm struct {
	ID      uint `gorm:"primarykey"`
	Content string
}

var gormDB *gorm.DB

func init() {
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
}

func BenchmarkGetAllS_GORM(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			a := []TestTableGorm{}
			err := gormDB.Find(&a).Error
			if err != nil {
				b.Error("error BenchmarkGetAllS_GORM:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetAllM_GORM(b *testing.B) {
	a := []map[string]any{}
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			err := gormDB.Find(&TestTableGorm{}).Scan(&a).Error
			if err != nil {
				b.Error("error BenchmarkGetAllM_GORM:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_GORM(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			u := TestTableGorm{}
			err := gormDB.Where(&TestTableGorm{
				Content: "test",
			}).First(&u).Error
			if err != nil {
				b.Error("error BenchmarkGetRowS_GORM:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetRowM_GORM(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			u := map[string]any{}
			err := gormDB.Model(&TestTableGorm{}).Where(&TestTableGorm{
				Content: "test",
			}).First(&u).Error
			if err != nil {
				b.Error("error BenchmarkGetRowS_GORM:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetAllS_Korm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			_, err := korm.Model[TestTable]().All()
			if err != nil {
				b.Error("error BenchmarkGetAllS:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetAllM_Korm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			_, err := korm.Table("test_table").All()
			if err != nil {
				b.Error("error BenchmarkGetAllM:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_Korm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			_, err := korm.Model[TestTable]().Where("content = ?", "test").One()
			if err != nil {
				b.Error("error BenchmarkGetRowS:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetRowM_Korm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			_, err := korm.Table("test_table").Where("content = ?", "test").One()
			if err != nil {
				b.Error("error BenchmarkGetRowM:", err)
			}
		})
	}
	p.Wait()
}
