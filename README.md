
# Pgx, Gorm, Korm (Sqlite, PostgreSQL, CockroachDB) vs Tarantool

## Test Env

- cockroachdb 21.1.11
- postgresql 14.0-1
- tarantool 2.8.2 (because latest tagged on this version, not 2.10.4)

```
goos: linux
goarch: amd64
```

## Usage

```shell
./clean-start.sh
go test -bench=Korm -benchmem .
go test -bench=Gorm -benchmem .
go test -bench=Taran -benchmem .
go test -bench=Pgx -benchmem .
```

## 2023-01-14 Result 10K rows, GetAll select all rows, concurrency: 32

- goos: linux, goarch: amd64
- S = struct, M = map, A = array
- disabled cache for KORM
- tarantool 10x less rps when only 1 core utilized (without `conc`)

```
## korm 1.3.8
## gorm 1.24.3
## go-tarantool 1.10.0

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
```

## 2023-01-15  Result 100K rows, GetAll select 1000 rows unordered, concurrency: 32

- SQLite = too slow
- Gorm = too many errors, connection reset by peer
- Korm = failed update postgres tests

```
## korm 1.4.1
## pgx 5.2.0
## go-tarantool 1.10.0

InsertS_Cockroach_Korm-32   100000   451436 ns/op  45.14 s
Insert_Cockroach_Pgx-32     100000    99197 ns/op   9.92 s
InsertS_Postgres_Korm-32    100000   172047 ns/op  17.20 s
Insert_Postgres_Pgx-32      100000    56311 ns/op   5.63 s
InsertS_Taran_ORM-32        100000    36685 ns/op   3.67 s -- fastest

Update_Cockroach_Korm-32    200000    48294 ns/op   9.66 s
Update_Cockroach_Pgx-32     200000   248171 ns/op  49.63 s
Update_Postgres_Pgx-32      200000    50967 ns/op  10.19 s
Update_Taran_ORM-32         200000      221 ns/op   0.04 s -- fastest

GetAllM_Cockroach_Korm-32     8346   132151 ns/op   417864 B/op  5972 allocs/op
GetAllM_Postgres_Korm-32     12662    93979 ns/op   391546 B/op  5716 allocs/op -- fastest
GetAllM_Taran_Raw-32          1640   742542 ns/op  1248536 B/op  6731 allocs/op

GetAllS_Cockroach_Korm-32     5997   200701 ns/op   167815 B/op  7999 allocs/op
GetAllS_Cockroach_Pgx-32     16476    73736 ns/op    58497 B/op  2951 allocs/op
GetAllS_Postgres_Korm-32      7095   166231 ns/op   165827 B/op  7752 allocs/op
GetAllS_Postgres_Pgx-32      40404    30255 ns/op    58503 B/op  2967 allocs/op -- fastest
GetAllS_Taran_ORM-32          4180   291447 ns/op   233928 B/op  4714 allocs/op
GetAllS_Taran_Raw-32          1689   734751 ns/op   936548 B/op  5731 allocs/op

GetAllA_Taran_ORM-32          4146   286855 ns/op   157546 B/op  4703 allocs/op

GetRowM_Cockroach_Korm-32    51686    23119 ns/op     1759 B/op    44 allocs/op
GetRowM_Postgres_Korm-32     93655    12851 ns/op     1759 B/op    44 allocs/op
GetRowM_Taran_Raw-32        130099     8951 ns/op     2498 B/op    55 allocs/op -- fastest

GetRowS_Cockroach_Korm-32     9432   157098 ns/op     2729 B/op    71 allocs/op
GetRowS_Cockroach_Pgx-32     78200    14916 ns/op      621 B/op    15 allocs/op
GetRowS_Postgres_Korm-32     10000   137250 ns/op     2724 B/op    71 allocs/op
GetRowS_Postgres_Pgx-32     226089     5308 ns/op      619 B/op    15 allocs/op
GetRowS_Taran_ORM-32        297463     3724 ns/op     1058 B/op    24 allocs/op -- fastest
GetRowS_Taran_Raw-32        113793    10017 ns/op     2509 B/op    56 allocs/op 
```
## 2023-01-18 100K Rows, GetAll Select 1000 Rows ordered, concurrency: 32

- korm still failed for postgres-update benchmark BenchmarkUpdate_Postgres_Korm
- korm return error when failed to set cache, so have to check for map.ErrLargeData
- enable korm cache but limit to 1MB, since it would make realistic benchmark for cases when database multitude times larger than RAM size

```
## korm 1.4.3
## pgx 5.2.0
## go-tarantool 1.10.0

GetAllA_Taran_ORM-32          3799   294895 ns/op    157528 B/op   4702 allocs/op
                                                                                
GetAllM_Cockroach_Korm-32     7966   136781 ns/op    417854 B/op   5972 allocs/op
GetAllM_Postgres_Korm-32     12549    96720 ns/op    391705 B/op   5734 allocs/op -- fastest
GetAllM_Taran_Raw-32          1560   778315 ns/op   1248589 B/op   6733 allocs/op
                                                                                
GetAllS_Cockroach_Korm-32     5606   209751 ns/op    167810 B/op   8000 allocs/op
GetAllS_Cockroach_Pgx-32     14605    81561 ns/op     58492 B/op   2951 allocs/op
GetAllS_Postgres_Korm-32      6732   167764 ns/op    165970 B/op   7770 allocs/op
GetAllS_Postgres_Pgx-32      37880    32951 ns/op     59516 B/op   2996 allocs/op -- fastest
GetAllS_Taran_ORM-32          3889   298250 ns/op    233923 B/op   4714 allocs/op
GetAllS_Taran_Raw-32          1372   779119 ns/op    936611 B/op   5735 allocs/op
                                                                                
GetRowM_Cockroach_Korm-32    51204    23929 ns/op      1759 B/op     44 allocs/op
GetRowM_Postgres_Korm-32     93279    13097 ns/op      1760 B/op     44 allocs/op
GetRowM_Taran_Raw-32        132504     8602 ns/op      2556 B/op     57 allocs/op -- fastest
                                                                                
GetRowS_Cockroach_Korm-32     9536   141470 ns/op      2756 B/op     72 allocs/op
GetRowS_Cockroach_Pgx-32     81296    14748 ns/op       619 B/op     15 allocs/op
GetRowS_Postgres_Korm-32     10000   117153 ns/op      2755 B/op     72 allocs/op
GetRowS_Postgres_Pgx-32     199897     6032 ns/op       619 B/op     15 allocs/op
GetRowS_Taran_ORM-32        304686     3676 ns/op      1114 B/op     26 allocs/op -- fastest
GetRowS_Taran_Raw-32        112923     9554 ns/op      2570 B/op     58 allocs/op
                                                                                    
Insert_Cockroach_Pgx-32     100000    96523 ns/op        9.65 s
Insert_Postgres_Pgx-32      100000    53829 ns/op        5.38 s
InsertS_Cockroach_Korm-32   100000   121974 ns/op       12.19 s
InsertS_Postgres_Korm-32    100000    85382 ns/op        8.54 s
InsertS_Taran_ORM-32        100000    31075 ns/op        3.11 s -- fastest
                                                                                    
Update_Cockroach_Korm-32    200000    35877 ns/op        7.18 s
Update_Cockroach_Pgx-32     200000   283554 ns/op       56.71 s
Update_Postgres_Pgx-32      200000    53921 ns/op       10.78 s
Update_Taran_ORM-32         200000      179 ns/op        0.04 s -- fastest

```

To be fair with korm that optimized for cases when database data size is smaller than RAM (cache set to 100MB)


```
BenchmarkGetAllM_Cockroach_Korm-32   2105178       546.8 ns/op     32 B/op     2 allocs/op
BenchmarkGetAllM_Postgres_Korm-32    2171280       568.3 ns/op     32 B/op     2 allocs/op

BenchmarkGetAllS_Cockroach_Korm-32   1852326       682.0 ns/op    256 B/op     3 allocs/op
BenchmarkGetAllS_Postgres_Korm-32    1690938       703.8 ns/op    256 B/op     3 allocs/op

BenchmarkGetRowM_Cockroach_Korm-32     52078     19679 ns/op     1655 B/op    38 allocs/op
BenchmarkGetRowM_Postgres_Korm-32      92900     11864 ns/op     1819 B/op    42 allocs/op

BenchmarkGetRowS_Cockroach_Korm-32      9546    140996 ns/op     3004 B/op    73 allocs/op
BenchmarkGetRowS_Postgres_Korm-32       7900    133152 ns/op     3005 B/op    73 allocs/op

BenchmarkInsertS_Postgres_Korm-32     100000     84714 ns/op       8.47 s
BenchmarkUpdate_Cockroach_Korm-32     200000     33460 ns/op       6.69 s
```

## Conclusion

Tarantool fastest for insert, update, get single row use-case, postgres with pgx fastest for get multi-row use-case.
