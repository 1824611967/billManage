package accessory

import (
	"encoding/json"
	"net/http"
)

func BillInfoFilter(w http.ResponseWriter, bill Bill) bool{
	var response Response
	boo := true

	if bill.BillId == ""{
		response.Code = ErrorGroup_idNilId
		response.Msg  = ErrorGroup_idNilMsg

		dataStr, _ := json.Marshal(response)
		w.Write([]byte(dataStr))

		boo = false

		return boo
	}

	if bill.BillMoney == 0{
		response.Code = ErrorGroup_idNilId
		response.Msg  = ErrorGroup_idNilMsg

		dataStr, _ := json.Marshal(response)

		w.Write([]byte(dataStr))

		boo = false

		return boo
	}

	return boo
}