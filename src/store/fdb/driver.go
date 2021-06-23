package fdb

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func Put(key string, value []byte) (didPut bool) {
	fdb.MustAPIVersion(630)
	db := fdb.MustOpenDefault()
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(fdb.Key(key), value)
		return
	})
	if err != nil {
		log.Fatalf("Unable to set value: (%v)", err)
		return false
	}
	return true
}

func Get(key string) (value []byte) {
	fdb.MustAPIVersion(630)
	db := fdb.MustOpenDefault()
	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(fdb.Key(key)).MustGet()
		return
	})
	if err != nil {
		log.Fatalf("Unable to read FDB database value: (%v)", err)
	}
	return ret.([]byte)
}

func Clear(key string) (didClear bool) {
	fdb.MustAPIVersion(630)
	db := fdb.MustOpenDefault()
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Clear(fdb.Key(key))
		return
	})
	if err != nil {
		log.Fatalf("Unable to clear FDB database key-value pair for key: (%v)", err)
		return false
	}
	return true
}
