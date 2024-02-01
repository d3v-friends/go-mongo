package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-mongo/mongoCodec"
	"github.com/d3v-friends/go-tools/fn/fnEnv"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"testing"
)

func TestMigrate(test *testing.T) {
	fnPanic.On(fnEnv.Load("../.env"))
	var client = fnPanic.Get(ConnectClient(&ConnectClientArgs{
		Host:     fnEnv.GetString("MG_HOST"),
		Username: fnEnv.GetString("MG_USERNAME"),
		Password: fnEnv.GetString("MG_PASSWORD"),
		CodecRegisters: []CodecRegister{
			mongoCodec.DecimalRegistry,
		},
	}))

	test.Run("migrate", func(t *testing.T) {
		var ctx = context.TODO()
		ctx = SetDB(ctx, client.Database(fnEnv.GetString("MG_DATABASE")))
		var err = RunMigrate(ctx, MangoModel)
		if err != nil {
			t.Fatal(err)
		}
	})
}
