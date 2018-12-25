package main

import (
	"dbslite"
	"net/http"
	"accessory"
	"time"
	"fmt"
	"log"
	"strconv"
)

//账单数据库插入函数
func billAddMethod(w http.ResponseWriter, data accessory.Bill){
	data.BillId = getId()
	//校验数据
	boo := accessory.BillInfoFilter(w, data)
	if boo == false{
		return
	}
	parse_time, _ := time.Parse("2006-01-02 15:04:05", data.BillDate)

	fmt.Println(data)

	result, err1 := dbslite.Db.Exec("insert into Bill values (:a, :b, :c, :d, :e)", data.BillId, data.BillMoney, data.BillCmt, parse_time, data.BillCategory)

	fmt.Print(result, err1)
	if err1 != nil{
		response.Code = accessory.ErrorAddInfoId
		response.Msg = accessory.ErrorAddInfoMsg

		return
	}

	response.Code = accessory.ErrorOKId
	response.Msg  = accessory.ErrorOKMsg 
}


//根据条数获取id
func getId() string{
	sql := "SELECT  /*+ROWID(Bill)*/ count(*)  FROM Bill t "
	i := 0
	err:= dbslite.Db.QueryRow(sql).Scan(&i)

	if err != nil{
		log.Fatal(err)
	}

	return strconv.Itoa(i)
}

//账单数据库修改函数
func billUpdateMethod(w http.ResponseWriter, data accessory.Bill){
	//校验数据
	boo := accessory.BillInfoFilter(w, data)
	if boo == false{
		return
	}

	//判断数据库中是否有对应数据
}