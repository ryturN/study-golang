package main

import (
	"fmt"
	"time"
)

func main() {
	person := map[string]string{
		"name":  "ryan",
		"email": "m.adryan@gmail.com",
		"phone": "08123456789",
		"age":   "20",
	}
	newPerson := make(map[string]string)
	newPerson["name"] = "joko"
	newPerson["email"] = "joko@gmail.com"

	for v, k := range newPerson {
		person[v] = k
		if k == "joko" {
			fmt.Println(len(person))
			fmt.Println(`<b>welcome to the person<b>?`, person[v])
		} else {
			fmt.Println(v, "\t:", k)
		}

	}
	waktu := time.Now()
	fmt.Println(waktu.Date())
	fmt.Println(waktu.Hour())
}
