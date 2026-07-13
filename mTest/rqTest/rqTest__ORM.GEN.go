package rqTest

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`korm1/mTest`

	`github.com/tarantool/go-tarantool/v2`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/X`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqTest__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqTest__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqTest__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqTest__ORM.GEN.go
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

// SpaceName returns full package and table name
func (t *TestTable2) SpaceName() string { //nolint:dupl false positive
	return string(mTest.TableTestTable2) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *TestTable2) SqlTableName() string { //nolint:dupl false positive
	return `"test_table2"`
}

func (t *TestTable2) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *TestTable2) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
		Index(t.UniqueIndexId()).
		Limit(1).
		Iterator(tarantool.IterEq).
		Key(Tt.Uint64Key{I:t.Id}),
	)
	if L.IsError(err, `TestTable2.FindById failed: `+t.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			t.FromArray(row)
			return true
		}
	}
	return false
}

// UniqueIndexContent return unique index name
func (t *TestTable2) UniqueIndexContent() string { //nolint:dupl false positive
	return `content`
}

// FindByContent Find one by Content
func (t *TestTable2) FindByContent() bool { //nolint:dupl false positive
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
		Index(t.UniqueIndexContent()).
		Limit(1).
		Iterator(tarantool.IterEq).
		Key(tarantool.StringKey{S:t.Content}),
	)
	if L.IsError(err, `TestTable2.FindByContent failed: `+t.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			t.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *TestTable2) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "content"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *TestTable2) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "content"
	`
}

// ToUpdateArray generate slice of update command
func (t *TestTable2) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, t.Id).
		Assign(1, t.Content)
}

// IdxId return name of the index
func (t *TestTable2) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *TestTable2) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxContent return name of the index
func (t *TestTable2) IdxContent() int { //nolint:dupl false positive
	return 1
}

// SqlContent return name of the column being indexed
func (t *TestTable2) SqlContent() string { //nolint:dupl false positive
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
func (t *TestTable2) FromArray(ax A.X) *TestTable2 { //nolint:dupl false positive
	t.Id = X.ToU(ax[0])
	t.Content = X.ToS(ax[1])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *TestTable2) FromUncensoredArray(ax A.X) *TestTable2 { //nolint:dupl false positive
	t.Id = X.ToU(ax[0])
	t.Content = X.ToS(ax[1])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *TestTable2) FindOffsetLimit(offset, limit uint32, idx string) []TestTable2 { //nolint:dupl false positive
	var rows []TestTable2
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
		Index(idx).
		Offset(offset).
		Limit(limit).
		Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `TestTable2.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := TestTable2{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TestTable2) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := t.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(t.SpaceName()).
		Index(idx).
		Offset(offset).
		Limit(limit).
		Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `TestTable2.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `TestTable2.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	rows = make([]A.X, len(res))
	for _, row := range res {
		row, ok := row.([]any)
		if ok {
			rows = append(rows, row)
		}
	}
	return rows, Tt.QueryMetaFrom(resp, nil)
}

// Total count number of rows
func (t *TestTable2) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TestTable2FieldTypeMap returns key value of field name and key
var TestTable2FieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
	`id`:      Tt.Unsigned,
	`content`: Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

