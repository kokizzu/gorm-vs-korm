package gorm_vs_korm

import (
	"testing"

	"github.com/kamalshkeir/korm"
	"github.com/kokizzu/gotro/S"
	"github.com/sourcegraph/conc/pool"
	"github.com/zeebo/assert"
)

func BenchmarkInsertS_Cockroach_Korm(b *testing.B) {
	if done() {
		b.SkipNow()
		return
	}
	defer timing()()
	b.N = total
	err := korm.Exec(kormCockroachDbName, `TRUNCATE TABLE `+kormTableName)
	assert.Nil(b, err)
	p := pool.New().WithMaxGoroutines(cores)
	for z := uint64(1); z <= total; z++ {
		z := z
		p.Go(func() {
			_, err := korm.Model[KormTestTable]().Database(kormCockroachDbName).Insert(&KormTestTable{
				Id:      z,
				Content: S.EncodeCB63(z, 0),
			})
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkUpdate_Cockroach_Korm(b *testing.B) {
	if done() {
		b.SkipNow()
		return
	}
	defer timing()(2)
	b.N = total
	p := pool.New().WithMaxGoroutines(cores)
	for z := uint64(1); z <= total; z++ {
		z := z
		p.Go(func() {
			_, err := korm.Model[KormTestTable]().Database(kormCockroachDbName).
				Where("id = ?", z).
				Set(
					`content = ?`, S.EncodeCB63(total+z, 0),
				)
			assert.Nil(b, err)
			_, err = korm.Model[KormTestTable]().Database(kormCockroachDbName).
				Where("id = ?", z).
				Set(
					`content = ?`, S.EncodeCB63(z, 0),
				)
			assert.Nil(b, err)
		})
	}
	b.N *= 2
	p.Wait()
}

func BenchmarkGetAllS_Cockroach_Korm(b *testing.B) {
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			_, err := korm.Model[KormTestTable]().Database(kormCockroachDbName).Limit(limit).All()
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetAllM_Cockroach_Korm(b *testing.B) {
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			_, err := korm.Table(kormTableName).Database(kormCockroachDbName).Limit(limit).All()
			if err != nil {
				b.Error("error BenchmarkGetAllM:", err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_Cockroach_Korm(b *testing.B) {
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			_, err := korm.Model[KormTestTable]().Where("content = ?",
				S.EncodeCB63(1+uint64(i)%total, 0),
			).Database(kormCockroachDbName).One()
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetRowM_Cockroach_Korm(b *testing.B) {
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			_, err := korm.Table(kormTableName).Database(kormCockroachDbName).Where("content = ?",
				S.EncodeCB63(1+uint64(i)%total, 0),
			).One()
			assert.Nil(b, err)
		})
	}
	p.Wait()
}
