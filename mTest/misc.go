package mTest

import (
	"strings"

	"github.com/kokizzu/gotro/A"
	. "github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/tarantool/go-tarantool"
)

func Migrate(taran *Adapter) {
	for name, props := range Tables {
		UpsertTable(taran, name, props) // copy of gotro/D/Tt/tarantool_migrator.go
	}
}

func UpsertTable(a *Adapter, tableName TableName, prop *TableProp) bool {
	if prop.Engine == `` {
		prop.Engine = Vinyl
	}
	if !a.CreateSpace(string(tableName), prop.Engine) {
		return false
	}
	if !a.ReformatTable(string(tableName), prop.Fields) {
		return false // failed to create table
	}
	// create one field unique index
	a.ExecBoxSpace(string(tableName)+`:format`, A.X{})
	if prop.AutoIncrementId {
		if len(prop.Fields) < 1 || prop.Fields[0].Name != IdCol || prop.Fields[0].Type != Unsigned {
			panic(`must create Unsigned id field on first field to use AutoIncrementId`)
		}

		seqName := string(tableName) + `_` + IdCol
		a.ExecTarantoolVerbose(`box.schema.sequence.create`, A.X{
			seqName,
		})
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			IdCol, Index{
				Sequence:    seqName,
				Parts:       []string{IdCol},
				IfNotExists: true,
				Unique:      true,
			},
		})
	}
	// only create unique if not "id"
	if prop.Unique1 != `` && !(prop.AutoIncrementId && prop.Unique1 == IdCol) {
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Unique1, Index{Parts: []string{prop.Unique1}, IfNotExists: true, Unique: true},
		})
		if prop.Unique2 != `` && prop.Unique1 == prop.Unique2 {
			panic(`Unique1 and Unique2 must be unique`)
		}
		if prop.Unique3 != `` && prop.Unique1 == prop.Unique3 {
			panic(`Unique1 and Unique3 must be unique`)
		}
	}
	if prop.Unique2 != `` && !(prop.AutoIncrementId && prop.Unique2 == IdCol) {
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Unique2, Index{Parts: []string{prop.Unique2}, IfNotExists: true, Unique: true},
		})
		if prop.Unique3 != `` && prop.Unique2 == prop.Unique3 {
			panic(`Unique2 and Unique3 must be unique`)
		}
	}
	if prop.Unique3 != `` && !(prop.AutoIncrementId && prop.Unique3 == IdCol) {
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Unique3, Index{Parts: []string{prop.Unique3}, IfNotExists: true, Unique: true},
		})
	}
	// create multi-field unique index: [col1, col2] will named col1__col2
	if len(prop.Uniques) > 1 {
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			strings.Join(prop.Uniques, `__`), Index{Parts: prop.Uniques, IfNotExists: true, Unique: true},
		})
	}
	// create other indexes
	for _, index := range prop.Indexes {
		//a.ExecBoxSpace(tableName+`.index.`+index+`:drop`, AX{index}) // TODO: remove this when index fixed
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			index, Index{Parts: []string{index}, IfNotExists: true},
		})
	}
	return true
}

func ConnectTarantool() *tarantool.Connection {
	taran, err := tarantool.Connect(`localhost:3301`, tarantool.Opts{})
	L.PanicIf(err, `failed to connect to tarantool`)
	return taran
}
