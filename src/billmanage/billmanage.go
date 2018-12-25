package main

import (
    "fmt"
    "net/http"
    "accessory"
    "io/ioutil"
    "bytes"
    "encoding/json"
)

var response *accessory.Response

func billIndex(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("账单初始化"))
    fmt.Print("sd")
}

func billAddFunc(w http.ResponseWriter, r *http.Request){
    response = accessory.NewResponse()

    defer response.Answer(w)

    if r.Method == "POST" {
        var request accessory.BillRequest

        bodyStr := GetData(r)

        if err := json.Unmarshal([]byte(bodyStr), &request); err != nil{
            response.Code = accessory.ErrorJsonErrId
            response.Msg  = accessory.ErrorJsonErrMsg
        }

        switch request.Com {
        case "add":
            billAddMethod(w, request.Data)
        }

    }
}
func GetData(res *http.Request) string{
    result, err := ioutil.ReadAll(res.Body)
    if err != nil{
        return `{"code":"10000", "msg":"failed"}`
    } else {
        return bytes.NewBuffer(result).String()
    }
}
func main()  {
    mux := http.NewServeMux()
    
    mux.Handle("/index", http.HandlerFunc(billIndex))
    mux.Handle("/add", http.HandlerFunc(billAddFunc))

    http.ListenAndServe(":3000", mux)
}