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

func InitFDB() {
	log.Default().Println("initializing FDB")
	fdb.MustAPIVersion(630)
	db = fdb.MustOpenDefault()
	productDir, err := directory.CreateOrOpen(db, []string{"product"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	productSubspace = productDir.Sub("product")
}

func Put(key string, value []byte) (didPut bool) {
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
	beginProductKey := productSubspace.Pack(tuple.Tuple{beginKey})
	endKeyInclusive, errStrinc := fdb.Strinc([]byte(endKey))
	if errStrinc != nil {
		log.Fatal("Could not get real end key from endKey")
	}
	endProductKey := productSubspace.Pack(tuple.Tuple{string(endKeyInclusive)})

	selectorRange := fdb.SelectorRange{Begin: fdb.FirstGreaterOrEqual(beginProductKey), End: fdb.FirstGreaterOrEqual(endProductKey)}

	var values []byte
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		retIter := tr.GetRange(selectorRange, fdb.RangeOptions{}).GetSliceOrPanic()
		for _, kv := range retIter {
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
