package hello_go

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"server_name"`
	ServerIp   string `json:"server_ip"`
	ServerPort int    `json:"server_port"`
}

// Json 结构体: 序列化
func SerializeStruct() {
	myServer := new(Server)
	myServer.ServerName = "Demo of json(struct) serialize:"
	myServer.ServerIp = "127.0.0.1"
	myServer.ServerPort = 8080

	b, err := json.Marshal(myServer)
	if nil != err {
		fmt.Println("Marshal err:", err.Error())
		return
	}

	fmt.Println("Marshal json(struct):", string(b))
	fmt.Println("")
}

// Json Map: 序列化
func SerializeMap() {
	myServer := make(map[string]interface{})
	myServer["ServerName"] = "Demo of json(map) serialize:"
	myServer["ServerIp"] = "198.0.0.1"
	myServer["ServerPort"] = 9090

	b, err := json.Marshal(myServer)
	if nil != err {
		fmt.Println("Marshal err:", err.Error())
		return
	}

	fmt.Println("Marshar json(map):", string(b))
	fmt.Println("")
}

// Json 结构体: 反序列化
func UnSerializeStruct() {
	jsonData := `{"server_name":"Demo of json(struct) serialize:","server_ip":"127.0.0.1","server_port":8080}`
	myServer := new(Server)

	err := json.Unmarshal([]byte(jsonData), &myServer)
	if nil != err {
		fmt.Println("Unmarshal err:", err.Error())
		return
	}

	fmt.Println("Unmarshal json(struct):", myServer)
	fmt.Println("")
}

// Json Map: 反序列化
func UnSerializeMap() {
	var jsonData string
	// jsonData = `{"ServerIp":"198.0.0.1","ServerName":"Demo of json(map) serialize:","ServerPort":9090}`
	// return `map[ServerIp:198.0.0.1 ServerName:Demo of json(map) serialize: ServerPort:9090]`
	jsonData = `{"server_name":"Demo of json(map) serialize:","server_ip":"198.0.0.1","server_port":9090}`
	// return `map[server_ip:198.0.0.1 server_name:Demo of json(map) serialize: server_port:9090]`
	myServer := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonData), &myServer)
	if nil != err {
		fmt.Println("Unmarshal err:", err.Error())
		return
	}

	fmt.Println("Unmarshal json(map):", myServer)
	fmt.Println("")
}
