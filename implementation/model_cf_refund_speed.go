package cashfree_pg_sdk_go

type CFRefundSpeed struct {
	RequestedSpeed *string `json:"requested"`
	AcceptedSpeed  string  `json:"accepted"`
	ProcessedSpeed *string `json:"processed"`
	Message        *string `json:"message"`
}
