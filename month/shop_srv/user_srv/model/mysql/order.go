package mysql

import (
	"gorm.io/gorm"
	"user_srv/from"
	"user_srv/utils"
)

// TODO：订单表字段设计
type Order struct {
	Id      int     `gorm:"primarykey"`
	UserID  int     `gorm:"column:user_id;type:int;not null;comment:用户id"`
	GoodsID int     `gorm:"column:goods_id;type:int;not null;comment:商品id"`
	OrderSn string  `gorm:"column:order_sn;type:varchar(50);not null;comment:订单编号"`
	Status  int     `gorm:"column:status;type:int;not null;comment:订单状态"`
	PayType int     `gorm:"column:pay_type;type:tinyint(1);not null;comment:支付类型"`
	PayTime string  `gorm:"column:pay_time;type:varchar(50);not null;comment:支付时间"`
	Total   float64 `gorm:"column:total;type:decimal(10,2);not null;comment:总金额"`
	Remark  string  `gorm:"column:remark;type:varchar(100);comment:订单备注"`
	gorm.Model
}

func GetOrder(userId int) (userOrder *from.UserOrder, err error) {
	err = utils.DB.Table("orders").Where("user_id = ?", userId).Select("name,mobile,order_sn,orders.status,total").Joins("left join users on users.id = user_id").Scan(&userOrder).Error
	return
}
