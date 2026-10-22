package rqTest

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedTestTable2OrmHelpers(t *testing.T) {
	q := NewTestTable2(nil)
	assert.NotNil(t, q)
	assert.NotEmpty(t, q.SpaceName())
	assert.NotEmpty(t, q.SqlTableName())
	q.Id = uint64(1)
	q.Content = "sample"
	arr := q.ToArray()
	assert.Len(t, arr, 2)
	assert.NotNil(t, q.ToUpdateArray())
	decoded := (&TestTable2{}).FromArray(arr)
	assert.Equal(t, q.Id, decoded.Id)
	assert.Equal(t, q.Content, decoded.Content)
	decoded2 := (&TestTable2{}).FromUncensoredArray(arr)
	assert.Equal(t, q.Id, decoded2.Id)
	assert.Equal(t, q.Content, decoded2.Content)
	assert.Equal(t, 0, q.IdxId())
	assert.Equal(t, `"id"`, q.SqlId())
	assert.Equal(t, 1, q.IdxContent())
	assert.Equal(t, `"content"`, q.SqlContent())
	_, ok := TestTable2FieldTypeMap[`id`]
	assert.True(t, ok)
	assert.NotEmpty(t, q.UniqueIndexId())
	assert.NotEmpty(t, q.UniqueIndexContent())
}

func TestGeneratedTestTable2DbMethodsPanic(t *testing.T) {
	q := NewTestTable2(&Tt.Adapter{})
	q.Id = uint64(1)
	q.Content = "sample"
	assert.Panics(t, func() { _ = q.FindById() })
	assert.Panics(t, func() { _ = q.FindByContent() })
	assert.Panics(t, func() { _ = q.FindOffsetLimit(0, 1, "") })
	assert.Panics(t, func() { _, _ = q.FindArrOffsetLimit(0, 1, "") })
	assert.Panics(t, func() { _ = q.Total() })
}
