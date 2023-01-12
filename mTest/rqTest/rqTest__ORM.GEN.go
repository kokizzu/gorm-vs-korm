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

type TestTable2 struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id      uint64
	Content string
}

func NewTestTable2(adapter *Tt.Adapter) *TestTable2 {
	return &TestTable2{Adapter: adapter}
}

func (t *TestTable2) SpaceName() string { //nolint:dupl false positive
	return string(mTest.TableTestTable2)
}

func (t *TestTable2) sqlTableName() string { //nolint:dupl false positive
	return `"test_table2"`
}

func (t *TestTable2) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

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

func (t *TestTable2) sqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "content"
	`
}

func (t *TestTable2) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.Content},
	}
}

func (t *TestTable2) IdxId() int { //nolint:dupl false positive
	return 0
}

func (t *TestTable2) sqlId() string { //nolint:dupl false positive
	return `"id"`
}

func (t *TestTable2) IdxContent() int { //nolint:dupl false positive
	return 1
}

func (t *TestTable2) sqlContent() string { //nolint:dupl false positive
	return `"content"`
}

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

func (t *TestTable2) FromArray(a A.X) *TestTable2 { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.Content = X.ToS(a[1])
	return t
}

func (t *TestTable2) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

