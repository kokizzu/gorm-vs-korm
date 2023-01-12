package mTest

import (
	"testing"

	"github.com/kamalshkeir/korm"
	"github.com/sourcegraph/conc/pool"
)

type TestTable struct {
	Id      uint   `korm:"pk"`
	Content string `korm:"size:50"`
}

func BenchmarkGetAllS_Sqlite_Korm(b *testing.B) {
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

func BenchmarkGetAllM_Sqlite_Korm(b *testing.B) {
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

func BenchmarkGetRowS_Sqlite_Korm(b *testing.B) {
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

func BenchmarkGetRowM_Sqlite_Korm(b *testing.B) {
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
