package common

type QueryParameters struct {
	Skip int64 `query:"skip"`
	Take int64 `query:"take"`
}

type Response struct {
	Success    bool
	Error      string
	InnerError string
}

type PagedResponse struct {
	Result interface{} `json:"result"`
	End    bool        `json:"end"`
}
