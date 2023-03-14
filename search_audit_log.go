package icigo

import (
	"context"

	"net/url"
)

//searchAuditLog
//supported query parameters
//userName
//userId
//toDate
//pageSize int
//pageNumber int
//orderBy
//fromDate 2021-10-25T14:12:06+0000
//eventName
//entityType repeatable header
//entityId
//caseId

func (self *Client) SearchAuditLog(ctx context.Context, qp map[string]string) ([]byte, error) {
	_url := `/als/api/v1/auditlogs/search`
	base, err := url.Parse(_url)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	if qp != nil {
		for k, v := range qp {
			q.Add(k, v)
		}
	}

	base.RawQuery = q.Encode()

	return self.GetBytes(ctx, base.String())
}
