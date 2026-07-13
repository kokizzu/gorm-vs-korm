package mTest

import (
	"context"
	"time"

	. "github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/tarantool/go-tarantool/v2"
)

func Migrate(taran *Adapter) {
	taran.MigrateTables(Tables)
}

func ConnectTarantool() *tarantool.Connection {
	taran, err := tarantool.Connect(context.Background(), tarantool.NetDialer{
		Address: `localhost:3301`,
	}, tarantool.Opts{
		Timeout: 8 * time.Second,
	})
	L.PanicIf(err, `failed to connect to tarantool`)
	return taran
}
