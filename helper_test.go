package pusher

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	type args struct {
		topic string
		data  any
	}
	type User struct {
		Name string
		Pass string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				topic: "test topic",
				data: User{
					Name: "test user",
					Pass: "123456",
				},
			},
		},
	}
	for _, tt := range tests {
		wg := sync.WaitGroup{}
		wg.Add(2)
		t.Run(tt.name, func(t *testing.T) {
			Push(tt.args.topic, tt.args.data)
			//there is no subscriber,so you need to sleep for 1 second
			time.Sleep(1 * time.Second)
			Subs(tt.args.topic, func(cb SubsCallback) {
				var user User
				cb.Bind(&user)
				fmt.Println(user)
				wg.Done()
			})
			Subs(tt.args.topic, func(cb SubsCallback) {
				var user User
				cb.Bind(&user)
				fmt.Println(user)
				wg.Done()
			})
		})
		wg.Wait()
	}
}
