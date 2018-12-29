package main

import (
	"dbslite"
	"net/http"
	"accessory"
	"time"
	"fmt"
	"log"
	"strconv"
	_ "encoding/json"
)

//账单数据库插入函数
func billAddMethod(w http.ResponseWriter, data accessory.Bill){
	data.BillId = getId()
	//校验数据
	boo := accessory.BillInfoFilter(w, data)
	if boo == false{
		return
	}

	data = filtetNull(data)

	parse_time, err := time.Parse("2006-01-02", data.BillDate)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data)


	result, err1 := dbslite.Db.Exec("insert into Bill values (:a, :b, :c, :d, :e)", data.BillId, data.BillMoney, data.BillCmt, parse_time, data.BillCategory)

	fmt.Print(result, err1)
	if err1 != nil{
		response.Code = accessory.ErrorAddInfoId
		response.Msg = accessory.ErrorAddInfoMsg

		return
	}

	response.Code = accessory.ErrorOKId
	response.Msg  = accessory.ErrorOKInsertMsg
	response.Data = data
}


//根据条数获取id
func getId() string{
	parse_time := time.Now()

	i := parse_time.Format("20060102150405")
	return i
}

//填充数据，防止向数据库中写入null
func filtetNull(data accessory.Bill) accessory.Bill {
	if data.BillCmt == ""{
		data.BillCmt = "无"
	}
	if data.BillCategory == ""{
		data.BillCategory = "无"
	}
	if data.BillDate == ""{
		data.BillDate = "2006-01-02"
	}

	return data
}

//账单数据库修改函数
func billUpdateMethod(w http.ResponseWriter, data accessory.Bill){
	//校验数据
	boo := accessory.BillInfoFilter(w, data)
	if boo == false{
		return
	}
	data = filtetNull(data)
	//判断数据库中是否有对应数据
	sql := "select * from Bill where b_id = :q"
	var bill accessory.Bill
	var parse_time time.Time

	err := dbslite.Db.QueryRow(sql, data.BillId).Scan(&bill.BillId, &bill.BillMoney,&bill.BillCmt, &parse_time, &bill.BillCategory)
	bill.BillDate = parse_time.Format("2006-01-02")
	fmt.Println(bill, err)
	if err != nil{
		response.Code = accessory.ErrorDataIsNilId
		response.Msg  = accessory.ErrorDataIsNilMsg
		return
	}

	sql = "update Bill set b_money = :a, b_cmt = :b, b_date = :c, b_category = :d where b_id = :e"

	parse_time, _ = time.Parse("2006-01-02", data.BillDate)
	result, err1 := dbslite.Db.Exec(sql, data.BillMoney, data.BillCmt, parse_time, data.BillCategory, data.BillId)

	fmt.Println(result)
	if err1 != nil{
		response.Code = accessory.ErrorUpdataInfoFailedId
		response.Msg = accessory.ErrorUpdataInfoFailedMsg
		return
	}

	response.Code = accessory.ErrorOKId
	response.Msg = accessory.ErrorOKModifyMsg
}

//账单数据库删除函数
func billDeleteMethod(w http.ResponseWriter, data accessory.Bill){
	//校验数据
	if data.BillId == ""{
		response.Code = accessory.ErrorGroup_idNilId
		response.Msg  = accessory.ErrorGroup_idNilMsg
		return 
	}

	//判断数据库中是否有对应数据
	sql := "select * from Bill where b_id = :q"
	var bill accessory.Bill
	var parse_time time.Time

	err := dbslite.Db.QueryRow(sql, data.BillId).Scan(&bill.BillId, &bill.BillMoney,&bill.BillCmt, &parse_time, &bill.BillCategory)
	bill.BillDate = parse_time.Format("2006-01-02")
	fmt.Println(bill, err)
	if err != nil || bill.BillId == "" {
		response.Code = accessory.ErrorDataIsNilId
		response.Msg  = accessory.ErrorDataIsNilMsg
		return
	}

	sql = "delete from Bill where b_id = :e"

	result, err1 := dbslite.Db.Exec(sql, data.BillId)

	fmt.Println(result)
	if err1 != nil{
		response.Code = accessory.ErrorUpdataInfoFailedId
		response.Msg = accessory.ErrorUpdataInfoFailedMsg
		return
	}

	response.Code = accessory.ErrorOKId
	response.Msg = accessory.ErrorOKDeleteMsg
}

//获取总钱数
func billGetMoney() string{
	sql := "select sum (b_money) from Bill"

	i := 0

	err:= dbslite.Db.QueryRow(sql).Scan(&i)

	if err != nil{
		log.Fatal(err)
	}

	return strconv.Itoa(i)
}
//获取所有账单
func getAllBill() []accessory.Bill{
	sql := "select * from Bill order by b_date desc"
	var bill accessory.Bill
	var bills []accessory.Bill
	var parse_time time.Time

	result, err := dbslite.Db.Query(sql)

	if err != nil{
		return bills
	}

	for result.Next(){
		err = result.Scan(&bill.BillId, &bill.BillMoney,&bill.BillCmt, &parse_time, &bill.BillCategory)
		bill.BillDate = parse_time.Format("2006-01-02")
		if err != nil{
			return bills
		}

		bills = append(bills, bill)
	}

	return bills
}
