package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	eventCnt := 0
	for eventCnt < 3 {
		fmt.Println("Retrieving events ...")
		eventCnt++
		if eventCnt == 3 {
			fmt.Printf("Got %v notif, updating Phone notifi\n", eventCnt)
		}
	}

	names := []string{"Isma", "You", "Zey", "Ami"}
	for i, n := range names {
		fmt.Printf("Username=%s (index=%d)\n", n, i)
	}

	for _, c := range "golang" {
		fmt.Printf("%v\n", string(c))
	}

}
