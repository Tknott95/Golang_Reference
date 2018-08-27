package main

// 8/27/2018 - Lightweight UFC rankings - TOP 7
// REF: ("http://m.ufc.com/rankings/lightweight")

import (
	"fmt"
)

func main() {
	fighters := make(map[string]float32)

	fighters["Khabib Nurmagomedov"] = 0
	fighters["Conor McGregor"] = 1
	fighters["Justin Gaethje"] = 7
	fighters["Dustin Porier"] = 3
	fighters["Tony Ferguson"] = 2
	fighters["Edson Barboza"] = 6
	fighters["Eddie Alvarez"] = 4
	fighters["Kevin Lee"] = 5

	fmt.Println(fighters)

	delete(fighters, "Conor McGregor")

	fmt.Printf("\n After Deletion: %v \n", fighters)

	for _key, _val := range fighters {
		fmt.Println(_key, ":", _val)	
	}
	

}
