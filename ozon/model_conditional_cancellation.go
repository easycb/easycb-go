package ozon

import "time"

type ConditionalCancellationListRsp struct {
	BaseRsp
	Result  []ConditionalCancellationListResultItem `json:"result"`
	Counter string                                  `json:"counter"`
	LastId  int64                                   `json:"last_id"`
}

type ConditionalCancellationListResultItem struct {
	ApproveComment        string    `json:"approve_comment"`
	ApproveDate           time.Time `json:"approve_date"`
	AutoApproveDate       time.Time `json:"auto_approve_date"`
	CancellationId        int64     `json:"cancellation_id"`
	CancellationInitiator string    `json:"cancellation_initiator"`
	CancellationReason    struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"cancellation_reason"`
	CancellationReasonMessage string    `json:"cancellation_reason_message"`
	CancelledAt               time.Time `json:"cancelled_at"`
	OrderDate                 time.Time `json:"order_date"`
	PostingNumber             string    `json:"posting_number"`
	State                     struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		State string `json:"state"`
	} `json:"state"`
	TplIntegrationType string `json:"tpl_integration_type"`
}

type ConditionalCancellationApproveRsp struct {
	BaseRsp
}

type ConditionalCancellationRejectRsp struct {
	BaseRsp
}
