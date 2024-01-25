package fnMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	CodecRegister func(*bsoncodec.Registry) *bsoncodec.Registry

	IfConnectClientArgs interface {
		GetHost() string
		GetUsername() string
		GetPassword() string
		GetCodecRegisters() []CodecRegister
	}

	IfConnectDatabaseArgs interface {
		IfConnectClientArgs
		GetDatabase() string
	}
)

func createOpt(i IfConnectClientArgs) (opt *options.ClientOptions) {
	opt = options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s", i.GetHost())).
		SetReadConcern(readconcern.Majority()).
		SetWriteConcern(writeconcern.Majority()).
		SetAuth(options.Credential{
			Username: i.GetUsername(),
			Password: i.GetPassword(),
		}).
		SetBSONOptions(&options.BSONOptions{
			UseLocalTimeZone: false,
		})

	opt.Registry = bson.DefaultRegistry

	if len(i.GetCodecRegisters()) != 0 {
		for _, registry := range i.GetCodecRegisters() {
			opt.Registry = registry(opt.Registry)
		}
	}

	return
}

func ConnectClient(i IfConnectClientArgs) (client *mongo.Client, err error) {
	var ctx = context.TODO()
	if client, err = mongo.Connect(ctx, createOpt(i)); err != nil {
		return
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}

	return
}

func ConnectDB(i IfConnectDatabaseArgs) (db *mongo.Database, err error) {
	var client *mongo.Client
	if client, err = ConnectClient(i); err != nil {
		return
	}
	db = client.Database(i.GetDatabase())
	return
}

/*------------------------------------------------------------------------------------------------*/

type ConnectDatabaseArgs struct {
	Host           string
	Username       string
	Password       string
	CodecRegisters []CodecRegister
	Database       string
}

func (x *ConnectDatabaseArgs) GetHost() string {
	return x.Host
}

func (x *ConnectDatabaseArgs) GetUsername() string {
	return x.Username
}

func (x *ConnectDatabaseArgs) GetPassword() string {
	return x.Password
}

func (x *ConnectDatabaseArgs) GetCodecRegisters() []CodecRegister {
	if len(x.CodecRegisters) != 0 {
		return x.CodecRegisters
	}
	return make([]CodecRegister, 0)
}

func (x *ConnectDatabaseArgs) GetDatabase() string {
	return x.Database
}

/*------------------------------------------------------------------------------------------------*/

type ConnectClientArgs struct {
	Host           string
	Username       string
	Password       string
	CodecRegisters []CodecRegister
}

func (x *ConnectClientArgs) GetHost() string {
	return x.Host
}

func (x *ConnectClientArgs) GetUsername() string {
	return x.Username
}

func (x *ConnectClientArgs) GetPassword() string {
	return x.Password
}

func (x *ConnectClientArgs) GetCodecRegisters() []CodecRegister {
	if len(x.CodecRegisters) != 0 {
		return x.CodecRegisters
	}
	return make([]CodecRegister, 0)
}
