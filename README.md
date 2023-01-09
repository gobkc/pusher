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
	type User struct {
		Name string
		Pass string
	}
    wg := sync.WaitGroup{}
    wg.Add(2)
    topic := "test topic"
    Push(topic, tt.args.data)
    //there is no subscriber,so you need to sleep for 1 second
    time.Sleep(1 * time.Second)
    Subs(topic, func(cb SubsCallback) {
        var user User
        cb.Bind(&user)
        fmt.Println(user)
        wg.Done()
    })
    Subs(topic, func(cb SubsCallback) {
        var user User
        cb.Bind(&user)
        fmt.Println(user)
        wg.Done()
    })
    wg.Wait()
````

### License
© Gobkc, 2022~time.Now

Released under the Apache License