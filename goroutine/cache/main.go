package main

import (
	"fmt"
	"sync"
	"time"
)

// cache item

type Item struct {
	Value      interface{}
	Expiration int64
}

func (item Item) isExpired() bool {
	if item.Expiration == 0 {
		return false
	}

	return time.Now().UnixNano() > item.Expiration
}

// cache

type Cache struct {
	data  map[string]Item
	mutex sync.Mutex
}


// create a new cache

func NewCache() *Cache {
	c := &Cache{
		data: make(map[string]Item),
	}

	// background cleanup
	go c.startGC()

	return c
}




// -----------SET---------
func(c *Cache)Set(key string, value interface{}, ttl time.Duration){
	var exp int64

	if ttl>0{
		exp=time.Now().Add(ttl).UnixNano()
	}

	c.mutex.Lock()

c.data[key]=Item{
	Value: value,
	Expiration: exp,
}

	c.mutex.Unlock()
}





// ----------------GET---------

	func (c *Cache)Get(key string)(interface{}, bool){

		c.mutex.Lock()

		item, found:=c.data[key]

		if !found{
			c.mutex.Unlock()
			return nil, false
		}

		if item.isExpired(){
			delete(c.data, key)
			c.mutex.Unlock()
			return nil, false
		}



		c.mutex.Unlock()
		return item.Value, true
	}

// ----------DELETE----------
func (c *Cache)Delete(key string)  {
	
	c.mutex.Lock()
	delete(c.data, key)
	c.mutex.Unlock()
}




// -----cleanup GC-----
func (c *Cache)startGC(){
	ticker:=time.NewTicker(2*time.Second)

	for range ticker.C{
		c.mutex.Lock()

		for key, item:=range c.data{
			if item.isExpired(){
				delete(c.data,key)
			}
		}

		c.mutex.Unlock()
	}

}



func main() {


	cache:=NewCache()

	cache.Set("user:1", "shariful", 5*time.Millisecond)
	cache.Set("user:2", "Adiyat", 0)

	// concurrent readers () race test)

	var wg sync.WaitGroup

	for i:=0; i<10;i++{
		wg.Add(1)


		go func (id int){

			defer wg.Done()

			val, ok:=cache.Get("user:1")

			if ok{
				fmt.Printf("Goroutine %d -> %v \n", id, val)
			}else{
				fmt.Printf("Goroutine %d expired/not found\n", id)
			}
		}(i)
	}


	wg.Wait()

	value, okay:=cache.Get("user:1")
	
	fmt.Println("\n After Expiration")
	fmt.Println("Found", okay, "Value: ", value)


}
