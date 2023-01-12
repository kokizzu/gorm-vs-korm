
# Gorm+Sqlite vs Korm+Sqlite vs Tarantool

```shell
goos: linux
goarch: amd64
pkg: korm1/mTest
cpu: AMD Ryzen 9 5950X 16-Core Processor
BenchmarkGetAllS_Sqlite_GORM-32    78362             14981 ns/op            3697 B/op         67 allocs/op
BenchmarkGetAllM_Sqlite_GORM-32    10000            446332 ns/op         1668883 B/op      35111 allocs/op
BenchmarkGetRowS_Sqlite_GORM-32    65410             17721 ns/op            4406 B/op         91 allocs/op
BenchmarkGetRowM_Sqlite_GORM-32    60702             18471 ns/op            5037 B/op        109 allocs/op
BenchmarkGetAllS_Sqlite_Korm-32   137373             10111 ns/op            1719 B/op         47 allocs/op
BenchmarkGetAllM_Sqlite_Korm-32   434544              2452 ns/op            1006 B/op         28 allocs/op
BenchmarkGetRowS_Sqlite_Korm-32   120147             10639 ns/op            2116 B/op         61 allocs/op
BenchmarkGetRowM_Sqlite_Korm-32   321590              3270 ns/op            1401 B/op         42 allocs/op
BenchmarkGetAllS_Taran-32         162075              7227 ns/op            2432 B/op         54 allocs/op
BenchmarkGetAllM_Taran-32         154174              7493 ns/op            2744 B/op         55 allocs/op
BenchmarkGetRowS_Taran-32         150128              8001 ns/op            2522 B/op         56 allocs/op
BenchmarkGetRowM_Taran-32         148023              8037 ns/op            2490 B/op         54 allocs/op
PASS
ok      korm1/mTest     18.601s
```

Note:
- disabled cache for KORM
- tarantool not using ORM nor faster API, just raw sql query
- tarantool 10x less rps when only 1 core utilized (without conc)