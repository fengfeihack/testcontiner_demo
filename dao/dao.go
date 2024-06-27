package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Product struct {
	Code  string
	Price int
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func OpenDB(dbUrl string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func (r *Repository) Select() (Product, error) {
	var product Product
	err := DB.First(&product, "code = ?", "D42").Error // 查找 code 字段值为 D42 的记录
	return product, err
}

func (r *Repository) Create(product Product) error {
	return DB.Create(&product).Error
}
