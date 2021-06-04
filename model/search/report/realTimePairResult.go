package report

import "encoding/json"

// 配对报告返回
type RealTimePairResult struct {
	KeywordId  int64         `json:"keywordId,omitempty"`  // 关键词ID
	CreativeId int64         `json:"creativeId,omitempty"` // 创意ID
	PairInfo   []string      `json:"pairInfo,omitempty"`   // 账户名，计划名，单元名，关键词字面，创意标题,创意描述一，创意描述二，创意显示url ，搜索引擎，精确匹配扩展(地域词扩展)触发
	Date       string        `json:"date,omitempty"`       // 统计开始时间
	KPIs       []json.Number `json:"KPIs,omitempty"`       // 按照请求顺序，返回KPI数据数组
}