package model


type RespOk struct {
	Status []Status `json:"status"`
}

type Status struct {
	Ratio float64 `json:"ratio"`
	Ip    string  `json:"ip"`
	Port  int  `json:"port"`
}

type ResBad struct {
	Msg string `json:"msg"`
}
