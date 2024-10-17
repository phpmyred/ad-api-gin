package response

import "github.com/goccy/go-json"

// ModelToDTO
// @Description  MODEL 表转 返回VO参数
// @Author aDuo 2024-09-02 14:45:10
// @Param m
// @Param dto
// @Return error

func ModelToVO(m interface{}, vo interface{}) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		panic("解析数据错误！")

	}
	err = json.Unmarshal(jsonData, &vo)

	if err != nil {
		panic("转化失败！")
	}
}

//
// ModelListToDTO[M any, DTO any]
// @Description  将结构体列表 化为另外一个结构体列表的方法
// @Author aDuo 2024-09-02 17:07:21
// @Param list
// @Return []DTO

func ModelListToVO[M any, VO any](list []M) []VO {
	out := make([]VO, len(list))
	for i, v := range list {
		var vo VO
		ModelToVO(v, &vo)
		out[i] = vo
	}
	return out
}
