
# Gorm+Korm (Sqlite, PostgreSQL, ) vs Tarantool

```shell
./clean-start.sh
go test -bench=Korm -benchmem .
go test -bench=Gorm -benchmem .
go test -bench=Taran -benchmem .

goos: linux
goarch: amd64

BenchmarkInsertS_Cockroach_Korm-32         10000            434172 ns/op              4.34 s
BenchmarkInsertS_Cockroach_Korm-32         10000            434175 ns/op            2764 B/op         55 allocs/op
BenchmarkGetAllS_Cockroach_Korm-32           715           1486551 ns/op         1992995 B/op      80002 allocs/op
BenchmarkGetAllM_Cockroach_Korm-32           722           1478724 ns/op         4278933 B/op      59979 allocs/op
BenchmarkGetRowS_Cockroach_Korm-32          9338            170849 ns/op            2687 B/op         71 allocs/op
BenchmarkGetRowM_Cockroach_Korm-32         32973             35975 ns/op            1695 B/op         43 allocs/op

BenchmarkInsertS_Postgres_Korm-32          10000            163079 ns/op              1.63 s
BenchmarkInsertS_Postgres_Korm-32          10000            163082 ns/op            2739 B/op         55 allocs/op
BenchmarkGetAllS_Postgres_Korm-32           1159           1041059 ns/op         1990956 B/op      79748 allocs/op
BenchmarkGetAllM_Postgres_Korm-32           1327            885765 ns/op         4036508 B/op      59714 allocs/op
BenchmarkGetRowS_Postgres_Korm-32          10000            137373 ns/op            2682 B/op         71 allocs/op
BenchmarkGetRowM_Postgres_Korm-32          70509             15877 ns/op            1695 B/op         42 allocs/op

BenchmarkInsertS_Sqlite_Korm-32            10000           2050994 ns/op             20.51 s
BenchmarkInsertS_Sqlite_Korm-32            10000           2050996 ns/op             20.51 s
BenchmarkInsertS_Sqlite_Korm-32            10000           2050997 ns/op            2075 B/op         52 allocs/op
BenchmarkGetAllS_Sqlite_Korm-32              736           1477125 ns/op         2393110 B/op     119993 allocs/op
BenchmarkGetAllM_Sqlite_Korm-32              978           1134521 ns/op         4439227 B/op      89977 allocs/op
BenchmarkGetRowS_Sqlite_Korm-32           117987             12081 ns/op            2148 B/op         64 allocs/op
BenchmarkGetRowM_Sqlite_Korm-32           323451              3202 ns/op            1432 B/op         45 allocs/op

BenchmarkInsertS_Cockroach_Gorm-32         10000            163963 ns/op              1.64 s         
BenchmarkInsertS_Cockroach_Gorm-32         10000            159524 ns/op           31334 B/op        260 allocs/op
BenchmarkGetAllS_Cockroach_Gorm-32           999           1071982 ns/op         2352898 B/op      79689 allocs/op
BenchmarkGetAllM_Cockroach_Gorm-32          1879            634957 ns/op          294824 B/op       6759 allocs/op
BenchmarkGetRowS_Cockroach_Gorm-32         20312             61121 ns/op           71997 B/op        512 allocs/op
BenchmarkGetRowM_Cockroach_Gorm-32         15852             74238 ns/op           64219 B/op        467 allocs/op

BenchmarkInsertS_Postgres_Gorm-32          10000            154924 ns/op              1.55 s
BenchmarkInsertS_Postgres_Gorm-32          10000            154424 ns/op           64887 B/op        496 allocs/op
BenchmarkGetAllS_Postgres_Gorm-32           1394            730572 ns/op         2351424 B/op      79682 allocs/op
BenchmarkGetAllM_Postgres_Gorm-32           3664            316556 ns/op          570431 B/op      12983 allocs/op
BenchmarkGetRowS_Postgres_Gorm-32          26478             46250 ns/op           23897 B/op        202 allocs/op
BenchmarkGetRowM_Postgres_Gorm-32          25344             45178 ns/op           23940 B/op        208 allocs/op

BenchmarkInsertS_Sqlite_Gorm-32            10000           1234051 ns/op             12.34 s
BenchmarkInsertS_Sqlite_Gorm-32            10000           1233827 ns/op            3967 B/op         72 allocs/op
BenchmarkGetAllS_Sqlite_Gorm-32             1047           1063228 ns/op         2828802 B/op     139728 allocs/op
BenchmarkGetAllM_Sqlite_Gorm-32            10000            431210 ns/op         1578220 B/op      35069 allocs/op
BenchmarkGetRowS_Sqlite_Gorm-32            70897             16205 ns/op            4183 B/op         92 allocs/op
BenchmarkGetRowM_Sqlite_Gorm-32            72108             16634 ns/op            4817 B/op        110 allocs/op

BenchmarkInsertS_Taran_ORM-32              10000             29680 ns/op              0.30 s
BenchmarkGetAllS_Taran_Raw-32                156           7687067 ns/op         1446471 B/op      59736 allocs/op
BenchmarkGetAllM_Taran_Raw-32                156           7777074 ns/op         4566475 B/op      69736 allocs/op
BenchmarkGetRowS_Taran_Raw-32             161505              7447 ns/op            2425 B/op         51 allocs/op
BenchmarkGetRowM_Taran_Raw-32             160676              7275 ns/op            2425 B/op         51 allocs/op
BenchmarkGetRowS_Taran_ORM-32             298116              3726 ns/op            1057 B/op         24 allocs/op
```

Note:
- disabled cache for KORM
- tarantool not using ORM nor faster API, just raw sql query
- tarantool 10x less rps when only 1 core utilized (without conc)
