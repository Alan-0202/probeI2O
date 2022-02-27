package model



//source : www.baidu.com_80_3_5_5_80_china1
//convert to Source struct
type Source struct {
	Address string `json:"address"`  //www.baidu.com
	Port int `json:"port"`  //80
	Timeout int `json:"timeout"` //3
	Concurrence int `json:"concurrence"` //5
	Counts int `json:"counts"` //5
	Threshold int `json:"threshold"`  //80
	Shopname string `json:"shopname"` //China1
}


