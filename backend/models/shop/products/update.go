package products

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func UpdateProduct(product models.Product) (*models.Product, error) {
	status := strconv.Itoa(int(product.Status))
	product.Slug = slug.Make(product.ProductName)
	productID := strconv.Itoa(int(product.ID))
	db := database.CreateCon()
	_, err := db.Exec(fmt.Sprintf("update ico_product set product_name='%v',slug='%v',status=%v where id=%v;", product.ProductName, product.Slug, status, productID))
	if len(product.Meta.Avatar) > 1 {
		err = UpdateProductMeta(product.ID, "avatar", product.Meta.Avatar, "")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateProductMeta(productID int64, metaKey string, metaValue string, currentValue string) error {
	db := database.CreateCon()
	if len(currentValue) > 1 {
		_, err := db.Exec(fmt.Sprintf("update ico_product_meta set meta_value='%v' where meta_key='%v' and product_id=%v and meta_value='%v'", metaValue, metaKey, productID, currentValue))
		if err != nil {
			return err
		}
	}
	_, err := db.Exec(fmt.Sprintf("update ico_product_meta set meta_value='%v' where meta_key='%v' and product_id=%v", metaValue, metaKey, productID))
	if err != nil {
		return err
	}
	return nil
}
