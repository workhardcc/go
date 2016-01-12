package main

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"io/ioutil"
	"net/http"
	"strconv"
	//	"reflect"
)

type sub struct {
	Id           string
	Sid          string
	Zone_id      string
	Record_name  string
	Record_view  string
	Record_type  string
	Record_value string
}

type data struct {
	Ret bool
	Val []sub
}

type mid []string
type result map[string]mid

func main() {
	domain_list := make(map[string]int)
	domain_list["xx.cn"] = 37
	domain_list["yy.cn"] = 38

	furl := "http://testcc.php?env=online&type=private&act=getDnsRecords&zone_id="
	r := make(result, 0) //这是个map,不管map里的val是什么
	for domain, num := range domain_list {
		url := furl + strconv.Itoa(num)
		fmt.Println(url)
		res := httpGet(url)
		for _, v := range res {
			full_record_name := v.Record_name + "." + domain
			r[v.Record_value] = append(r[v.Record_value], full_record_name) //如果不存在也append,存在就更append了
		}
	}
	//	final, err := json.Marshal(r)	   json_encode后展示出来
	//	if err != nil {
	//		fmt.Println("ERROR:", err)
	//	}
	//	fmt.Println(string(final))
	mc := memcache.New("127.0.0.1:11211")
	for key, val := range r {
		d, err := json.Marshal(val) //由于memcache.set方法只能接受[]byte类型，所以必须是字符串
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		mc.Set(&memcache.Item{Key: key, Value: d})
	}
}

//func httpGet() (tmp []sub) {
func httpGet(url string) (tmp []sub) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Get http failed")
	}

	defer res.Body.Close()

	var r data
	body, err := ioutil.ReadAll(res.Body)
	jsonerr := json.Unmarshal(body, &r) //json.Unmarshal()函数将一个JSON对象解码到空接口r中，最终r将会是一个键值对的 map[string]interface{} 结构
	if jsonerr != nil {
		fmt.Println("ERROR: Parse json faile!", jsonerr)
	}
	tmp = r.Val
	//	fmt.Println(tmp)
	//	fmt.Println("type:", reflect.TypeOf(tmp))
	return tmp
}
