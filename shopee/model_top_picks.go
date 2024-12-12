package shopee

type GetTopPicksListRsp struct {
	BaseRsp
	Response struct {
		CollectionList []struct {
			IsActivated bool `json:"is_activated"`
			ItemList    []struct {
				ItemName                    string `json:"item_name"`
				ItemId                      int64  `json:"item_id"`
				CurrentPrice                string `json:"current_price"`
				InflatedPriceOfCurrentPrice string `json:"inflated_price_of_current_price"`
				Sales                       int    `json:"sales"`
			} `json:"item_list"`
			TopPicksId int    `json:"top_picks_id"`
			Name       string `json:"name"`
		} `json:"collection_list"`
	} `json:"response"`
}

type AddTopPicksRsp struct {
	BaseRsp
	Response struct {
		CollectionList []struct {
			IsActivated bool `json:"is_activated"`
			ItemList    []struct {
				ItemName                    string `json:"item_name"`
				ItemId                      int64  `json:"item_id"`
				CurrentPrice                string `json:"current_price"`
				InflatedPriceOfCurrentPrice string `json:"inflated_price_of_current_price"`
				Sales                       int    `json:"sales"`
			} `json:"item_list"`
			TopPicksId int    `json:"top_picks_id"`
			Name       string `json:"name"`
		} `json:"collection_list"`
	} `json:"response"`
}

type UpdateTopPicksRsp struct {
	BaseRsp
	Response struct {
		CollectionList []struct {
			IsActivated bool `json:"is_activated"`
			ItemList    []struct {
				ItemName                    string `json:"item_name"`
				ItemId                      int64  `json:"item_id"`
				CurrentPrice                string `json:"current_price"`
				InflatedPriceOfCurrentPrice string `json:"inflated_price_of_current_price"`
				Sales                       int    `json:"sales"`
			} `json:"item_list"`
			TopPicksId int    `json:"top_picks_id"`
			Name       string `json:"name"`
		} `json:"collection_list"`
	} `json:"response"`
}

type DeleteTopPicksRsp struct {
	BaseRsp
	Response struct {
		TopPicksId int `json:"top_picks_id"`
	} `json:"response"`
}
