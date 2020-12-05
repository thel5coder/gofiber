package viewmodel

type ResponseErrorVm struct {
	Messages interface{} `json:"messages"`
}

type ResponseSuccessVm struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}
