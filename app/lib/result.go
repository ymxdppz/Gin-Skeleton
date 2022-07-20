package lib

type result struct {
	Code   int            `json:"code"`
	Status bool           `json:"status"`
	Msg    string         `json:"msg"`
	Data   map[string]any `json:"data"`
}

func NewResult() *result {
	return &result{
		Code:   200,
		Status: true,
		Msg:    "ok",
		Data:   nil,
	}
}

func (p *result) GetResult() result {
	return *p
}

func (p *result) SetErrInfo(code int, msg string) *result {
	p.Code = code
	p.Msg = msg
	return p
}

func (p *result) SetData(data map[string]any) *result {
	p.Data = data
	return p
}
