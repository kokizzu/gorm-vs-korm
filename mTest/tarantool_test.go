package mTest

import (
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/sourcegraph/conc/pool"
)

var taran *Tt.Adapter

func TestGenerateOrm(t *testing.T) {
	Tt.GenerateOrm(tables, false)
}

func init() {
	taran = &Tt.Adapter{Connection: connectTarantool(), Reconnect: connectTarantool}
	taran.MigrateTables(tables)
	row1 := TestTable2{Id: 1, Content: `test`}
	_, err := taran.Replace(TableTestTable2, row1.ToArray())
	L.PanicIf(err, `failed to insert data`)
}

const queryAll = `SELECT * FROM "test_table2"`
const queryOne = `SELECT * FROM "test_table2" WHERE "content" = `

func BenchmarkGetAllS_Taran(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			var res []*TestTable2
			_ = taran.QuerySql(queryAll, func(row []any) {
				obj := &TestTable2{}
				obj.FromArray(row)
				res = append(res, obj)
			})
		})
	}
	p.Wait()
}

func BenchmarkGetAllM_Taran(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			obj := &TestTable2{}
			var res []map[string]any
			_ = taran.QuerySql(queryAll, func(row []any) {
				m := obj.ToMapFromSlice(row)
				res = append(res, m)
			})
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_Taran(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			var res []*TestTable2
			_ = taran.QuerySql(queryOne+S.Z(`test`), func(row []any) {
				obj := &TestTable2{}
				obj.FromArray(row)
				res = append(res, obj)
			})
		})
	}
	p.Wait()
}

func BenchmarkGetRowM_Taran(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(32)
	for i := 0; i < b.N; i++ {
		p.Go(func() {
			obj := &TestTable2{}
			_ = taran.QuerySql(queryOne+S.Z(`test`), func(row []any) {
				_ = obj.ToMapFromSlice(row)
			})
		})
	}
	p.Wait()
}
