package category

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func RegisterCategory(category models.Category) (*models.Category, error) {
	db := database.CreateCon()
	parent := strconv.Itoa(int(category.Parent))
	status := strconv.Itoa(int(category.Status))
	slugCat := slug.Make(category.CatName)
	cat, err := db.Exec(fmt.Sprintf("insert into ico_category (type, category_name, slug, parent, status) values ('%v','%v','%v',%v,%v);", category.Type, category.CatName, slugCat, parent, status))
	if err != nil {
		return nil, err
	}
	catID, _ := cat.LastInsertId()
	category.CatID = catID
	category.Slug = slugCat
	return &category, nil
}
