package mTest

import (
	"os"
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
)

// only need to do once before compile
func TestGenerateOrm(t *testing.T) {
	if os.Getenv("RUN_TARANTOOL_GENERATOR") != "1" {
		t.Skip("use make generate-orm to regenerate Tarantool ORM files")
	}
	Tt.GenerateOrm(Tables, false)
	t.SkipNow()
}
func TestMigrate(t *testing.T) {
	if os.Getenv("RUN_TARANTOOL_INTEGRATION") != "1" {
		t.Skip("set RUN_TARANTOOL_INTEGRATION=1 to run Tarantool migration test")
	}
	taran := &Tt.Adapter{Connection: ConnectTarantool(), Reconnect: ConnectTarantool}
	_, err := taran.Ping()
	L.PanicIf(err, `taran.Ping`)
	Migrate(taran)
}
