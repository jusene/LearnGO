package models

type Ress struct {
	Names []Res `json:"names"`
}

type Res struct {
	Name string `json:"name,omitempty"`
	Msg  string `json:"msg,omitempty" `
}

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty" `
}
