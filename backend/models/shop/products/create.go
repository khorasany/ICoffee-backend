package products

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
	"strconv"
)

func RegisterProduct(product models.Product) (*models.Product, error) {
	status := strconv.Itoa(int(product.Status))
	product.Slug = slug.Make(product.ProductName)
	db := database.CreateCon()
	productID, err := db.Exec(fmt.Sprintf("insert into ico_product (product_name,slug,status) values ('%v','%v',%v);", product.ProductName, product.Slug, status))
	if err != nil {
		return nil, err
	}
	productIDInt, _ := productID.LastInsertId()
	product.ID = productIDInt
	err = InsertTaxonomyProductAndShop(product)
	if err != nil {
		return nil, err
	}
	err = InsertProductMeta(productIDInt, "avatar", product.Meta.Avatar)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func InsertTaxonomyProductAndShop(product models.Product) error {
	db := database.CreateCon()
	productID := strconv.Itoa(int(product.ID))
	shopID := strconv.Itoa(int(product.ShopID))
	_, err := db.Exec(fmt.Sprintf("insert into ico_taxonomy_products_shop (shop_id,product_id) values (%v,%v);", shopID, productID))
	if err != nil {
		return err
	}
	return nil
}

func InsertProductMeta(productID int64, metaKey string, metaValue string) error {
	db := database.CreateCon()
	strProductID := strconv.Itoa(int(productID))
	_, err := db.Exec(fmt.Sprintf("insert into ico_product_meta (product_id,meta_key,meta_value) values (%v,'%v','%v');", strProductID, metaKey, metaValue))
	if err != nil {
		return err
	}
	return nil
}
