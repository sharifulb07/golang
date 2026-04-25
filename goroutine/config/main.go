package main

import (
	"fmt"
	"sync"
	"time"
)

// ----config------

type Config struct{
	AppName string 
	MaxUsers int 
	EnableSignUp bool 
	RateLimit int 
}


// ---------configStore-------

type configStore struct{
	mu sync.RWMutex
	config Config
}



// ---------create newConfigStore----------


func NewConfigStore()*configStore{
	return &configStore{
		config: Config{
		AppName: "My Saas App",
		MaxUsers: 1000,
		EnableSignUp: true,
		RateLimit: 100,
		},
	
	}
}



// ---READ----


func (c *configStore)GetConfig()Config{
	c.mu.Lock()


	defer c.mu.Unlock()

	return c.config
	
}


// ---------WRITE --------------

func (c *configStore)UpdateStore(newConfig Config){
	c.mu.Lock()

	defer c.mu.Unlock()

	c.config=newConfig

}


// -------partial update example---------


func (c *configStore) UpdateRateLimit(limit int){
	c.mu.Lock()

	defer c.mu.Unlock()

	c.config.RateLimit=limit
}


// ----simulation------


func reader(id int, store *configStore, wg *sync.WaitGroup){

	defer wg.Done()


	for i:=0;i<3;i++{

		cfg:=store.GetConfig()

		fmt.Printf(
			"Reader %d | App -> %s | rate -> %d \n",
			id,
			cfg.AppName,
			cfg.RateLimit,
		)
		time.Sleep(500*time.Millisecond)
	}
}


// writer ----


func writer(store *configStore, wg *sync.WaitGroup){
	defer wg.Done()

	time.Sleep(2*time.Second)

	fmt.Printf("\n Writer updating config ....")
	store.UpdateRateLimit(500)

}


func main(){

	store:=NewConfigStore()

	var wg sync.WaitGroup


	for i:=0; i<5; i++{
		wg.Add(1)

		go reader(i, store, &wg)
	}

	wg.Add(1)
	go writer(store, &wg)


	wg.Wait()


	fmt.Println("\n final Config:  ")

	fmt.Println(store.GetConfig())

}