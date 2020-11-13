package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDatabase(t *testing.T) {
	dname, e := ioutil.TempDir("", "sampledir")
	if e != nil {
		t.Fatal(e)
	}
	defer os.RemoveAll(dname)

	db, err := badger.Open(badger.DefaultOptions(dname))
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	var val []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			t.Fatal(err)
		}
		err = item.Value(func(v []byte) error {
			val = append(val, v...)
			return nil
		})
		return nil
	})
	fmt.Printf("Answer: %s\n", val)
}

func TestDatabase2(t *testing.T) {
	db := new(DB)
	db.Init(0)
	db.Put("test", []byte("answer"), []byte("42"))
	answer := db.Get("test", []byte("answer"))
	fmt.Println("The Answer is ", answer)
}
