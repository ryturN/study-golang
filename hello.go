package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Decrement(4)
	fmt.Println(message)

	name := "ryan"

	var e = name[0]
	// var eString = string(e)
	fmt.Println(e)
	rows := 5

	for i := 0; i < rows; i++ {
		// Membuat spasi sebelum bintang
		for j := 0; j < rows-i-1; j++ {
			fmt.Print(" ")
		}

		// Membuat bintang
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}

		// Pindah ke baris baru setelah setiap baris selesai
		fmt.Println()
	}

	var names = [...]string{
		"ryan",
		"ganteng",
		"keren",
		"Ziyad",
		"Jeje",
		"jhon doe",
		"jhhonny",
		"jlobbe",
		"nikky",
		"willy",
	}

	var keren = map[string]int{
		"ryan":     14,
		"Ziyad":    23,
		"Jeje":     12,
		"jhon doe": 12,
		"keren":    13,
		"ganteng":  14,
		"jlobbe":   12,
		"jhhonny":  1,
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for key, val := range names {
		fmt.Println(key, val)
	}
	for key, val := range keren {
		fmt.Println(key, "	\t", val)
	}
	fmt.Println(names[0])
	slice := names[:]
	slice2 := append(slice, "keren2222", "ganteng2222")
	fmt.Println(slice2)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println(slice[:len(slice)-9])
	newSlice := make([]string, 3, 5)
	newSlice[0] = "hehe"
	newSlice[1] = "hehe"
	newSlice[2] = "he"
	for i := 0; i < len(newSlice); i++ {
		fmt.Println(newSlice[i])
	}
	fmt.Println(newSlice[1])

	copySlice := make([]string, 3, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
