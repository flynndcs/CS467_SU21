package fdb

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

var (
	productSubspace subspace.Subspace
	db              fdb.Database
	buffer          bytes.Buffer
	enc             *gob.Encoder
)

func encodeKey(exactScope []string) (returnBuffer bytes.Buffer) {
	enc = gob.NewEncoder(&buffer)
	enc.Encode(exactScope)
	return buffer
}

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

func Put(exactScope []string, value []byte) (didPut bool) {
	buffer = encodeKey(exactScope)
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(productSubspace.Pack(tuple.Tuple{buffer.Bytes()}), value)
		return
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to set value: (%v)", err)
	}
	buffer.Reset()
	return true
}

func GetSingle(exactScope []string) (value []byte) {
	buffer = encodeKey(exactScope)
	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(productSubspace.Pack(tuple.Tuple{buffer.Bytes()})).MustGet()
		return
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to read FDB database value: (%v)", err)
	}
	buffer.Reset()
	return ret.([]byte)
}

func GetAllForScope(scope []string) (repeatedValue []byte) {
	buffer = encodeKey(scope)
	endKeyInclusive, errStrinc := fdb.Strinc([]byte(scope[len(scope)-1]))
	if errStrinc != nil {
		buffer.Reset()
		log.Fatal("Could not get real end key from scope", errStrinc)
	}

	selectorRange := fdb.SelectorRange{
		Begin: fdb.FirstGreaterOrEqual(productSubspace.Pack(tuple.Tuple{buffer.Bytes()})),
		End:   fdb.FirstGreaterOrEqual(productSubspace.Pack(tuple.Tuple{string(endKeyInclusive)}))}

	var values []byte

	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		retIter := tr.GetRange(selectorRange, fdb.RangeOptions{}).GetSliceOrPanic()
		for _, kv := range retIter {
			values = append(values, kv.Value...)
		}
		return values, nil
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to read FDB database value: (%v)", err)
	}
	buffer.Reset()
	return values
}

func ClearSingle(exactScope []string) (didClear bool) {
	buffer = encodeKey(exactScope)
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Clear(productSubspace.Pack(tuple.Tuple{buffer.Bytes()}))
		return
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to clear FDB database key-value pair for key: (%v)", err)
	}
	buffer.Reset()
	return true
}
