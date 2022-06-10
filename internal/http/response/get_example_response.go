package response

import "time"

type GetExampleResponse struct {
	Field1 string                          `json:"field_1"`
	Field2 int                             `json:"field_2"`
	Cars   []GetExampleResponseInnerCarObj `json:"cars_array"`
}

type GetExampleResponseInnerCarObj struct {
	Model      string    `json:"car_model"`
	Year       int       `json:"year"`
	LastUsedBy time.Time `json:"last-used-by"`
}

func GenerateFakeGetResponse() GetExampleResponse {
	return GetExampleResponse{
		Field1: "value-of-filed-1",
		Field2: 1999999,
		Cars: []GetExampleResponseInnerCarObj{
			{
				Model:      "audi",
				Year:       666,
				LastUsedBy: time.Now(),
			},
			{
				Model:      "reno",
				Year:       777,
				LastUsedBy: time.Now().AddDate(-2, -2, -5),
			},
		},
	}
}
