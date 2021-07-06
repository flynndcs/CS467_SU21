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
	accountSubspace subspace.Subspace
	productSubspace subspace.Subspace
	db              fdb.Database
	buffer          bytes.Buffer
	enc             *gob.Encoder
)

func encodeKey(exactcategorySequence []string) (returnBuffer bytes.Buffer) {
	enc = gob.NewEncoder(&buffer)
	enc.Encode(exactcategorySequence)
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

	accountDir, err := directory.CreateOrOpen(db, []string{"account"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	accountSubspace = accountDir.Sub("account")
}

func CreateAccount(accountName string, userArray []byte) (didCreate bool) {
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		tr.Set(accountSubspace.Pack(tuple.Tuple{accountName}), userArray)
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
	_, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		userArray := tr.Get(accountSubspace.Pack(tuple.Tuple{accountName})).MustGet()
		userArray = append(userArray, []byte(username)...)

		tr.Set(accountSubspace.Pack(tuple.Tuple{accountName}), userArray)
		tr.Set(accountSubspace.Pack(tuple.Tuple{accountName, username}), hash)
		return
	})
	if err != nil {
		log.Fatalf("Could not create user (%v)", username)
	}
	return true
}

func CheckCredentials(accountName string, username string, password string) (isValid bool) {
	ret, err := db.Transact(func(tr fdb.Transaction) (ret interface{}, e error) {
		ret = tr.Get(accountSubspace.Pack(tuple.Tuple{accountName, username})).MustGet()
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

func Put(name string, categorySequence []string, value []byte) (didPut bool) {
	buffer = encodeKey(append(categorySequence, name))
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

func GetSingle(name string, categorySequence []string) (value []byte) {
	buffer = encodeKey(append(categorySequence, name))
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

func GetAllForCategorySequence(categorySequence []string) (repeatedValue []byte) {
	buffer = encodeKey(categorySequence)
	endKeyInclusive, errStrinc := fdb.Strinc([]byte(categorySequence[len(categorySequence)-1]))
	if errStrinc != nil {
		buffer.Reset()
		log.Fatal("Could not get real end key from categorySequence", errStrinc)
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

func ClearSingle(name string, categorySequence []string) (didClear bool) {
	buffer = encodeKey(append(categorySequence, name))
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
