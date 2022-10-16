package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool digunakan untuk menyimpan data(mirip slice atau array) untuk mengatasi problem race condition

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Romzi")
	pool.Put("Farhan")
	pool.Put("Ghozi")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
