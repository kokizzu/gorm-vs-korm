
# Gorm+Korm (Sqlite, PostgreSQL, ) vs Tarantool

```shell
./clean-start.sh
go test -bench=Korm -benchmem .
go test -bench=Gorm -benchmem .
go test -bench=Taran -benchmem .

goos: linux
goarch: amd64

S = struct, M = map, A = array

10K rows, GetAll select all rows, concurrency: 32

InsertS_Cockroach_Gorm-32   10000    163963 ns/op       1.64 s         
InsertS_Cockroach_Korm-32   10000    434172 ns/op       4.34 s
InsertS_Postgres_Gorm-32    10000    154924 ns/op       1.55 s
InsertS_Postgres_Korm-32    10000    163079 ns/op       1.63 s
InsertS_Sqlite_Gorm-32      10000   1234051 ns/op      12.34 s
InsertS_Sqlite_Korm-32      10000   2050994 ns/op      20.51 s
InsertS_Taran_ORM-32        10000     29680 ns/op       0.30 s

GetAllM_Cockroach_Gorm-32    1879    634957 ns/op   294824 B/op       6759 allocs/op
GetAllM_Cockroach_Korm-32     722   1478724 ns/op  4278933 B/op      59979 allocs/op
GetAllM_Postgres_Gorm-32     3664    316556 ns/op   570431 B/op      12983 allocs/op
GetAllM_Postgres_Korm-32     1327    885765 ns/op  4036508 B/op      59714 allocs/op
GetAllM_Sqlite_Gorm-32      10000    431210 ns/op  1578220 B/op      35069 allocs/op
GetAllM_Sqlite_Korm-32        978   1134521 ns/op  4439227 B/op      89977 allocs/op
GetAllM_Taran_Raw-32          156   7777074 ns/op  4566475 B/op      69736 allocs/op

GetAllS_Cockroach_Gorm-32     999   1071982 ns/op  2352898 B/op      79689 allocs/op
GetAllS_Cockroach_Korm-32     715   1486551 ns/op  1992995 B/op      80002 allocs/op
GetAllS_Postgres_Gorm-32     1394    730572 ns/op  2351424 B/op      79682 allocs/op
GetAllS_Postgres_Korm-32     1159   1041059 ns/op  1990956 B/op      79748 allocs/op
GetAllS_Sqlite_Gorm-32       1047   1063228 ns/op  2828802 B/op     139728 allocs/op
GetAllS_Sqlite_Korm-32        736   1477125 ns/op  2393110 B/op     119993 allocs/op
GetAllS_Taran_Raw-32          156   7687067 ns/op  1446471 B/op      59736 allocs/op

GetRowM_Cockroach_Gorm-32   15852     74238 ns/op    64219 B/op        467 allocs/op
GetRowM_Cockroach_Korm-32   32973     35975 ns/op     1695 B/op         43 allocs/op
GetRowM_Postgres_Gorm-32    25344     45178 ns/op    23940 B/op        208 allocs/op
GetRowM_Postgres_Korm-32    70509     15877 ns/op     1695 B/op         42 allocs/op
GetRowM_Sqlite_Gorm-32      72108     16634 ns/op     4817 B/op        110 allocs/op
GetRowM_Sqlite_Korm-32     323451      3202 ns/op     1432 B/op         45 allocs/op
GetRowM_Taran_Raw-32       160676      7275 ns/op     2425 B/op         51 allocs/op

GetRowS_Cockroach_Gorm-32   20312     61121 ns/op    71997 B/op        512 allocs/op
GetRowS_Cockroach_Korm-32    9338    170849 ns/op     2687 B/op         71 allocs/op
GetRowS_Postgres_Gorm-32    26478     46250 ns/op    23897 B/op        202 allocs/op
GetRowS_Postgres_Korm-32    10000    137373 ns/op     2682 B/op         71 allocs/op
GetRowS_Sqlite_Gorm-32      70897     16205 ns/op     4183 B/op         92 allocs/op
GetRowS_Sqlite_Korm-32     117987     12081 ns/op     2148 B/op         64 allocs/op
GetRowS_Taran_ORM-32       298116      3726 ns/op     1057 B/op         24 allocs/op
GetRowS_Taran_Raw-32       161505      7447 ns/op     2425 B/op         51 allocs/op

100K rows, GetAll select 1000 rows unordered, concurrency: 32

SQLite = too slow
Gorm = too many errors, connection reset by peer

InsertS_Cockroach_Korm-32   100000   451436 ns/op      45.14 s
InsertS_Postgres_Korm-32    100000   172047 ns/op      17.20 s
InsertS_Taran_ORM-32        100000    36685 ns/op       3.67 s
Insert_Cockroach_Pgx-32     100000    99197 ns/op       9.92 s
Insert_Postgres_Pgx-32      100000    56311 ns/op       5.63 s

GetAllM_Cockroach_Korm-32     5757   191306 ns/op   418047 B/op  5973 allocs/op
GetAllM_Postgres_Korm-32     10000   117168 ns/op   391829 B/op  5729 allocs/op
GetAllM_Taran_Raw-32          1640   742542 ns/op  1248536 B/op  6731 allocs/op

GetAllS_Cockroach_Korm-32     4272   244311 ns/op   167806 B/op  7998 allocs/op
GetAllS_Cockroach_Pgx-32     17805    67558 ns/op    34493 B/op  1951 allocs/op
GetAllS_Postgres_Korm-32      6567   185029 ns/op   165941 B/op  7764 allocs/op
GetAllS_Postgres_Pgx-32      45964    26192 ns/op    34505 B/op  1965 allocs/op
GetAllS_Taran_ORM-32          4180   291447 ns/op   233928 B/op  4714 allocs/op
GetAllS_Taran_Raw-32          1689   734751 ns/op   936548 B/op  5731 allocs/op

GetAllA_Taran_ORM-32          4146   286855 ns/op   157546 B/op  4703 allocs/op

GetRowM_Cockroach_Korm-32    32676    36936 ns/op     1697 B/op    43 allocs/op
GetRowM_Postgres_Korm-32     68497    16393 ns/op     1696 B/op    43 allocs/op
GetRowM_Taran_Raw-32        130099     8951 ns/op     2498 B/op    55 allocs/op

GetRowS_Cockroach_Korm-32     9468   153688 ns/op     2684 B/op    71 allocs/op
GetRowS_Cockroach_Pgx-32     78200    14916 ns/op      621 B/op    15 allocs/op
GetRowS_Postgres_Korm-32     10000   128368 ns/op     2682 B/op    71 allocs/op
GetRowS_Postgres_Pgx-32     226089     5308 ns/op      619 B/op    15 allocs/op
GetRowS_Taran_ORM-32        297463     3724 ns/op     1058 B/op    24 allocs/op
GetRowS_Taran_Raw-32        113793    10017 ns/op     2509 B/op    56 allocs/op
```

Note:
- disabled cache for KORM
- tarantool not using ORM nor faster API, just raw sql query
- tarantool 10x less rps when only 1 core utilized (without conc)
