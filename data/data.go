package data

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Responses struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    []Student `json:"data"`
}
