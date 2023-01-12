package mTest

import (
	"testing"

	"github.com/sourcegraph/conc/pool"
	"gorm.io/gorm"
)

type TestTableGorm struct {
	ID      uint `gorm:"primarykey"`
	Content string
}

var gormDB *gorm.DB

func BenchmarkGetAllS_Sqlite_GORM(b *testing.B) {
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

func BenchmarkGetAllM_Sqlite_GORM(b *testing.B) {
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

func BenchmarkGetRowS_Sqlite_GORM(b *testing.B) {
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

func BenchmarkGetRowM_Sqlite_GORM(b *testing.B) {
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
