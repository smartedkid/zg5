package mysql

import "gorm.io/gorm"

// todo: 商品表
type Goods struct {
	Id         int     `gorm:"primary"`
	GoodsName  string  `gorm:"column:goods_name;type:varchar(20);comment:商品名称"`
	GoodsPrice float64 `gorm:"column:goods_price;type:decimal(10,2);comment:商品价格"`
	GoodsImg   string  `gorm:"column:goods_img;type:varchar(200);comment:商品封面"`
	GoodsStock int     `gorm:"column:goods_stock;type:int;comment:商品库存"`
	GoodsType  int     `gorm:"column:goods_type;type:tinyint;comment:商品类型"`
	IsHot      int     `gorm:"column:is_hot;type:tinyint;comment:是否热门"`
	Status     int     `gorm:"column:status;type:tinyint;comment:商品状态"`
	gorm.Model
}
