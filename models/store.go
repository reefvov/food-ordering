package models

type Store struct {
	StoreId          string `json:"storeId"`
	StoreName        string `json:"storeName"`
	StorePhone       string `json:"storePhone"`
	StoreDescription string `json:"storeDescription"`
	StorePicture     string `json:"profilePicture"`
	IsActive         string `json:"isActive"`
}
