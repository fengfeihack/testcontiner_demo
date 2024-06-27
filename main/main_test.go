package main

import (
	"testcontainer_demo/dao"
	"testing"
	
	_ "testcontainer_demo/testhelper"
	
	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestQueryDataUseContainer(t *testing.T) {
	mockey.PatchConvey("23", t, func() {
		//初始化需要测试的表
		err := dao.DB.AutoMigrate(dao.Product{})
		assert.Nil(t, err)
		r := dao.NewRepository()
		//写入临时测试数据
		err = r.Create(dao.Product{
			Code:  "D42",
			Price: 1,
		})
		assert.Nil(t, err)
		//执行测试
		mockey.Mock(DoSomethingUseProduct).Return(nil).Build()
		product, err := QueryData()
		assert.Nil(t, err)
		assert.Equal(t, 1, product.Price)
	})
}

func TestQueryData(t *testing.T) {
	mockey.PatchConvey("22", t, func() {
		dao.DB = &gorm.DB{}
		mockey.Mock((*dao.Repository).Select).Return(dao.Product{
			Price: 1,
		}, nil).Build()
		defer mockey.UnPatchAll()
		mockey.Mock(DoSomethingUseProduct).Return(nil).Build()
		product, err := QueryData()
		assert.Nil(t, err)
		assert.Equal(t, 1, product.Price)
		dao.DB = nil
	})
}
