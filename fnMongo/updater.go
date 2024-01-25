package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-snippet/fn/fnParam"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne(
	ctx context.Context,
	f Filter,
	u Updater,
	opts ...*options.UpdateOptions,
) (err error) {
	var filter bson.M
	if filter, err = f.GetFilter(); err != nil {
		return
	}

	var updater bson.M
	if updater, err = u.GetUpdater(); err != nil {
		return
	}

	var col = GetColP(ctx, f.GetColNm())
	var opt = fnParam.Get(opts)

	if _, err = col.UpdateOne(ctx, filter, updater, opt); err != nil {
		return
	}

	return
}

func UpdateMany(
	ctx context.Context,
	f Filter,
	u Updater,
	opts ...*options.UpdateOptions,
) (err error) {
	var filter bson.M
	if filter, err = f.GetFilter(); err != nil {
		return
	}

	var updater bson.M
	if updater, err = u.GetUpdater(); err != nil {
		return
	}

	var col = GetColP(ctx, f.GetColNm())
	var opt = fnParam.Get(opts)

	if _, err = col.UpdateMany(ctx, filter, updater, opt); err != nil {
		return
	}

	return
}
