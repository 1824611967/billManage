package main

import (
    "fmt"
    "net/http"
    "accessory"
    "io/ioutil"
    "bytes"
    "encoding/json"
    "html/template"
    "mime"
    "math"
    "time"
    "strconv"
    "dbslite"
)

var response *accessory.Response

func add(left int, right int) int{
    return left + right
}

var mapInfo = map[string]interface{}{}

//翻页查询函数
func billPageFunc(w http.ResponseWriter, r *http.Request){
    //解析
    if err := r.ParseForm(); err != nil {
        response.Code = accessory.ComErrorId
        response.Msg = accessory.ComErrorMsg
        return
    }
    
    //根据传入的参数执行对应的方法
    
    defer func(){
        if r := recover(); r != nil {
            FindPagefunc(w, 1)
            fmt.Println("sss")
        }
    }()

    page, err := strconv.Atoi(r.Form["page"][0])
    
    
    if err != nil{
        response.Code = accessory.ErrorParameterIsErrId
        response.Msg = accessory.ErrorParameterIsErrMsg
        return
    }

    if page != 0{
        fmt.Println(page)
        FindPagefunc(w, page)
    }
    
}

//查询函数
func searchBill(w http.ResponseWriter, r *http.Request){
    if err := r.ParseForm(); err != nil{
        response.Code = accessory.ComErrorId
        response.Msg = accessory.ComErrorMsg
        return
    }

    keyword := r.Form["keyword"][0]
    page, _ := strconv.Atoi(r.Form["page"][0]) 

    sql1 := "select * from Bill where b_money >= " + keyword
    sql2 := "select * from Bill where b_category like '" + keyword + "'"
    sql3 := "select * from Bill where b_money like %'" + keyword + "'%"
    bills := getSqlBill(sql2)

    if len(bills) ==0{
        money, err1 := strconv.Atoi(keyword)

        if err1 != nil{
            fmt.Println(money)
            money = 0
            bills = getSqlBill(sql3)
        }else{
            bills = getSqlBill(sql1)
        }
        
    }

    if keyword == ""{
        bills = getAllBill()
    }

    funcMap := template.FuncMap{"add":add}
    
	t := template.Must(template.New("homework.html").Funcs(funcMap).ParseFiles("G:/billmanage/src/view/homework.html"))
    
    totalMoney := 0

    for _, v := range bills{
        totalMoney += v.BillMoney
    }
    var billlength int 

    billlength = len(bills)

    if billlength >= 18*page{
        mapInfo["totalBill"] = bills[(page-1)*18:(page)*18]
    }else{
        mapInfo["totalBill"] = bills[(page-1)*18:]
    }
  
    mapInfo["totalMoney"] = totalMoney
	
    mapInfo["billlen"] = int(math.Ceil(float64(len(bills)) / 18))
    mapInfo["presentPage"] = page
	err := t.Execute(w,mapInfo) 
    fmt.Println(err)
	return

}

//根据输入的sql语句返回[]Bill
func getSqlBill(sql string) []accessory.Bill{
    var bill accessory.Bill
	var bills []accessory.Bill
    var parse_time time.Time
    
    result, err := dbslite.Db.Query(sql)
    
	if err != nil{
        fmt.Println(err)
		return bills
	}

	for result.Next(){
        err = result.Scan(&bill.BillId, &bill.BillMoney,&bill.BillCmt, &parse_time, &bill.BillCategory)
        
		bill.BillDate = parse_time.Format("2006-01-02")
		if err != nil{
            fmt.Println(err)
			return bills
		}

		bills = append(bills, bill)
	}

    return bills
}

//根据page查询对应页数的数据
func FindPagefunc(w http.ResponseWriter, page int){
    var bill accessory.Bill
	var bills []accessory.Bill
	var parse_time time.Time
    var i int
	sql := "select * from (select Bill.*,rownum rc from Bill where rownum <= " + strconv.Itoa((page) * 18) + ") Bill where Bill.rc >= " + strconv.Itoa((page-1) * 18 + 1)

	result, err := dbslite.Db.Query(sql)
    
	if err != nil{
        fmt.Println(err)
		return 
	}

	for result.Next(){
        err = result.Scan(&bill.BillId, &bill.BillMoney,&bill.BillCmt, &parse_time, &bill.BillCategory, &i)
        
		bill.BillDate = parse_time.Format("2006-01-02")
		if err != nil{
            fmt.Println(err)
			return 
		}

		bills = append(bills, bill)
	}

	funcMap := template.FuncMap{"add":add}
    
	t := template.Must(template.New("homework.html").Funcs(funcMap).ParseFiles("G:/billmanage/src/view/homework.html"))
    
    totalMoney := billGetMoney()
    totalBill := getAllBill()
    mapInfo["totalMoney"] = totalMoney
	mapInfo["totalBill"] = bills
    mapInfo["billlen"] = int(math.Ceil(float64(len(totalBill)) / 18))
    mapInfo["presentPage"] = page
	err = t.Execute(w,mapInfo) 

	return
}

//添加账单对应函数
func billAddFunc(w http.ResponseWriter, r *http.Request){
    response = accessory.NewResponse()

    setHeader(w)

    defer response.Answer(w)

    if r.Method == "POST" {
        var request accessory.BillRequest
        

        bodyStr := GetData(r)

        if err := json.Unmarshal([]byte(bodyStr), &request); err != nil{
            response.Code = accessory.ErrorJsonErrId
            response.Msg  = accessory.ErrorJsonErrMsg
        }
        fmt.Println(request.Data)
        switch request.Com {
        case "add":
            billAddMethod(w, request.Data)
        case "update":
            billUpdateMethod(w, request.Data)
        case "delete":
            billDeleteMethod(w, request.Data)
        }
    }

}

//获取request中信息函数
func GetData(res *http.Request) string{
    result, err := ioutil.ReadAll(res.Body)

    if err != nil{
        return `{"code":"10000", "msg":"failed"}`
    } else {
        return bytes.NewBuffer(result).String()
    }
}

func setHeader(w http.ResponseWriter) {
	//一种解决资源跨域资源的策略
	w.Header().Add("Access-Control-Allow-Origin", "*")
	//用于 preflight request （预检请求）中，列出了将会在正式请求的 Access-Control-Expose-Headers 字段中出现的首部信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	//是否允许发送cookie
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	//表明服务器支持的所有的跨域请求方法
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

//程序的进口
func main()  {
    mux := http.NewServeMux()

    mime.AddExtensionType(".css", "text/css")
    

    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("G:/billmanage/src/static"))))
   
    mux.Handle("/index/", http.HandlerFunc(billPageFunc))
    mux.Handle("/bill/manage", http.HandlerFunc(billAddFunc))
    mux.Handle("/bill/search", http.HandlerFunc(searchBill))

    http.ListenAndServe(":3000", mux)
}