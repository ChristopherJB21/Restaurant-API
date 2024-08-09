package model

type FoodResponse struct {
	IDFood      uint   `json:"IDFood"`
	FoodName    string `json:"FoodName"`
	IDCuisine   uint   `json:"IDCuisine"`
	CuisineName string `json:"CuisineName"`
}

func ToFoodResponse(item Food) FoodResponse {
	return FoodResponse{
		IDFood:      item.ID,
		FoodName:    item.FoodName,
		IDCuisine:   item.IDCuisine,
		CuisineName: item.Cuisine.CuisineName,
	}
}

func ToFoodResponses(items []Food) []FoodResponse {
	var foodResponses []FoodResponse
	for _, item := range items {
		foodResponses = append(foodResponses, ToFoodResponse(item))
	}
	return foodResponses
}
