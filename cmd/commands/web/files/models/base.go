package models

type TimeFields struct {
	CreatedAt int64 `bson:"createdAt" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt int64 `bson:"updatedAt" json:"updatedAt,omitempty"` // 更新时间
	DeletedAt int64 `bson:"deletedAt" json:"-,omitempty"`         // 删除时间
}

//InitModels 需要初始化的操作
func InitModels() {

}
