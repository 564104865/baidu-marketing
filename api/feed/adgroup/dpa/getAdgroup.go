package dpa

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/feed/adgroup/dpa"
)

// GetAdgroup 查询商品推广单元
func GetAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *dpa.GetAdgroupFeedRequest) (*model.ResponseHeader, []dpa.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp dpa.GetAdgroupFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
