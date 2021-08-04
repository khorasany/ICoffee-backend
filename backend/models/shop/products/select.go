package products

import (
	"fmt"
	"github.com/khorasany/coffee/api/backend/database"
	"github.com/khorasany/coffee/api/backend/models"
)

func GetProductByProductID(productID string) (*models.Product, error) {
	var product models.Product
	db := database.CreateCon()
	err := db.QueryRow(fmt.Sprintf("select ico_product.id as id,ico_product.product_name as product_name,ico_product.slug as slug,ico_product.status as status,itps.shop_id as shop_id,ipm.meta_value as avatar from ico_product join ico_taxonomy_products_shop itps on ico_product.id = itps.product_id,ico_product_meta as ipm where ipm.meta_key='avatar' and ipm.product_id = ico_product.id and ico_product.id=%v", productID)).Scan(&product.ID, &product.ProductName, &product.Slug, &product.Status, &product.ShopID, &product.Meta.Avatar)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProductByOwnerID(ownerID string) ([]*models.Product, error) {
	db := database.CreateCon()
	productsResult, err := db.Query(fmt.Sprintf("select ip.id as id,ip.product_name as product_name,ip.slug as slug,ip.status as status,itps.shop_id as shop_id,icm.meta_value as avatar from ico_product as ip join ico_taxonomy_products_shop itps on ip.id = itps.product_id,ico_shop as shop,ico_product_meta as icm where icm.product_id = ip.id and icm.meta_key='avatar' and shop.id = itps.shop_id and shop.owner_id =%v;", ownerID))
	if err != nil {
		return nil, err
	}

	products := []*models.Product{}
	for productsResult.Next() {
		product := models.Product{}
		_ = productsResult.Scan(&product.ID, &product.ProductName, &product.Slug, &product.Status, &product.ShopID, &product.Meta.Avatar)
		products = append(products, &product)
	}
	return products, nil
}
