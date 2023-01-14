package gorm_vs_korm

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/kamalshkeir/korm"
	_ "github.com/kamalshkeir/pgdriver"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"korm1/mTest"
)

const total = 100000
const limit = 1000
const cores = 32

// connection strings

const gormPostgresConnStr = "host=localhost user=pgroot password=password dbname=postgres port=5432 sslmode=disable"
const gormCockroachConnStr = "host=localhost user=root dbname=defaultdb port=26257 sslmode=disable"
const kormPostgresConnStr = `pgroot:password@localhost:5432`
const kormCockroachConnStr = `root@localhost:26257`

// gorm
type GormTestTable struct {
	ID      uint64 `gorm:"primarykey"`
	Content string `gorm:"index:idx_content,unique"`
}

func (g *GormTestTable) TableName() string {
	return gormTestTableName
}

const gormTestTableName = `test_table`

// korm
const kormSqliteDbName = `benchkorm`
const kormTableName = `test_table`
const kormCockroachDbName = `defaultdb`
const kormPostgresDbName = `postgres`

type KormTestTable struct {
	Id      uint64 `korm:"pk"`
	Content string `korm:"unique"`
}

func TestMain(m *testing.M) {
	var err error

	//log.Println(`sqlite`)
	//{
	//	log.Println(`gorm`)
	//	{
	//		sqlitedriver.Use()
	//		gormSqlite, err = gorm.Open(sqlite.Open("benchgorm.sqlite"), &gorm.Config{
	//			SkipDefaultTransaction: true,
	//		})
	//		L.PanicIf(err, `gorm.Open`)
	//		err = gormSqlite.AutoMigrate(&GormTestTable{})
	//		L.IsError(err, `gormSqlite.AutoMigrate`)
	//	}
	//
	//	log.Println(`korm`)
	//	{
	//		err = korm.New(korm.SQLITE, kormSqliteDbName)
	//		L.PanicIf(err, `korm.New`)
	//		err = korm.AutoMigrate[KormTestTable](kormTableName, kormSqliteDbName)
	//		L.PanicIf(err, `korm.AutoMigrate`)
	//	}
	//}

	log.Println(`postgres`)
	{
		log.Println(`gorm`)
		{
			gormPostgres, err = gorm.Open(postgres.Open(gormPostgresConnStr))
			L.PanicIf(err, `gorm.Open`)
			err = gormPostgres.AutoMigrate(&GormTestTable{})
			L.PanicIf(err, `gormPostgres.AutoMigrate`)
		}

		log.Println(`korm`)
		{
			err = korm.New(korm.POSTGRES, kormPostgresDbName, kormPostgresConnStr)
			L.PanicIf(err, `korm.New`)
			err = korm.AutoMigrate[KormTestTable](kormTableName, kormPostgresDbName)
			L.PanicIf(err, `korm.AutoMigrate`)
		}
	}

	log.Println(`cockroach`)
	{
		//log.Println(`gorm`)
		//{
		//	gormCockroach, err = gorm.Open(postgres.Open(gormCockroachConnStr))
		//	L.PanicIf(err, `gorm.Open`)
		//	err = gormCockroach.AutoMigrate(&GormTestTable{})
		//	L.PanicIf(err, `gormCockroach.AutoMigrate`)
		//}

		log.Println(`korm`)
		{
			err = korm.New(korm.COCKROACH, kormCockroachDbName, kormCockroachConnStr)
			L.PanicIf(err, `korm.New`)
			err = korm.AutoMigrate[KormTestTable](kormTableName, kormCockroachDbName)
			L.PanicIf(err, `korm.AutoMigrate`)
		}
	}

	korm.DisableCache()
	// ^ the cheat for 2m performance
	// caveat:
	// - database cannot be more than half of RAM size, since the cache has no limit
	// - you cannot update record using another tool (eg. database IDEs) since it didn't evict the cache

	log.Println(`tarantool`)
	{
		taran = &Tt.Adapter{Connection: mTest.ConnectTarantool(), Reconnect: mTest.ConnectTarantool}
		_, err = taran.Ping()
		L.PanicIf(err, `taran.Ping`)
		mTest.Migrate(taran)
	}

	log.Println(`start test`)
	m.Run()
}

var runOnce = map[string]bool{}

func done() bool {
	caller := L.CallerInfo(2).FuncName
	log.Println(caller)
	if runOnce[caller] {
		return true
	}
	runOnce[caller] = true
	return false
}

var timer = map[string]time.Time{}

func timing() func() {
	start := time.Now()
	return func() {
		dur := time.Since(start)
		// BenchmarkInsertS_Taran_ORM-32              10000             48616 ns/op            0.49 s
		fmt.Printf(`%-36s %11d %17d ns/op   %15.2f s`+"\n",
			fmt.Sprintf("%s-%d", L.CallerInfo(2).FuncName, cores),
			total,
			dur.Nanoseconds()/int64(total),
			dur.Seconds(),
		)
	}
}
