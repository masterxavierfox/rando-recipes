package actions

import (
	"encoding/json"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

const mealdbURL = "https://www.themealdb.com/api/json/v1/1/random.php"

type Themealdb struct {
	Meals []struct {
		IDMeal                      string      `json:"idMeal"`
		StrMeal                     string      `json:"strMeal"`
		StrDrinkAlternate           interface{} `json:"strDrinkAlternate"`
		StrCategory                 string      `json:"strCategory"`
		StrArea                     string      `json:"strArea"`
		StrInstructions             string      `json:"strInstructions"`
		StrMealThumb                string      `json:"strMealThumb"`
		StrTags                     string      `json:"strTags"`
		StrYoutube                  string      `json:"strYoutube"`
		StrIngredient1              string      `json:"strIngredient1"`
		StrIngredient2              string      `json:"strIngredient2"`
		StrIngredient3              string      `json:"strIngredient3"`
		StrIngredient4              string      `json:"strIngredient4"`
		StrIngredient5              string      `json:"strIngredient5"`
		StrIngredient6              string      `json:"strIngredient6"`
		StrIngredient7              string      `json:"strIngredient7"`
		StrIngredient8              string      `json:"strIngredient8"`
		StrIngredient9              string      `json:"strIngredient9"`
		StrIngredient10             string      `json:"strIngredient10"`
		StrIngredient11             string      `json:"strIngredient11"`
		StrIngredient12             string      `json:"strIngredient12"`
		StrIngredient13             string      `json:"strIngredient13"`
		StrIngredient14             string      `json:"strIngredient14"`
		StrIngredient15             string      `json:"strIngredient15"`
		StrIngredient16             interface{} `json:"strIngredient16"`
		StrIngredient17             interface{} `json:"strIngredient17"`
		StrIngredient18             interface{} `json:"strIngredient18"`
		StrIngredient19             interface{} `json:"strIngredient19"`
		StrIngredient20             interface{} `json:"strIngredient20"`
		StrMeasure1                 string      `json:"strMeasure1"`
		StrMeasure2                 string      `json:"strMeasure2"`
		StrMeasure3                 string      `json:"strMeasure3"`
		StrMeasure4                 string      `json:"strMeasure4"`
		StrMeasure5                 string      `json:"strMeasure5"`
		StrMeasure6                 string      `json:"strMeasure6"`
		StrMeasure7                 string      `json:"strMeasure7"`
		StrMeasure8                 string      `json:"strMeasure8"`
		StrMeasure9                 string      `json:"strMeasure9"`
		StrMeasure10                string      `json:"strMeasure10"`
		StrMeasure11                string      `json:"strMeasure11"`
		StrMeasure12                string      `json:"strMeasure12"`
		StrMeasure13                string      `json:"strMeasure13"`
		StrMeasure14                string      `json:"strMeasure14"`
		StrMeasure15                string      `json:"strMeasure15"`
		StrMeasure16                interface{} `json:"strMeasure16"`
		StrMeasure17                interface{} `json:"strMeasure17"`
		StrMeasure18                interface{} `json:"strMeasure18"`
		StrMeasure19                interface{} `json:"strMeasure19"`
		StrMeasure20                interface{} `json:"strMeasure20"`
		StrSource                   string      `json:"strSource"`
		StrImageSource              interface{} `json:"strImageSource"`
		StrCreativeCommonsConfirmed interface{} `json:"strCreativeCommonsConfirmed"`
		DateModified                interface{} `json:"dateModified"`
	} `json:"meals"`
}

// FeedFoodfeed default implementation.
//func FeedFoodfeed(c buffalo.Context) error {
//	return c.Render(http.StatusOK, r.HTML("feed/foodfeed.html"))
//}

// method to get the data from the API https://www.themealdb.com/api/json/v1/1/random.php and store response in struct Themealdb
func FeedFoodfeed(c buffalo.Context) error {
	var themealdb Themealdb

	resp, err := http.Get(mealdbURL)

	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&themealdb); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	strArea := themealdb.Meals[0].StrArea
	strMealThumbnail := themealdb.Meals[0].StrMealThumb
	strMeal := themealdb.Meals[0].StrMeal
	strYoutube := themealdb.Meals[0].StrYoutube
	StrCategory := themealdb.Meals[0].StrCategory
	StrTags := themealdb.Meals[0].StrTags
	strInstructions := themealdb.Meals[0].StrInstructions

	c.Set("area", strArea)
	c.Set("thumbnail", strMealThumbnail)
	c.Set("meal", strMeal)
	c.Set("video", strYoutube)
	c.Set("category", StrCategory)
	c.Set("tags", StrTags)
	c.Set("instructions", strInstructions)

	//return c.Render(http.StatusOK, r.JSON(strArea+" "+strMealThumbnail+" "+strMeal+" "+strYoutube))

	return c.Render(http.StatusOK, r.HTML("feed/foodfeed.html"))
}
