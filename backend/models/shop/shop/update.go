package shop

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func UpdateShop(shop models.Shop) (*models.Shop, error) {
	ownerID := strconv.Itoa(int(shop.OwnerID))
	catID := strconv.Itoa(int(shop.CatID))
	status := strconv.Itoa(int(shop.Status))
	shopID := strconv.Itoa(int(shop.ID))
	shopSlug := slug.Make(shop.ShopName)
	db := database.CreateCon()
	_, err := db.Exec(fmt.Sprintf("update ico_shop set owner_id=%v,cat_id=%v,shop_name='%v',slug='%v',status=%v where id=%v;", ownerID, catID, shop.ShopName, shopSlug, status, shopID))
	if err != nil {

		return nil, err
	}

	return &shop, nil
}
