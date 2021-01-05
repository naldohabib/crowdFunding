package model

type Respons struct {
	//Success bool
	//Message string      `json:"Message"`
	//Data    interface{} `json:"data,omitempty"`

	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int `json:"code"`
	Status  bool `json:"status"`
}

