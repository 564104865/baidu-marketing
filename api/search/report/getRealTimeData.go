package report

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/search/report"
)

// GetRealTimeData 推广报告
func GetRealTimeData(clt *core.SDKClient, auth model.RequestHeader, realTimeRequest *report.RealTimeRequest) (*model.ResponseHeader, []report.RealTimeResult, error) {
	req := &model.Request{
		Header: auth,
		Body: report.GetRealTimeDataRequest{
			RealTimeRequestType: realTimeRequest,
		},
	}
	var resp report.GetRealTimeDataResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
