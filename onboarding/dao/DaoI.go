package dao

// 数据层操作
type Dao interface {
	// 保存数据
	add(po interface{})
	// 条件查询数据，获取多条
	query(po interface{}) []interface{}
	// 条件查询单条数据，一般是根据唯一索引或id
	findByCondition(po interface{}) interface{}
	// 修改数据
	update(po interface{})
}
