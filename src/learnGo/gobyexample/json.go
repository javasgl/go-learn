package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	p := fmt.Println

	bolB, _ := json.Marshal(true)
	p(string(bolB))

	intB, _ := json.Marshal(2)
	p(string(intB))

	floatB, _ := json.Marshal(2.34)
	p(string(floatB))

	strB, _ := json.Marshal("gopher")
	p(string(strB))

	sliceB, _ := json.Marshal([]string{"apple", "huawei", "xiaomi"})
	p(string(sliceB))

	mapB, _ := json.Marshal(map[string]int{"apple": 1, "xiaomi": 3, "huawei": 2})
	p(string(mapB))
	mapB, _ = json.Marshal(map[int]string{1: "apple", 3: "xiaomi", 2: "huawei"})
	p(string(mapB))

	byt := []byte(`{"total":12.2,"list":["a","c","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p(dat)
	total := dat["total"].(float64)
	p(total)
	strs := dat["list"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
