package models

import (
	"time"
)

type Admin struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Status    int64  `json:"status,omitempty"`
	Role      Role   `json:"role,omitempty"`
	Meta      Meta   `json:"meta,omitempty"`
}

type MetaInfo struct {
	ID        int64  `json:"id,omitempty"`
	TypeID    int64  `json:"type_id,omitempty"`
	MetaKey   string `json:"meta_key,omitempty"`
	MetaValue string `json:"meta_value,omitempty"`
}

type AdminMeta struct {
	ID        int64  `json:"id,omitempty"`
	AdminID   int64  `json:"admin_id,omitempty"`
	MetaKey   string `json:"meta_key,omitempty"`
	MetaValue string `json:"meta_value,omitempty"`
}

type UserMeta struct {
	ID        int64  `json:"id,omitempty"`
	UserID    int64  `json:"user_id,omitempty"`
	MetaKey   string `json:"meta_key,omitempty"`
	MetaValue string `json:"meta_value,omitempty"`
}

type Meta struct {
	BirthDate string `json:"birthdate,omitempty"`
	Address   string `json:"address,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

type Customer struct {
	ID        int64     `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Status    int64     `json:"status,omitempty"`
}

type Role struct {
	ID       int64  `json:"id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	Status   int64  `json:"status,omitempty"`
}

type JwtToken struct {
	Token string `json:"jwt_token,omitempty"`
}

type Shop struct {
	ID           int64  `json:"id,omitempty"`
	ShopName     string `json:"shop_name,omitempty"`
	OwnerID      int64  `json:"owner_id,omitempty"`
	CatID        int64  `json:"cat_id,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Status       int64  `json:"status,omitempty"`
	CategoryName string `json:"cat_name,omitempty"`
	Meta         ShopMeta
}

type ShopMeta struct {
	TelePhone         string `json:"tele_phone,omitempty"`
	Email             string `json:"email,omitempty"`
	Address           string `json:"address,omitempty"`
	Location          string `json:"location,omitempty"`
	NumberOfEmployers int64  `json:"number_of_employers,omitempty"`
}

type Product struct {
	ID          int64  `json:"id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	ShopID      int64  `json:"shop_id,omitempty"`
	Status      int64  `json:"status,omitempty"`
	Meta        ProductMeta
}

type ProductMeta struct {
	Avatar string `json:"avatar,omitempty"`
}

type Order struct {
	ID              int64     `json:"order_id"`
	CustomerID      int64     `json:"customer_id"`
	Products        string    `json:"products"`
	RefID           string    `json:"ref_id"`
	OrderToken      string    `json:"order_token"`
	Authority       string    `json:"authority"`
	TransportPrice  string    `json:"transport_price"`
	TotalPrice      string    `json:"total_price"`
	DiscountAmount  string    `json:"discount_amount,omitempty"`
	DiscountCoupon  string    `json:"discount_coupon,omitempty"`
	WalletReduction string    `json:"wallet_reduction,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	DatePaid        time.Time `json:"date_paid,omitempty"`
	Description     string    `json:"description"`
	Status          string    `json:"order_status"`
}

type ApiToken struct {
	ApiKey string `json:"api_key,omitempty"`
}

type Card struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Quantity    string `json:"qty"`
	Discount    string `json:"discount"`
	TotalPrice  string `json:"total_price"`
}

type FullCard struct {
	CardID     string `json:"card_id"`
	CustomerID string `json:"customer_id"`
	Card       string `json:"card"`
	CreatedAt  string `json:"created_at"`
	Status     string `json:"status"`
}

type Category struct {
	CatID   int64  `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	CatName string `json:"cat_name,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Parent  int64  `json:"parent,omitempty"`
	Status  int64  `json:"status,omitempty"`
}
