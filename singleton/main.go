package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

func (DataBase) CreateConnection() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(time.Second * 2)
	fmt.Println("Done!")
}

var db *DataBase
var lock sync.Mutex

func GetDataBase() *DataBase {
	lock.Lock()
        defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating database...")
		db = &DataBase{}
		db.CreateConnection()
	} else {
		fmt.Println("DB Already Created!")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDataBase()
		}()
	}

	wg.Wait()
}
