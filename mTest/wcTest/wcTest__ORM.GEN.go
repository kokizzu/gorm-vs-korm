package wcTest

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`korm1/mTest/rqTest`

	`github.com/tarantool/go-tarantool/v2`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/M`
	`github.com/kokizzu/gotro/S`
	`github.com/kokizzu/gotro/X`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcTest__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcTest__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcTest__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcTest__ORM.GEN.go
// TestTable2Mutator DAO writer/command struct
type TestTable2Mutator struct {
	rqTest.TestTable2
	mutations *tarantool.Operations
	logs	  []A.X
}

// NewTestTable2Mutator create new ORM writer/command object
func NewTestTable2Mutator(adapter *Tt.Adapter) (res *TestTable2Mutator) {
	res = &TestTable2Mutator{TestTable2: rqTest.TestTable2{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (t *TestTable2Mutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TestTable2Mutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TestTable2Mutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = tarantool.NewOperations()
	t.logs = []A.X{}
}

// toOverwriteArray generate slice of update command without primary index fields
func (t *TestTable2Mutator) toOverwriteArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(1, t.Content)
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (t *TestTable2Mutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.RetryDo(tarantool.NewUpdateRequest(t.SpaceName()).
		Index(t.UniqueIndexId()).
		Key(Tt.Uint64Key{I:t.Id}).
		Operations(t.toOverwriteArray()),
	)
	return !L.IsError(err, `TestTable2.DoOverwriteById failed: `+t.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TestTable2Mutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.RetryDo(
		tarantool.NewUpdateRequest(t.SpaceName()).
		Index(t.UniqueIndexId()).
		Key(Tt.Uint64Key{I:t.Id}).
		Operations(t.mutations),
	)
	return !L.IsError(err, `TestTable2.DoUpdateById failed: `+t.SpaceName())
}

// DoDeletePermanentById permanent delete
func (t *TestTable2Mutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.RetryDo(
		tarantool.NewDeleteRequest(t.SpaceName()).
		Index(t.UniqueIndexId()).
		Key(Tt.Uint64Key{I:t.Id}),
	)
	return !L.IsError(err, `TestTable2.DoDeletePermanentById failed: `+t.SpaceName())
}

// DoOverwriteByContent update all columns, error if not exists, not using mutations/Set*
func (t *TestTable2Mutator) DoOverwriteByContent() bool { //nolint:dupl false positive
	_, err := t.Adapter.RetryDo(tarantool.NewUpdateRequest(t.SpaceName()).
		Index(t.UniqueIndexContent()).
		Key(tarantool.StringKey{S:t.Content}).
		Operations(t.toOverwriteArray()),
	)
	return !L.IsError(err, `TestTable2.DoOverwriteByContent failed: `+t.SpaceName())
}

// DoUpdateByContent update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TestTable2Mutator) DoUpdateByContent() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.RetryDo(
		tarantool.NewUpdateRequest(t.SpaceName()).
		Index(t.UniqueIndexContent()).
		Key(tarantool.StringKey{S:t.Content}).
		Operations(t.mutations),
	)
	return !L.IsError(err, `TestTable2.DoUpdateByContent failed: `+t.SpaceName())
}

// DoDeletePermanentByContent permanent delete
func (t *TestTable2Mutator) DoDeletePermanentByContent() bool { //nolint:dupl false positive
	_, err := t.Adapter.RetryDo(
		tarantool.NewDeleteRequest(t.SpaceName()).
		Index(t.UniqueIndexContent()).
		Key(tarantool.StringKey{S:t.Content}),
	)
	return !L.IsError(err, `TestTable2.DoDeletePermanentByContent failed: `+t.SpaceName())
}

// DoInsert insert, error if already exists
func (t *TestTable2Mutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.RetryDo(
		tarantool.NewInsertRequest(t.SpaceName()).
		Tuple(arr),
	)
	if err == nil {
		if len(row) > 0 {
			if cells, ok := row[0].([]any); ok && len(cells) > 0 {
				t.Id = X.ToU(cells[0])
			}
		}
	}
	return !L.IsError(err, `TestTable2.DoInsert failed: `+t.SpaceName() + `\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (t *TestTable2Mutator) DoUpsertById() bool { //nolint:dupl false positive
	if t.Id > 0 {
		return t.DoUpdateById()
	}
	return t.DoInsert()
}

// SetId create mutations, should not duplicate
func (t *TestTable2Mutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations.Assign(0, val)
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetContent create mutations, should not duplicate
func (t *TestTable2Mutator) SetContent(val string) bool { //nolint:dupl false positive
	if val != t.Content {
		t.mutations.Assign(1, val)
		t.logs = append(t.logs, A.X{`content`, t.Content, val})
		t.Content = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TestTable2Mutator) SetAll(from rqTest.TestTable2, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		t.Id = from.Id
		changed = true
	}
	if !excludeMap[`content`] && (forceMap[`content`] || from.Content != ``) {
		t.Content = S.Trim(from.Content)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

