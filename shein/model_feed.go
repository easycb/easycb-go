package shein

type CreateFeedDocumentRsp struct {
	Info struct {
		Result struct {
			FeedDocumentId string `json:"feedDocumentId"`
			Url            string `json:"url"`
		} `json:"result"`
	} `json:"info"`
}

type GetFeedDocumentRsp struct {
	Info struct {
		Result struct {
			FeedDocumentId string `json:"feedDocumentId"`
			Url            string `json:"url"`
		} `json:"result"`
	} `json:"info"`
}

type UploadDocumentContentRsp struct {
	BaseRsp
	Info []struct {
		NPid string `json:"npid"`
	} `json:"info"`
}

type CreateFeedRsp struct {
	BaseRsp
	Info struct {
		Result int64 `json:"result"`
	} `json:"info"`
}

type GetFeedRsp struct {
	BaseRsp
	Info struct {
		Result struct {
			FeedId               int64  `json:"feedId"`
			FeedType             string `json:"feedType"`
			CreatedTime          string `json:"createdTime"`
			ProcessingStatus     string `json:"processingStatus"`
			ProcessingStartTime  string `json:"processingStartTime"`
			ProcessingEndTime    string `json:"processingEndTime"`
			ResultFeedDocumentId string `json:"resultFeedDocumentId"`
			ResultDocumentUrl    string `json:"resultDocumentUrl"`
		} `json:"result"`
	} `json:"info"`
}

type CancelFeedRsp struct {
	BaseRsp
}
