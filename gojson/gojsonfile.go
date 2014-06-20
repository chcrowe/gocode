package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  // "net/http"
  // "time"
   "os"
  // "os/signal"
  // "runtime"
  // "path/filepath"
)

//use the json editor at http://www.jsoneditoronline.org/

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func unmarshallTest(f string) {

	//byt, err := ioutil.ReadFile("./gojsonfile.json")
	byt, err := ioutil.ReadFile(f)
	check(err)

	// var objmap map[string]*json.RawMessage
	// err := json.Unmarshal(data, &objmap)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
	//fmt.Println(dat)


 	fmt.Printf("\n\n")
 	   
 	fmt.Printf("name: %v\n", dat["name"].(string))
 	fmt.Printf("email: %v\n", dat["email"].(string))
 	fmt.Printf("altemail: %v\n\n", dat["altemail"].(string))


 	phones := dat["phones"].([]interface{})		
	fmt.Printf("phones: %d items\n", len(phones))

	for _, v := range phones {
		v1 := v.(map[string]interface{})
		
		carrier, ok := m["carrier"]
		if !ok{ carrier = "(none)"}

        fmt.Printf("  keys(%d) => %v, carrier: %v\n", len(v1), v1, carrier)		
	}

	fmt.Printf("\n\n")
 
 	addresses := dat["addresses"].([]interface{})		
	fmt.Printf("addresses: %d items\n", len(addresses))

	for _, v := range addresses {
		v1 := v.(map[string]interface{})
        fmt.Printf("  %-12v => %v %v\n", v1["name"], v1["street"], v1["street2"])		
        fmt.Printf("  %-12v    %v %v %v (%v)\n", "", v1["city"], v1["state"], v1["zip"], v1["country"])		
	}

	fmt.Printf("\n\n")


	// email := dat["email"].(string)
	// altemail := dat["altemail"].(string)

	// dat["email"] = altemail
	// dat["altemail"] = email


	// d1, _ := json.MarshalIndent(dat, "", "    ")
	// err = ioutil.WriteFile(f, d1, 0644)
	// check(err)
}


func main() {


    unmarshallTest(os.Args[1])
 
}