package main

import "fmt"

var cache Cache

func main() {
	var err error
	cache, err = NewSqlCacheService()
	if err != nil {
		panic(err.Error())
	}

	err = cache.Set("lahmacun", "acılı")
	if err != nil {
		fmt.Println("Set => " + err.Error())
	}

	val, err := cache.Get("lahmacun")
	if err != nil {
		fmt.Println("Get => " + err.Error())
	}
	fmt.Println(val)
}
