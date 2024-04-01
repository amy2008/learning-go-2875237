package main

import (
	"fmt"
	"sort"
)

func main() {
	provinces := make(map[string]string)
	fmt.Println(provinces)

	provinces["HB"] = "HeBei"
	provinces["BJ"] = "BeiJing"
	provinces["TJ"] = "TianJin"
	fmt.Println(provinces)

	tianjin := provinces["TJ"]
	fmt.Println(tianjin)

	delete(provinces, "HB")
	provinces["SD"] = "ShanDong"
	fmt.Println(provinces)

	for k, v := range provinces {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(provinces))
	i := 0
	for k := range provinces {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(provinces[keys[i]])
	}

}
