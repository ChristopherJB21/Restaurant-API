package model

type CuisineResponse struct {
	IDCuisine   uint   `json:"IDCuisine"`
	CuisineName string `json:"CuisineName"`
}

func ToCuisineResponse(item Cuisine) CuisineResponse {
	return CuisineResponse{
		IDCuisine:   item.ID,
		CuisineName: item.CuisineName,
	}
}

func ToCuisineResponses(items []Cuisine) []CuisineResponse {
	var cuisineResponses []CuisineResponse
	for _, item := range items {
		cuisineResponses = append(cuisineResponses, ToCuisineResponse(item))
	}
	return cuisineResponses
}
