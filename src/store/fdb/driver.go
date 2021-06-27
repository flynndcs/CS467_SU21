package fdb

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

var (
	productSubspace subspace.Subspace
	db              fdb.Database
)

func initFDB() (db fdb.Database) {
	fdb.MustAPIVersion(630)
	db = fdb.MustOpenDefault()
	productDir, err := directory.CreateOrOpen(db, []string{"product"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	productSubspace = productDir.Sub("product")
	return db
}

func Put(key string, value []byte) (didPut bool) {
	db = initFDB()
	productKey := productSubspace.Pack(tuple.Tuple{key})
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(productKey, value)
		return
	})
	if err != nil {
		log.Fatalf("Unable to set value: (%v)", err)
		return false
	}
	return true
}

func Get(key string) (value []byte) {
	db = initFDB()
	productKey := productSubspace.Pack(tuple.Tuple{key})
	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(productKey).MustGet()
		return
	})
	if err != nil {
		log.Fatalf("Unable to read FDB database value: (%v)", err)
	}
	return ret.([]byte)
}

func GetRange(beginKey string, endKey string) (repeatedValue []byte) {
	db = initFDB()
	beginProductKey := productSubspace.Pack(tuple.Tuple{beginKey})
	endProductKey := productSubspace.Pack(tuple.Tuple{endKey})

	log.Default().Println("begin: ", beginKey)
	log.Default().Println("end: ", endKey)
	selectorRange := fdb.SelectorRange{Begin: fdb.FirstGreaterOrEqual(beginProductKey), End: fdb.FirstGreaterOrEqual(endProductKey)}

	var values []byte
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		retIter := tr.GetRange(selectorRange, fdb.RangeOptions{}).Iterator()
		for retIter.Advance() {
			kv := retIter.MustGet()
			values = append(values, kv.Value...)
		}
		return values, nil
	})
	if err != nil {
		log.Fatalf("Unable to read FDB database value: (%v)", err)
	}
	repeatedValue = values
	return
}

func Clear(key string) (didClear bool) {
	db = initFDB()
	productKey := productSubspace.Pack(tuple.Tuple{key})
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Clear(productKey)
		return
	})
	if err != nil {
		log.Fatalf("Unable to clear FDB database key-value pair for key: (%v)", err)
		return false
	}
	return true
}
