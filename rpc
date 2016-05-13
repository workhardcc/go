package main

import (
	//    "encoding/json"
	"fmt"
	"os"
	//    "io/ioutil"
	"net/http"
	"net/rpc"
	"strconv"
)

type Create struct {
	Cpu  int
	Mem  int
	Disk int
}

type Delete struct {
	Vname string
	Uuid  string
}

func rpc_clinet(cpu int, mem int, disk int) int {
	client, err := rpc.DialHTTP("tcp", "10.10.34.89:5978")
	if err != nil {
		fmt.Println("Dialerror" + err.Error())
	}

	args := &Create{1, 2, 100}
	var reply int
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		fmt.Println("Arith.Add err" + err.Error())
	}
	fmt.Printf("Arith: %d", reply)
	return reply
}
func rpc_client_create(cpu int, mem int, disk int) bool {
	client, err := rpc.DialHTTP("tcp", "10.10.34.89:5978")
	if err != nil {
		fmt.Println("Dialerror" + err.Error())
	}

	args := &Create{1, 2, 100}
	var reply bool
	err = client.Call("Arith.Create", args, &reply)
	if err != nil {
		fmt.Println("Arith.Create err" + err.Error())
	}
	fmt.Printf("Arith: %t", reply) //  %t  is  bool
	return reply
}

func main() {
	http.HandleFunc("/http_rpcclient.go", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Method == "GET" || req.Method == "POST" {
			classic := req.FormValue("classic")
			if classic == "" {
				fmt.Println("[ERROR] Args is missing")
				os.Exit(2)
			}
			switch classic {
			case "create":
				cpu := req.FormValue("cpu")
				mem := req.FormValue("mem")
				disk := req.FormValue("disk")
				println(cpu)
				println(mem)
				println(disk)
				CPU, _ := strconv.Atoi(cpu)
				MEM, _ := strconv.Atoi(mem)
				DISK, _ := strconv.Atoi(disk)
				rpc_client_create(CPU, MEM, DISK)
			case "delete":
				vname := req.FormValue("vname")
				uuid := req.FormValue("uuid")
				println(vname)
				println(uuid)
			default:
				fmt.Println("error")
				os.Exit(2)
			}
		} else {
			http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
		}
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed:", err.Error())
	}
}
