# Pusher
Simple use of publish/subscribe in web, grpc, graphQL

### Contributing
You can commit PR to this repository

### Overview
- This is a sample pub/sub library, but useful
- Irregularly updated

### How to get it?
````
go get -u github.com/gobkc/pusher
````

### Quick start
````
package main

import (
	"fmt"
	"github.com/gobkc/pusher"
	"sync"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	pushUser := User{
		Name:     "test user",
		Password: "123456",
	}
	topic := "test topic"
	wg := sync.WaitGroup{}
	wg.Add(2)
	pusher.Subs(topic, func(cb pusher.SubsCallback) {
		var user User
		cb.Bind(&user)
		fmt.Println(user)
		wg.Done()
	})
	pusher.Subs(topic, func(cb pusher.SubsCallback) {
		var user User
		cb.Bind(&user)
		fmt.Println(user)
		wg.Done()
	})
	pusher.Push(topic, pushUser)
	wg.Wait()
}
````

### License
© Gobkc, 2022~time.Now

Released under the Apache License