package rqTest

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`korm1/mTest`

	`github.com/tarantool/go-tarantool`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/X`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqTest__ORM.GEN.go
//go:generate replacer 'Id" form' 'Id,string" form' type rqTest__ORM.GEN.go
//go:generate replacer 'json:"id"' 'json:"id,string"' type rqTest__ORM.GEN.go
//go:generate replacer 'By" form' 'By,string" form' type rqTest__ORM.GEN.go
// go:generate msgp -tests=false -file rqTest__ORM.GEN.go -o rqTest__MSG.GEN.go

// TestTable2 DAO reader/query struct
type TestTable2 struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id      uint64
	Content string
}

// NewTestTable2 create new ORM reader/query object
func NewTestTable2(adapter *Tt.Adapter) *TestTable2 {
	return &TestTable2{Adapter: adapter}
}

// sqlTableName returns full package and table name
func (t *TestTable2) SpaceName() string { //nolint:dupl false positive
	return string(mTest.TableTestTable2)
}

// sqlTableName returns quoted table name
func (t *TestTable2) sqlTableName() string { //nolint:dupl false positive
	return `"test_table2"`
}

func (t *TestTable2) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *TestTable2) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{t.Id})
	if L.IsError(err, `TestTable2.FindById failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexContent return unique index name
func (t *TestTable2) UniqueIndexContent() string { //nolint:dupl false positive
	return `content`
}

// FindByContent Find one by Content
func (t *TestTable2) FindByContent() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexContent(), 0, 1, tarantool.IterEq, A.X{t.Content})
	if L.IsError(err, `TestTable2.FindByContent failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// sqlSelectAllFields generate sql select fields
func (t *TestTable2) sqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "content"
	`
}

// ToUpdateArray generate slice of update command
func (t *TestTable2) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.Content},
	}
}

// IdxId return name of the index
func (t *TestTable2) IdxId() int { //nolint:dupl false positive
	return 0
}

// sqlIdxId return name of the column being indexed
func (t *TestTable2) sqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxContent return name of the index
func (t *TestTable2) IdxContent() int { //nolint:dupl false positive
	return 1
}

// sqlIdxContent return name of the column being indexed
func (t *TestTable2) sqlContent() string { //nolint:dupl false positive
	return `"content"`
}

// ToArray receiver fields to slice
func (t *TestTable2) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.Content, // 1
	}
}

// FromArray convert slice to receiver fields
func (t *TestTable2) FromArray(a A.X) *TestTable2 { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.Content = X.ToS(a[1])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *TestTable2) FindOffsetLimit(offset, limit uint32, idx string) []TestTable2 { //nolint:dupl false positive
	var rows []TestTable2
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, 2, A.X{})
	if L.IsError(err, `TestTable2.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := TestTable2{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TestTable2) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, 2, A.X{})
	if L.IsError(err, `TestTable2.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (t *TestTable2) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

