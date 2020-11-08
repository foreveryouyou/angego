package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type PipUtil struct {
	pipeline []bson.M
	match    bson.M
}

func NewPipUtil() (pl *PipUtil) {
	pl = &PipUtil{
		pipeline: []bson.M{},
		match:    bson.M{},
	}
	return
}

func (p *PipUtil) GetPipeline() []bson.M {
	if len(p.pipeline) < 1 {
		p.pipeline = append(p.pipeline, bson.M{"$skip": 0})
	}
	return p.pipeline
}

func (p *PipUtil) GetMatch() bson.M {
	return p.match
}

func (p *PipUtil) Match(match bson.M) *PipUtil {
	p.match = match
	if len(match) > 0 {
		p.pipeline = append(p.pipeline, bson.M{"$match": match})
	}
	return p
}

func (p *PipUtil) Sort(sort bson.M) *PipUtil {
	if len(sort) > 0 {
		p.pipeline = append(p.pipeline, bson.M{"$sort": sort})
	}
	return p
}

func (p *PipUtil) Skip(skip int64) *PipUtil {
	p.pipeline = append(p.pipeline, bson.M{"$skip": skip})
	return p
}

func (p *PipUtil) Limit(limit int64) *PipUtil {
	p.pipeline = append(p.pipeline, bson.M{"$limit": limit})
	return p
}

func (p *PipUtil) Project(projection bson.M) *PipUtil {
	if len(projection) > 0 {
		p.pipeline = append(p.pipeline, bson.M{"$project": projection})
	}
	return p
}

//Lookup
// 示例：
//{"$lookup":
// 	bson.M{
//		"from":         "表1",
//		"localField":   "当前表字段",
//		"foreignField": "表1对应关联字段",
//		"as":           "关联后结果中字段名",
//	}
// }
func (p *PipUtil) Lookup(lookup bson.M) *PipUtil {
	if len(lookup) > 0 {
		p.pipeline = append(p.pipeline, bson.M{"$lookup": lookup})
	}
	return p
}

//Unwind
// 示例：
//{"$unwind": "$salesmen"}
func (p *PipUtil) Unwind(unwind string) *PipUtil {
	if unwind != "" {
		p.pipeline = append(p.pipeline, bson.M{"$unwind": unwind})
	}
	return p
}

//Unwind
// 示例：
//{"$unwind": "$salesmen"}
func (p *PipUtil) UnwindEmpty(unwind string) *PipUtil {
	if unwind != "" {
		p.pipeline = append(p.pipeline, bson.M{"$unwind": bson.M{
			"path":                       unwind,
			"preserveNullAndEmptyArrays": true,
		}})
	}
	return p
}

func (p *PipUtil) GeoNear(geoNear GeoNear) *PipUtil {
	p.match = geoNear.Query
	p.pipeline = append(p.pipeline, bson.M{"$geoNear": geoNear})
	return p
}

// 暂未实现的操作先借由这个实现
func (p *PipUtil) AddStage(stage string, value interface{}) *PipUtil {
	p.pipeline = append(p.pipeline, bson.M{stage: value})
	return p
}

//
//GeoNear{
//	Near:               [2]float64{},
//	Query:              bson.M{},
//	Spherical:          true,
//	DistanceField:      "distance",
//	DistanceMultiplier: 6378137,
//	MaxDistance:        0,
//}
type GeoNear struct {
	Near               [2]float64 `bson:"near"`                  // []float64{121.522473, 31.264108}
	Query              bson.M     `bson:"query"`                 // 查询条件
	Spherical          bool       `bson:"spherical"`             // 是否计算球面距离
	DistanceField      string     `bson:"distanceField"`         // 结果中的距离字段名
	DistanceMultiplier int64      `bson:"distanceMultiplier"`    // 6378137
	MaxDistance        float64    `bson:"maxDistance,omitempty"` // 最大距离: m/6378137
}
