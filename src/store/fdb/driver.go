package fdb

import (
	"bytes"
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"golang.org/x/crypto/bcrypt"
)

var (
	accountSubspace subspace.Subspace
	ProductSubspace subspace.Subspace
	db              fdb.Database
	buffer          bytes.Buffer
)

func EncodeCategories(givenRange []string) (returnBytes []byte) {
	var bytes []byte
	for _, v := range givenRange {
		bytes = append(bytes, []byte(v)...)
	}
	return bytes
}

func InitFDB() {
	log.Println("initializing FDB")
	fdb.MustAPIVersion(630)
	db = fdb.MustOpenDefault()
	productDir, err := directory.CreateOrOpen(db, []string{"product"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	ProductSubspace = productDir.Sub("product")

	accountDir, err := directory.CreateOrOpen(db, []string{"account"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	accountSubspace = accountDir.Sub("account")
}

func CreateAccount(accountName string, userArray []byte) (didCreate bool) {
	var keyBytes []byte
	keyBytes = append(keyBytes, accountSubspace.Bytes()...)
	keyBytes = append(keyBytes, []byte(accountName)...)

	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(fdb.Key(keyBytes), userArray)
		return
	})
	if err != nil {
		log.Println("Could not create account: ", err)
		return false
	}
	return true
}

func CreateUser(accountName string, username string, password string) (didCreate bool) {
	hash, hashError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashError != nil {
		log.Println("could not generate hash: ", hashError)
	}
	var keyBytes []byte
	keyBytes = append(keyBytes, accountSubspace.Bytes()...)
	keyBytes = append(keyBytes, []byte(accountName)...)
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		userArray := tr.Get(fdb.Key(keyBytes)).MustGet()
		userArray = append(userArray, []byte(username)...)

		tr.Set(fdb.Key(keyBytes), userArray)

		keyBytes = append(keyBytes, []byte(username)...)
		tr.Set(fdb.Key(keyBytes), hash)
		return
	})
	if err != nil {
		log.Fatalf("Could not create user (%v)", username)
	}
	return true
}

func CheckCredentials(accountName string, username string, password string) (isValid bool) {
	var keyBytes []byte
	keyBytes = append(keyBytes, accountSubspace.Bytes()...)
	keyBytes = append(keyBytes, []byte(accountName)...)
	keyBytes = append(keyBytes, []byte(username)...)

	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(fdb.Key(keyBytes)).MustGet()
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

func Put(key []byte, value []byte) (didPut bool) {
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(fdb.Key(key), value)
		return
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to set value: (%v)", err)
	}
	buffer.Reset()
	return true
}

func Get(key []byte) (value []byte) {
	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(fdb.Key(key)).MustGet()
		return
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to read FDB database value: (%v)", err)
	}
	buffer.Reset()
	return ret.([]byte)
}

func GetRange(key []byte) (repeatedValue [][]byte) {
	prefixRange, prefixError := fdb.PrefixRange(key)

	if prefixError != nil {
		log.Println("Could not get prefix for key: ", prefixError)
	}

	var values [][]byte

	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		retIter := tr.GetRange(prefixRange, fdb.RangeOptions{}).GetSliceOrPanic()
		for _, kv := range retIter {
			values = append(values, kv.Value)
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

func Clear(key []byte) (didClear bool) {
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Clear(fdb.Key(key))
		return
	})
	if err != nil {
		buffer.Reset()
		log.Fatalf("Unable to clear FDB database key-value pair for key: (%v)", err)
	}
	buffer.Reset()
	return true
}
