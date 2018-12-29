package accessory

import ()

type BillRequest struct{
	Com 	string `json:"com"`
	Data 	Bill	`json:"data"`
}
