package ozon

import "time"

type QuestionAnswerCreateRsp struct {
	BaseRsp
	AnswerId string `json:"answer_id"`
}

type QuestionAnswerDeleteRsp struct {
	BaseRsp
	Code    int    `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type QuestionAnswerListRsp struct {
	BaseRsp
	Answers []QuestionAnswerListItem `json:"answers"`
	LastId  string                   `json:"last_id"`
}

type QuestionAnswerListItem struct {
	AuthorName        string    `json:"author_name"`
	Id                string    `json:"id"`
	PublishedAt       time.Time `json:"published_at"`
	QuestionId        string    `json:"question_id"`
	Sku               int64     `json:"sku"`
	StatusPublication string    `json:"status_publication"`
	Text              string    `json:"text"`
}

type QuestionChangeStatusRsp struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type QuestionChangeCountRsp struct {
	BaseRsp
	All         int64 `json:"all"`
	New         int64 `json:"new"`
	Processed   int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
	Viewed      int64 `json:"viewed"`
}

type QuestionInfoRsp struct {
	BaseRsp
	AnswersCount string    `json:"answers_count"`
	AuthorName   string    `json:"author_name"`
	ProductUrl   string    `json:"product_url"`
	Sku          int64     `json:"sku"`
	Id           string    `json:"id"`
	Text         string    `json:"text"`
	QuestionLink string    `json:"question_link"`
	PublishedAt  time.Time `json:"published_at"`
	Status       string    `json:"status"`
}

type QuestionListRsp struct {
	BaseRsp
	HasNext   bool               `json:"has_next"`
	Questions []QuestionListItem `json:"questions"`
	LastId    string             `json:"last_id"`
}

type QuestionListItem struct {
	AnswersCount int64     `json:"answers_count"`
	AuthorName   string    `json:"author_name"`
	Id           string    `json:"id"`
	Sku          int64     `json:"sku"`
	ProductUrl   string    `json:"product_url"`
	PublishedAt  time.Time `json:"published_at"`
	QuestionLink string    `json:"question_link"`
	Text         string    `json:"text"`
	Status       string    `json:"status"`
}

type QuestionTopSkuRsp struct {
	BaseRsp
	Sku []int64 `json:"sku"`
}
