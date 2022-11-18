package repositories

import (
	"database/sql"
	"product/datamodels"
)

//开发接口
type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error) //获取插入后的ID
	Delete(int64) error
	Update(*datamodels.Product) (int64, error)
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func (p *ProductManager) Conn() (err error) {

}

func (p *ProductManager) Insert(product *datamodels.Product) (int64, error) {

}

func (p *ProductManager) Delete(ID int64) error {

}

func (p *ProductManager) Update(product *datamodels.Product) (int64, error) {

}

func (p *ProductManager) SelectByKey(ID int64) (*datamodels.Product, error) {

}

func (p *ProductManager) SelectAll() ([]*datamodels.Product, error) {

}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table: table, mysqlConn: db}
}
