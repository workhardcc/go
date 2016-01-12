package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"net/http"
)

type sub struct {
	Record_type  string
	Record_view  string
	Ttl          int
	Record_name  string
	Record_value string
}

type data struct {
	Ret bool
	Val []sub
}

func main() {
	//	var result []byte
	http.HandleFunc("/memcache_to_http.go", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Method == "GET" || req.Method == "POST" {
			ip := req.FormValue("ip")
			result := memcacheGet(ip)
			//			result = memcacheGet("1.1.1.1")
			//			r, err := json.Marshal(string(result))
			//			if err != nil {
			//				fmt.Println("ERROR: marshal failed!")
			//			} else {
			//				w.Write([]byte(string(r))) //这里必须先string(r),Write方法只接收字符切片
			//				//				fmt.Println(string(r))
			//			}
			res := result.([]byte)
			w.Write([]byte(string(res)))
		} else {
			http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed: ", err)
	}
}

//func memcacheGet(domain string) (res []byte) { //这里必须加上[]string否则报
func memcacheGet(domain string) (res interface{}) { //这里必须加上[]string否则报
	con := memcache.New("127.0.0.1:11211")
	result, err := con.Get(domain)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR: Get data from memcache failed!")
	}
	//	} else {
	res = result.Value
	return res
}
