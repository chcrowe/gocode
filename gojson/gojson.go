package main

import (
  "fmt"
  "encoding/json"
  // "net/http"
  // "time"
  // "os"
  // "os/signal"
  // "runtime"
  // "path/filepath"
)

//use the json editor at http://www.jsoneditoronline.org/

func unmarshallTest() {

	byt := []byte(`{

  "name" : "Charles Rowe",
  "email" : "chcrowe@gmail.com",
  "altemail" : "chrisrowemd@hotmail.com",
  "hash" : "872UBDEISY8723HE",
  "phones" : [
        {
            "name":"mobile",
            "number":"7034316906",
            "carrier":"AT&T"
        },
        {
            "name":"office",
            "number":"7034219101x233"
        },
        {
            "name":"facsimile",
            "number":"7034219158"
        }
      ],
  "addresses": [
        {
            "name":"billing",
            "street": "13120 Mercury Lane",
            "street2": "",
            "city": "Fairfax",
            "state": "VA",
            "zip": "22033",
            "country": "US"
      },
        {
        	"name":"shipping",
            "street": "13120 Mercury Lane",
            "street2": "",
            "city": "Fairfax",
            "state": "VA",
            "zip": "22033",
            "country": "US"
      },
        {
        	"name":"work",
            "street": "10 Pidgeon Hill Drive",
            "street2": "Suite 210",
            "city": "Sterling",
            "state": "VA",
            "zip": "20165",
            "country": "US"
      },
        {
        	"name":"work 2",
            "street": "45195 Business Court",
            "street2": "Suite 350",
            "city": "Dulles",
            "state": "VA",
            "zip": "20166",
            "country": "US"
      }
  ]
  	}`)

	// var objmap map[string]*json.RawMessage
	// err := json.Unmarshal(data, &objmap)


	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
	fmt.Println(dat)

 	fmt.Printf("\n\n")
 	   
 	fmt.Printf("name: %v\n", dat["name"].(string))
 	fmt.Printf("email: %v\n", dat["email"].(string))
 	fmt.Printf("altemail: %v\n\n", dat["altemail"].(string))

 	phones := dat["phones"].([]interface{})		
	fmt.Printf("phones: %d items\n", len(phones))

	for _, v := range phones {
		v1 := v.(map[string]interface{})
        fmt.Printf("  keys(%d) => %v\n", len(v1), v1)		
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
}


func main() {


    unmarshallTest()
 
}