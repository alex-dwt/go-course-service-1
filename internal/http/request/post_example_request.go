package request

type PostExampleRequest struct {
	Year  int                         `json:"year"`
	Items []PostExampleRequestOneItem `json:"items_list"`
}

type PostExampleRequestOneItem struct {
	Name string                  `json:"name"`
	Cars []PostExampleRequestCar `json:"cars"`
}

type PostExampleRequestCar struct {
	Year int    `json:"year"`
	Id   string `json:"id"`
}
