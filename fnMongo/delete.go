package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fn/fnParam"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteOne(
	ctx context.Context,
	f Filter,
	opts ...*options.DeleteOptions,
) (err error) {
	var filter bson.M
	if filter, err = f.GetFilter(); err != nil {
		return
	}

	var opt = fnParam.Get(opts)

	if _, err = GetColP(ctx, f.GetColNm()).DeleteOne(ctx, filter, opt); err != nil {
		return
	}

	return
}

func DeleteAll(
	ctx context.Context,
	f Filter,
	opts ...*options.DeleteOptions,
) (err error) {
	var filter bson.M
	if filter, err = f.GetFilter(); err != nil {
		return
	}

	var opt = fnParam.Get(opts)

	if _, err = GetColP(ctx, f.GetColNm()).DeleteMany(ctx, filter, opt); err != nil {
		return
	}

	return
}
