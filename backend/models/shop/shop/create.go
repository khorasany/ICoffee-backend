package shop

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"github.com/khorasany/coffee/api/backend/models/shop/category"
	"strconv"
)

func RegisterShop(shop models.Shop) (*models.Shop, error) {
	db := database.CreateCon()
	ownerID := strconv.Itoa(int(shop.OwnerID))
	status := strconv.Itoa(int(shop.Status))
	shopSlug := slug.Make(shop.ShopName)
	cat, _ := category.GetCategoryByCatName(shop.CategoryName)
	catID := strconv.Itoa(int(cat.CatID))
	shopRes, err := db.Exec(fmt.Sprintf("insert into ico_shop (owner_id, cat_id, shop_name, slug, status) values (%v,%v,'%v','%v',%v);", ownerID, catID, shop.ShopName, shopSlug, status))
	if err != nil {
		return nil, err
	}

	shopID, _ := shopRes.LastInsertId()
	strShop := strconv.Itoa(int(shopID))
	res := db.QueryRow(fmt.Sprintf("select * from ico_shop where id=%v;", strShop)).Scan(
		&shop.ID, &shop.OwnerID, &shop.CatID, &shop.ShopName, &shop.Slug, &shop.Status)
	if res != nil {
		return nil, res
	}
	return &shop, nil
}
