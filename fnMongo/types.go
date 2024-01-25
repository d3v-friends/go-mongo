package fnMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Pager interface {
		GetSize() int64
		GetPage() int64
	}

	Indexer interface {
		GetIndex() []mongo.IndexModel
	}

	Filter interface {
		GetFilter() (bson.M, error)
		GetColNm() string
	}

	Sorter interface {
		GetSorter() (bson.M, error)
	}

	Updater interface {
		GetUpdater() (bson.M, error)
	}

	ResultList[T any] struct {
		Page  int64
		Size  int64
		Total int64
		List  []*T
	}

	MongoCodec interface {
		bsoncodec.ValueEncoder
		bsoncodec.ValueDecoder
	}
)

/*------------------------------------------------------------------------------------------------*/

type IdxDir int

func (x IdxDir) Int() int {
	return int(x)
}

func (x IdxDir) Valid() bool {
	for _, dir := range IdxDirAll {
		if x == dir {
			return true
		}
	}
	return false
}

var (
	IdxDirASC  IdxDir = 1
	IdxDirDESC IdxDir = -1
)

var IdxDirAll = []IdxDir{
	IdxDirASC,
	IdxDirDESC,
}

/*------------------------------------------------------------------------------------------------*/

// Identifier
// unique 한 값을 저장하기 위한 field 로 사용한다
type Identifier map[IdentifierKey]string
type IdentifierKey string

const (
	IdentifierKeyUsername IdentifierKey = "username"
	IdentifierKeyEmail    IdentifierKey = "email"
)

func (x IdentifierKey) String() string {
	return string(x)
}

/*------------------------------------------------------------------------------------------------*/

// Property
// unique 하지 않는 값을 저장하기 위한 field 로 사용된다.
// go -> mongo 로 번역될 때, 기본값으로 변경 되므로
// mongo -> go 로 번역될 때는 기본값으로 변경된다.
// 왠만해서 struct 정의해서 사용하는 편이 좋다.
type Property map[PropertyKey]any
type PropertyKey string
