package wcTest

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`context`
	`fmt`
	`log`
	`os`
	`testing`
	`time`
	`korm1/mTest`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/S`
	`github.com/ory/dockertest/v3`
	`github.com/stretchr/testify/assert`
	`github.com/tarantool/go-tarantool/v2`
)

var globalPool *dockertest.Pool
var reconnect func() *tarantool.Connection
var dbConn *tarantool.Connection

func prepareDb(onReady func(db *tarantool.Connection) int) {
	const dockerRepo = "tarantool/tarantool"
	const dockerVer = Tt.LatestTarantoolDockerTag
	const ttPort = "3301/tcp"
	const dbConnStr = "127.0.0.1:%s"
	const dbUser = "guest"
	const dbPass = ""
	var err error
	if globalPool == nil {
		globalPool, err = dockertest.NewPool("")
		if err != nil {
			log.Printf("Could not connect to docker: %s\n", err)
			os.Exit(onReady(nil))
		}
	}
	resource, err := globalPool.Run(dockerRepo, dockerVer, []string{})
	if err != nil {
		log.Printf("Could not start resource: %s\n", err)
		os.Exit(onReady(nil))
	}
	var db *tarantool.Connection
	if err := globalPool.Retry(func() error {
		var err error
		connStr := fmt.Sprintf(dbConnStr, resource.GetPort(ttPort))
		reconnect = func() *tarantool.Connection {
			db, err = tarantool.Connect(context.Background(), tarantool.NetDialer{
				Address:  connStr,
				User:     dbUser,
				Password: dbPass,
			}, tarantool.Opts{
				Timeout: 8 * time.Second,
			})
			if err != nil && !S.Contains(err.Error(), "failed to read greeting: EOF") {
				L.IsError(err, "tarantool.Connect")
			}
			return db
		}
		reconnect()
		if err != nil {
			return err
		}
		_, err = db.Do(tarantool.NewPingRequest()).Get()
		return err
	}); err != nil {
		log.Printf("Could not connect to docker: %s\n", err)
		os.Exit(onReady(nil))
	}
	code := onReady(db)
	if err := globalPool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestMain(m *testing.M) {
	prepareDb(func(db *tarantool.Connection) int {
		dbConn = db
		return m.Run()
	})
}

func TestGeneratedSanity(t *testing.T) {
	if dbConn == nil {
		t.Skip("docker unavailable")
	}
	conn := dbConn
	a := &Tt.Adapter{Connection: conn, Reconnect: reconnect}
	a.MigrateTables(map[Tt.TableName]*Tt.TableProp{
		"bands": {
			Fields: []Tt.Field{
				{"id", Tt.Unsigned},
				{"band_name", Tt.String},
				{"year", Tt.Unsigned},
			},
			AutoIncrementId: true,
			Unique1:         "band_name",
			Indexes:         []string{"year"},
		},
	})
	tuples := [][]any{
		{1, "Roxette", 1986},
		{2, "Scorpions", 1965},
		{3, "Ace of Base", 1987},
		{4, "The Beatles", 1960},
	}
	for _, tuple := range tuples {
		_, err := conn.Do(tarantool.NewInsertRequest("bands").Tuple(tuple)).Get()
		assert.NoError(t, err)
	}
	_, err := conn.Do(tarantool.NewSelectRequest("bands").Limit(10).Iterator(tarantool.IterEq).Key([]any{uint(1)})).Get()
	assert.NoError(t, err)
	_, err = conn.Do(tarantool.NewSelectRequest("bands").Index("band_name").Limit(10).Iterator(tarantool.IterEq).Key([]any{"The Beatles"})).Get()
	assert.NoError(t, err)
	_, err = conn.Do(tarantool.NewUpdateRequest("bands").Key(tarantool.IntKey{2}).Operations(tarantool.NewOperations().Assign(1, "Pink Floyd"))).Get()
	assert.NoError(t, err)
	_, err = conn.Do(tarantool.NewUpsertRequest("bands").Tuple([]any{uint(5), "The Rolling Stones", 1962}).Operations(tarantool.NewOperations().Assign(1, "The Doors"))).Get()
	assert.NoError(t, err)
	_, err = conn.Do(tarantool.NewReplaceRequest("bands").Tuple([]any{1, "Queen", 1970})).Get()
	assert.NoError(t, err)
	_, err = conn.Do(tarantool.NewDeleteRequest("bands").Key([]any{uint(5)})).Get()
	assert.NoError(t, err)
}

func TestGeneratedTestTable2Unit(t *testing.T) {
	m := NewTestTable2Mutator(nil)
	assert.NotNil(t, m)
	assert.False(t, m.HaveMutation())
	assert.Empty(t, m.Logs())
	assert.True(t, m.SetId(uint64(1)))
	assert.False(t, m.SetId(uint64(1)))
	assert.True(t, m.SetContent("sample"))
	assert.False(t, m.SetContent("sample"))
	assert.True(t, m.HaveMutation())
	from := m.TestTable2
	assert.True(t, m.SetAll(from, nil, nil))
	m2 := NewTestTable2Mutator(nil)
	fromZero := m2.TestTable2
	assert.False(t, m.SetAll(fromZero, nil, nil))
	m.ClearMutations()
	assert.False(t, m.HaveMutation())
	assert.True(t, m.DoUpdateById())
	assert.True(t, m.DoUpdateByContent())
}

func TestGeneratedTestTable2CRUD(t *testing.T) {
	if dbConn == nil {
		t.Skip("docker unavailable")
	}
	a := &Tt.Adapter{Connection: dbConn, Reconnect: reconnect}
	ok := a.UpsertTable(mTest.TableTestTable2, mTest.TarantoolTables[mTest.TableTestTable2])
	assert.True(t, ok)
	_ = a.TruncateTable(string(mTest.TableTestTable2))
	seed := func() *TestTable2Mutator {
		x := NewTestTable2Mutator(a)
		x.Content = "sample"
		assert.True(t, x.DoInsert())
		assert.Greater(t, x.Id, uint64(0))
		return x
	}
	rec0 := seed()
	readInsert0 := NewTestTable2Mutator(a)
	readInsert0.Id = rec0.Id
	assert.True(t, readInsert0.FindById())
	assert.Equal(t, rec0.Id, readInsert0.Id)
	assert.Equal(t, rec0.Content, readInsert0.Content)
	assert.True(t, rec0.DoUpdateById())
	readUpdate0 := NewTestTable2Mutator(a)
	readUpdate0.Id = rec0.Id
	assert.True(t, readUpdate0.FindById())
	assert.Equal(t, rec0.Id, readUpdate0.Id)
	assert.Equal(t, rec0.Content, readUpdate0.Content)
	assert.True(t, rec0.DoOverwriteById())
	readOverwrite0 := NewTestTable2Mutator(a)
	readOverwrite0.Id = rec0.Id
	assert.True(t, readOverwrite0.FindById())
	assert.Equal(t, rec0.Id, readOverwrite0.Id)
	assert.Equal(t, rec0.Content, readOverwrite0.Content)
	rows0 := rec0.FindOffsetLimit(0, 10, rec0.UniqueIndexId())
	assert.NotNil(t, rows0)
	arrRows0, _ := rec0.FindArrOffsetLimit(0, 10, rec0.UniqueIndexId())
	assert.NotNil(t, arrRows0)
	assert.GreaterOrEqual(t, rec0.Total(), int64(0))
	assert.True(t, rec0.DoDeletePermanentById())
	assert.False(t, rec0.FindById())
	rec1 := seed()
	readInsert1 := NewTestTable2Mutator(a)
	readInsert1.Content = rec1.Content
	assert.True(t, readInsert1.FindByContent())
	assert.Equal(t, rec1.Id, readInsert1.Id)
	assert.Equal(t, rec1.Content, readInsert1.Content)
	assert.True(t, rec1.DoUpdateByContent())
	readUpdate1 := NewTestTable2Mutator(a)
	readUpdate1.Content = rec1.Content
	assert.True(t, readUpdate1.FindByContent())
	assert.Equal(t, rec1.Id, readUpdate1.Id)
	assert.Equal(t, rec1.Content, readUpdate1.Content)
	assert.True(t, rec1.DoOverwriteByContent())
	readOverwrite1 := NewTestTable2Mutator(a)
	readOverwrite1.Content = rec1.Content
	assert.True(t, readOverwrite1.FindByContent())
	assert.Equal(t, rec1.Id, readOverwrite1.Id)
	assert.Equal(t, rec1.Content, readOverwrite1.Content)
	rows1 := rec1.FindOffsetLimit(0, 10, rec1.UniqueIndexContent())
	assert.NotNil(t, rows1)
	arrRows1, _ := rec1.FindArrOffsetLimit(0, 10, rec1.UniqueIndexContent())
	assert.NotNil(t, arrRows1)
	assert.GreaterOrEqual(t, rec1.Total(), int64(0))
	assert.True(t, rec1.DoDeletePermanentByContent())
	assert.False(t, rec1.FindByContent())
	deletedCheck1 := NewTestTable2Mutator(a)
	deletedCheck1.Id = rec1.Id
	assert.False(t, deletedCheck1.FindById())
	u := NewTestTable2Mutator(a)
	u.Content = "sample"
	assert.True(t, u.DoUpsertById())
	assert.Greater(t, u.Id, uint64(0))
	uRead := NewTestTable2Mutator(a)
	uRead.Id = u.Id
	assert.True(t, uRead.FindById())
	assert.Equal(t, u.Id, uRead.Id)
	assert.Equal(t, u.Content, uRead.Content)
	assert.True(t, u.DoDeletePermanentById())
}

func TestGeneratedTestTable2LargeUnsignedRoundTrip(t *testing.T) {
	if dbConn == nil {
		t.Skip("docker unavailable")
	}
	a := &Tt.Adapter{Connection: dbConn, Reconnect: reconnect}
	ok := a.UpsertTable(mTest.TableTestTable2, mTest.TarantoolTables[mTest.TableTestTable2])
	assert.True(t, ok)
	_ = a.TruncateTable(string(mTest.TableTestTable2))
	const largeUnsignedId = uint64(1<<63) + 77
	const largeUnsignedValue = uint64(1<<63) + 123
	const largeSigned = int64(1<<62) + 1234
	const largeSignedUpdated = int64(1<<62) + 5678
	rec := NewTestTable2Mutator(a)
	rec.Id = largeUnsignedId
	rec.Content = "sample"
	assert.True(t, rec.DoInsert())
	assert.Equal(t, largeUnsignedId, rec.Id)
	rawRows, err := a.RetryDo(
		tarantool.NewSelectRequest(rec.SpaceName()).
			Index(rec.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(Tt.Uint64Key{I: largeUnsignedId}),
	)
	assert.NoError(t, err)
	if assert.Len(t, rawRows, 1) {
		row, ok := rawRows[0].([]any)
		if assert.True(t, ok) && assert.NotEmpty(t, row) {
			assert.IsType(t, uint64(0), row[0])
			assert.Equal(t, largeUnsignedId, row[0])
		}
	}
	readInsert := NewTestTable2Mutator(a)
	readInsert.Id = largeUnsignedId
	assert.True(t, readInsert.FindById())
	assert.Equal(t, rec.Id, readInsert.Id)
	assert.Equal(t, rec.Content, readInsert.Content)
	assert.True(t, rec.DoOverwriteById())
	readOverwrite := NewTestTable2Mutator(a)
	readOverwrite.Id = largeUnsignedId
	assert.True(t, readOverwrite.FindById())
	assert.Equal(t, rec.Id, readOverwrite.Id)
	assert.Equal(t, rec.Content, readOverwrite.Content)
	assert.True(t, rec.DoDeletePermanentById())
	deleted := NewTestTable2Mutator(a)
	deleted.Id = largeUnsignedId
	assert.False(t, deleted.FindById())
}

