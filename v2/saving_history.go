package binance

package binance

import (
	"context"
	"net/http"
)

// AssetDividendService fetches the saving purchases
type SavingPurchaseHistoryService struct {
	c         *Client
	lendingType     *string
	startTime *int64
	endTime   *int64
	asset     *int
	current *int
	size *int
	recvWindow 		*int
}

func (s *AssetDividendService) LendingType(t string) *AssetDividendService {
	s.lendingType = &t
	return s
}

func (s *AssetDividendService) LendingType(t string) *AssetDividendService {
	s.lendingType = &t
	return s
}

func (s *AssetDividendService) StartTime(startTime int64) *AssetDividendService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *AssetDividendService) EndTime(endTime int64) *AssetDividendService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *AssetDividendService) Do(ctx context.Context) (*DividendResponseWrapper, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/assetDividend",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	} else {
		r.setParam("limit", 20)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(SavingPurchaseHistoryResponseWrapper)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SavingPurchaseHistoryResponseWrapper struct {
	Rows  *[]SavingPurchaseHistoryResponse `json:"rows"`
	Total int32               `json:"total"`
}

type SavingPurchaseHistoryResponse struct {
	ID     int64  `json:"id"`
	Amount string `json:"amount"`
	Asset  string `json:"asset"`
	Info   string `json:"enInfo"`
	Time   int64  `json:"divTime"`
	TranID int64  `json:"tranId"`
}
