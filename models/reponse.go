package models

type JSONDataResult struct {
	Code    int         `json:"code" `
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} //@name ResponseData

type JSONCodeResult struct {
	Code    int    `json:"code" `
	Status  string `json:"status"`
	Message string `json:"message"`
} //@name ResponseCode

type JSONDataResultWithPagination struct {
	Code   int         `json:"code" `
	Count  int         `json:"count"`
	Data   interface{} `json:"data"`
	Offset int         `json:"offset"`
	Status string      `json:"status"`
	Total  int         `json:"total"`
} //@name ResponseDataWithPagination
