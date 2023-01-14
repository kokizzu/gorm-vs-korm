package gorm_vs_korm

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kokizzu/gotro/S"
	"github.com/sourcegraph/conc/pool"
	"github.com/zeebo/assert"
)

var pgxPostgres *pgxpool.Pool

func BenchmarkInsert_Postgres_Pgx(b *testing.B) {
	if done() {
		b.SkipNow()
		return
	}
	defer timing()()
	b.ReportAllocs()
	b.ResetTimer()
	b.N = total
	ctx := context.Background()
	_, err := pgxPostgres.Exec(ctx, `TRUNCATE TABLE `+pgxTableName)
	assert.Nil(b, err)
	p := pool.New().WithMaxGoroutines(cores)
	for z := uint64(1); z <= total; z++ {
		z := z
		p.Go(func() {
			_, err := pgxPostgres.Exec(ctx, `INSERT INTO `+pgxTableName+` (id, content) VALUES ($1, $2)`, z, S.EncodeCB63(z, 0))
			assert.Nil(b, err)
		})
	}
	p.Wait()
}

func BenchmarkGetAllS_Postgres_Pgx(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			rows, err := pgxPostgres.Query(ctx, `SELECT * FROM `+pgxTableName+` LIMIT $1`, limit)
			assert.Nil(b, err)
			defer rows.Close()
			var row PgxTestTable
			for rows.Next() {
				err = rows.Scan(&row.Id, &row.Content)
				assert.Nil(b, err)
			}
		})
	}
	p.Wait()
}

func BenchmarkGetRowS_Postgres_Pgx(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()
	p := pool.New().WithMaxGoroutines(cores)
	for i := uint64(1); i <= uint64(b.N); i++ {
		p.Go(func() {
			row := pgxPostgres.QueryRow(ctx, `SELECT * FROM `+pgxTableName+` WHERE content = $1 LIMIT 1`, S.EncodeCB63(1+(i%total), 0))
			var row2 PgxTestTable
			err := row.Scan(&row2.Id, &row2.Content)
			assert.Nil(b, err)
		})
	}
	p.Wait()
}
