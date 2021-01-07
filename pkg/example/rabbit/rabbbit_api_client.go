package rabbit

import (
	"encoding/json"
	"fmt"
	"github.com/michaelklishin/rabbit-hole/v2"
)

func channels() {
	client, _ := rabbithole.NewClient("http://192.168.0.201:8080", "guest", "guest")
	listChannels, _ := client.ListChannels()

	for _, cInfo := range listChannels {
		marshal, _ := json.Marshal(cInfo)
		fmt.Println(string(marshal))
	}
}

func overview() {
	client, _ := rabbithole.NewClient("http://192.168.0.201:8080", "guest", "guest")
	ov, _ := client.Overview()
	marshal, _ := json.Marshal(ov)
	fmt.Println(string(marshal))
}
