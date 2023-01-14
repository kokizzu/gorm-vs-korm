package gorm_vs_korm

import (
	"testing"

	"github.com/kokizzu/gotro/S"
	"github.com/sourcegraph/conc/pool"
	"github.com/zeebo/assert"
	"gorm.io/gorm"
)

var gormPostgres *gorm.DB

func BenchmarkInsertS_Postgres_Gorm(b *testing.B) {
	if done() {
		b.SkipNow()
		return
	}
	defer timing()()
	err := gormPostgres.Exec(`TRUNCATE TABLE ` + gormTestTableName).Error
	assert.Nil(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	b.N = total
	p := pool.New().WithMaxGoroutines(cores)
	for z := uint64(1); z <= total; z++ {
		z := z
		p.Go(func() {
			err := gormPostgres.Create(&GormTestTable{
				ID:      z,
				Content: S.EncodeCB63(z, 0),
			}).Error
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetAllS_Postgres_Gorm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			a := []GormTestTable{}
			err := gormPostgres.Find(&a).Error
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetAllM_Postgres_Gorm(b *testing.B) {
	a := []map[string]any{}
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			err := gormPostgres.Find(&GormTestTable{}).Scan(&a).Error
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_Postgres_Gorm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			u := GormTestTable{}
			err := gormPostgres.Where(&GormTestTable{
				Content: S.EncodeCB63(1+uint64(i)%total, 0),
			}).First(&u).Error
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetRowM_Postgres_Gorm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		i := i
		p.Go(func() {
			u := map[string]any{}
			err := gormPostgres.Model(&GormTestTable{}).Where(&GormTestTable{
				Content: S.EncodeCB63(1+uint64(i)%total, 0),
			}).First(&u).Error
			assert.Nil(b, err)
		})
	}
	p.Wait()
}
