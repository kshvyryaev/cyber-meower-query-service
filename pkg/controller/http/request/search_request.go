package request

type SearchRequest struct {
	Query string `form:"query"`
	Skip  int    `form:"skip"`
	Take  int    `form:"take"`
}
