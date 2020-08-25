package products_lib

type Product struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Price    int64  `json:"price,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type ProductUpdate struct {
	Id       int64   `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Price    *int64  `json:"price,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
}

type ProductStore interface {
	List() ([]Product, error)
	Create(product *Product) (*Product, error)
	GetById(id int64) (*Product, error)
	Update(product *ProductUpdate) (*Product, error)
	Delete(id int64) error
}
