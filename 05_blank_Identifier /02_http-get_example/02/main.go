
package main

import (
	"net/http"	
	"io/ioutil"
	"fmt"
)

func main()  {
		res, _ := http.Get("https://www.mcleodracing.com/")
		page, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Printf("%s",page)
	}
	