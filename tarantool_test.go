package gorm_vs_korm

import (
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/sourcegraph/conc/pool"
	"github.com/tarantool/go-tarantool"
	"github.com/zeebo/assert"

	"korm1/mTest"
	"korm1/mTest/rqTest"
	"korm1/mTest/wcTest"
)

func connectTarantool() *tarantool.Connection {
	taran, err := tarantool.Connect(`localhost:3301`, tarantool.Opts{})
	L.PanicIf(err, `failed to connect to tarantool`)
	return taran
}

var taran *Tt.Adapter

// only need to do once before compile
func TestGenerateOrm(t *testing.T) {
	Tt.GenerateOrm(mTest.Tables, false)
	t.SkipNow()
}

const queryAll = `SELECT * FROM "test_table2"`
const queryOne = `SELECT * FROM "test_table2" WHERE "content" = `

func BenchmarkInsertS_Taran_ORM(b *testing.B) {
	if done() {
		b.SkipNow()
		return
	}
	defer timing()()
	b.ReportAllocs()
	b.ResetTimer()
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
			var res []*mTest.TestTable2
			_ = taran.QuerySql(queryAll, func(row []any) {
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
			var res []map[string]any
			_ = taran.QuerySql(queryAll, func(row []any) {
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
			var res []*mTest.TestTable2
			_ = taran.QuerySql(queryOne+S.Z(
				S.EncodeCB63(i, 0)), func(row []any) {
				obj := &mTest.TestTable2{}
				obj.FromArray(row)
				res = append(res, obj)
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