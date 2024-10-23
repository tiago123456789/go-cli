package types

type Todos struct {
	Items []Todo `json:"todos"`
}

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}
