package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fn/fnReflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](
	ctx context.Context,
	filter Filter,
	sorter ...Sorter,
) (res *T, err error) {
	var f bson.M
	if f, err = filter.GetFilter(); err != nil {
		return
	}

	var opt = &options.FindOneOptions{}
	if len(sorter) != 0 && sorter[0] != nil {
		if opt.Sort, err = sorter[0].GetSorter(); err != nil {
			return
		}
	}

	var col = GetDBP(ctx).Collection(filter.GetColNm())
	var cur *mongo.SingleResult
	if cur = col.FindOne(ctx, f, opt); cur.Err() != nil {
		err = cur.Err()
		return
	}

	res = new(T)
	if err = cur.Decode(res); err != nil {
		return
	}

	return
}

func FindAll[T any](
	ctx context.Context,
	filter Filter,
	sorter ...Sorter,
) (res []*T, err error) {
	var f bson.M
	if f, err = filter.GetFilter(); err != nil {
		return
	}

	var opt = &options.FindOptions{}
	if len(sorter) != 0 && sorter[0] != nil {
		if opt.Sort, err = sorter[0].GetSorter(); err != nil {
			return
		}
	}

	var col = GetDBP(ctx).Collection(filter.GetColNm())
	var cur *mongo.Cursor
	if cur, err = col.Find(ctx, f, opt); err != nil {
		return
	}

	res = make([]*T, 0)
	if err = cur.All(ctx, &res); err != nil {
		return
	}

	return
}

func FindList[T any](
	ctx context.Context,
	filter Filter,
	pager Pager,
	sorter ...Sorter,
) (res *ResultList[T], err error) {
	var f bson.M
	if f, err = filter.GetFilter(); err != nil {
		return
	}

	var col = GetDBP(ctx).Collection(filter.GetColNm())
	var total int64
	if total, err = col.CountDocuments(ctx, f); err != nil {
		return
	}

	var opt = &options.FindOptions{}
	if len(sorter) != 0 && sorter[0] != nil {
		if opt.Sort, err = sorter[0].GetSorter(); err != nil {
			return
		}
	}

	opt.Skip = fnReflect.Pointer(pager.GetPage() * pager.GetSize())
	opt.Limit = fnReflect.Pointer(pager.GetSize())

	var cur *mongo.Cursor
	if cur, err = col.Find(ctx, f, opt); err != nil {
		return
	}

	var ls = make([]*T, 0)
	if err = cur.All(ctx, &ls); err != nil {
		return
	}

	res = &ResultList[T]{
		Page:  pager.GetPage(),
		Size:  pager.GetSize(),
		Total: total,
		List:  ls,
	}

	return
}
