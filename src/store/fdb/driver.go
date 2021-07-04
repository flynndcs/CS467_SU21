package fdb

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"golang.org/x/crypto/bcrypt"
)

var (
	userSubspace    subspace.Subspace
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
	log.Println("initializing FDB")
	fdb.MustAPIVersion(630)
	db = fdb.MustOpenDefault()
	productDir, err := directory.CreateOrOpen(db, []string{"product"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	productSubspace = productDir.Sub("product")

	userDir, err := directory.CreateOrOpen(db, []string{"user"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	userSubspace = userDir.Sub("user")
}

func CreateUser(username string, password string) (didCreate bool) {
	hash, hashError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashError != nil {
		log.Println("could not generate hash: ", hashError)
	}
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(userSubspace.Pack(tuple.Tuple{username, password}), hash)
		return
	})
	if err != nil {
		log.Fatalf("Could not create user (%v)", username)
	}
	return true
}

func CheckCredentials(username string, password string) (isValid bool) {
	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(userSubspace.Pack(tuple.Tuple{username, password})).MustGet()
		log.Println(ret)
		if bcrypt.CompareHashAndPassword(ret.([]byte), []byte(password)) != nil {
			log.Println("Password did not match")
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		log.Fatalf("Could not authenticate user (%v)", username)
	}
	return ret.(bool)
}

func Put(name string, scope []string, value []byte) (didPut bool) {
	buffer = encodeKey(append(scope, name))
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

func GetSingle(name string, scope []string) (value []byte) {
	buffer = encodeKey(append(scope, name))
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
		End:   fdb.FirstGreaterOrEqual(productSubspace.Pack(tuple.Tuple{endKeyInclusive}))}

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

func ClearSingle(name string, scope []string) (didClear bool) {
	buffer = encodeKey(append(scope, name))
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
