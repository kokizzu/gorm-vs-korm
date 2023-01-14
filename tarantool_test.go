package gorm_vs_korm

import (
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/S"
	"github.com/sourcegraph/conc/pool"
	"github.com/zeebo/assert"

	"korm1/mTest"
	"korm1/mTest/rqTest"
	"korm1/mTest/wcTest"
)

var taran *Tt.Adapter

const queryAll = `SELECT * FROM "test_table2"`
const queryOne = `SELECT * FROM "test_table2" WHERE "content" = `
const limit1k = ` LIMIT 1000`

func BenchmarkInsertS_Taran_ORM(b *testing.B) {
	if done() {
		b.SkipNow()
		return
	}
	defer timing()()
	b.ReportAllocs()
	b.ResetTimer()
	b.N = total
	r := taran.ExecSql(`DELETE FROM ` + S.ZZ(mTest.TableTestTable2))
	assert.Equal(b, len(r), 1)

	p := pool.New().WithMaxGoroutines(2)
	for z := uint64(1); z <= total; z++ {
		z := z
		p.Go(func() {
			row := wcTest.NewTestTable2Mutator(taran)
			row.Id = z
			row.Content = S.EncodeCB63(z, 0)
			ok := row.DoReplace()
			assert.True(b, ok)
		})
	}
	p.Wait()
}

func BenchmarkGetAllS_Taran_Raw(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			res := make([]*mTest.TestTable2, 0, total)
			_ = taran.QuerySql(queryAll+limit1k, func(row []any) {
				obj := &mTest.TestTable2{}
				obj.FromArray(row)
				res = append(res, obj)
			})
		})
	}
	p.Wait()
}

func BenchmarkGetAllM_Taran_Raw(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			obj := &mTest.TestTable2{}
			res := make([]map[string]any, 0, total)
			_ = taran.QuerySql(queryAll+limit1k, func(row []any) {
				m := obj.ToMapFromSlice(row)
				res = append(res, m)
			})
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_Taran_Raw(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			obj := mTest.TestTable2{}
			_ = taran.QuerySql(queryOne+S.Z(
				S.EncodeCB63(i, 0)), func(row []any) {
				obj.FromArray(row)
			})
		})
	}
	p.Wait()
}

func BenchmarkGetRowM_Taran_Raw(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			obj := &mTest.TestTable2{}
			_ = taran.QuerySql(queryOne+S.Z(
				S.EncodeCB63(i, 0)), func(row []any) {
				_ = obj.ToMapFromSlice(row)
			})
		})
	}
	p.Wait()
}
func BenchmarkGetRowS_Taran_ORM(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			obj := rqTest.NewTestTable2(taran)
			obj.Content = S.EncodeCB63(1+i%total, 0)
			ok := obj.FindByContent()
			assert.True(b, ok)
		})
	}
	p.Wait()
}
