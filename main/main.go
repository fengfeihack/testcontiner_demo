package main

import (
	"fmt"
	"testcontainer_demo/dao"
)

func init() {
	if dao.DB != nil {
		return
	}
	db, err := dao.OpenDB("")
	if err != nil {
		panic(err)
	}
	dao.DB = db
}

func QueryData() (*dao.Product, error) {
	r := dao.NewRepository()
	product, err := r.Select()
	if err != nil {
		return nil, err
	}
	err = DoSomethingUseProduct(product)
	return &product, err
}

func DoSomethingUseProduct(product dao.Product) error {
	//todo
	fmt.Println(product)
	return nil
}
