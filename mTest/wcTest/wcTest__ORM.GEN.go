package wcTest

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`korm1/mTest/rqTest`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/X`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcTest__ORM.GEN.go
//go:generate replacer 'Id" form' 'Id,string" form' type wcTest__ORM.GEN.go
//go:generate replacer 'json:"id"' 'json:"id,string"' type wcTest__ORM.GEN.go
//go:generate replacer 'By" form' 'By,string" form' type wcTest__ORM.GEN.go
// go:generate msgp -tests=false -file wcTest__ORM.GEN.go -o wcTest__MSG.GEN.go

type TestTable2Mutator struct {
	rqTest.TestTable2
	mutations []A.X
}

func NewTestTable2Mutator(adapter *Tt.Adapter) *TestTable2Mutator {
	return &TestTable2Mutator{TestTable2: rqTest.TestTable2{Adapter: adapter}}
}

func (t *TestTable2Mutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// Overwrite all columns, error if not exists
func (t *TestTable2Mutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.ToUpdateArray())
	return !L.IsError(err, `TestTable2.DoOverwriteById failed: `+t.SpaceName())
}

// Update only mutated, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TestTable2Mutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.mutations)
	return !L.IsError(err, `TestTable2.DoUpdateById failed: `+t.SpaceName())
}

func (t *TestTable2Mutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id})
	return !L.IsError(err, `TestTable2.DoDeletePermanentById failed: `+t.SpaceName())
}

// func (t *TestTable2Mutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := t.Adapter.Upsert(t.SpaceName(), t.ToArray(), A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.Content},
//	})
//	return !L.IsError(err, `TestTable2.DoUpsert failed: `+t.SpaceName())
// }

// Overwrite all columns, error if not exists
func (t *TestTable2Mutator) DoOverwriteByContent() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexContent(), A.X{t.Content}, t.ToUpdateArray())
	return !L.IsError(err, `TestTable2.DoOverwriteByContent failed: `+t.SpaceName())
}

// Update only mutated, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TestTable2Mutator) DoUpdateByContent() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexContent(), A.X{t.Content}, t.mutations)
	return !L.IsError(err, `TestTable2.DoUpdateByContent failed: `+t.SpaceName())
}

func (t *TestTable2Mutator) DoDeletePermanentByContent() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexContent(), A.X{t.Content})
	return !L.IsError(err, `TestTable2.DoDeletePermanentByContent failed: `+t.SpaceName())
}

// insert, error if exists
func (t *TestTable2Mutator) DoInsert() bool { //nolint:dupl false positive
	row, err := t.Adapter.Insert(t.SpaceName(), t.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TestTable2.DoInsert failed: `+t.SpaceName())
}

// replace = upsert, only error when there's unique secondary key
func (t *TestTable2Mutator) DoReplace() bool { //nolint:dupl false positive
	_, err := t.Adapter.Replace(t.SpaceName(), t.ToArray())
	return !L.IsError(err, `TestTable2.DoReplace failed: `+t.SpaceName())
}

func (t *TestTable2Mutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.Id = val
		return true
	}
	return false
}

func (t *TestTable2Mutator) SetContent(val string) bool { //nolint:dupl false positive
	if val != t.Content {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.Content = val
		return true
	}
	return false
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

