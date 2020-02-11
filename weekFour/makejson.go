package main
import "fmt"
import "encoding/json"

func makejson() {
	var name string
	//user enters name
	fmt.Printf("Enter a name: ")
	fmt.Scan(&name)

	var address string
	//user enters address
	fmt.Printf("Enter an address: ")
	fmt.Scan(&address)

	//create map
	userMap := make(map[string]string)
	
	//add name, address as keys and values
	userMap["name"] = name
	userMap["address"] = address

	//Marshal() to create JSON object from the map
	mapBytes, _ := json.Marshal(userMap)
	fmt.Println("Marshal to create JSON object from the map:", mapBytes)

	//print JSON object
	fmt.Println("JSON object:", string(mapBytes))
}