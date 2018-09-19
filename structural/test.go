package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Rankings struct {
    Keyword  string `json:"keyword"`
    GetCount uint32 `json:"get_count"`
    Engine   string `json:"engine"`
    Locale   string `json:"locale"`
    Mobile   bool   `json:"mobile"`
}

func main() {
    var jsonBlob = []byte(`
        {"keyword":"hipaa compliance form", "get_count":157, "engine":"google", "locale":"en-us", "mobile":false}
    `)
    rankings := Rankings{}
    err := json.Unmarshal(jsonBlob, &rankings)
    if err != nil {
        // nozzle.printError("opening config file", err.Error())
    }

    rankingsJson, _ := json.Marshal(rankings)
    err = ioutil.WriteFile("output.json", rankingsJson, 0644)
    fmt.Printf("%+v", rankings)
    fmt.Println()
}