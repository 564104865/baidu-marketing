package account

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/feed/account"
)

// GetAccountFeed 查询账户
func GetAccountFeed(clt *core.SDKClient, auth model.RequestHeader, accountFields []string) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: account.GetAccountFeedRequest{
			AccountFeedFields: accountFields,
		},
	}
	var resp account.GetAccountFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
