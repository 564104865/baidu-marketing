package adgroup

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/search/adgroup"
)

// DeleteAdgroup 删除单元
func DeleteAdgroup(clt *core.SDKClient, auth model.RequestHeader, adgroupIds []int64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: adgroup.DeleteAdgroupRequest{
			AdgroupIds: adgroupIds,
		},
	}

	return clt.Do(req, nil)
}
