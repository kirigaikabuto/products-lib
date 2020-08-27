package products_lib

type ListProductCommand struct {
}

func (cmd *ListProductCommand) Exec(service ProductService) (interface{}, error) {
	return service.ListProducts()
}

type CreateProductCommand struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func (cmd *CreateProductCommand) Exec(service ProductService) (interface{}, error) {
	return service.CreateProduct(cmd)
}

type GetProductByIdCommand struct {
	Id int64 `json:"id"`
}

func (cmd *GetProductByIdCommand) Exec(service ProductService) (interface{}, error) {
	return service.GetProductById(cmd)
}

type UpdateProductCommand struct {
	Id       int64   `json:"id"`
	Name     *string `json:"name"`
	Price    *int64  `json:"price"`
	ImageUrl *string `json:"image_url"`
}

func (cmd *UpdateProductCommand) Exec(service ProductService) (interface{}, error) {
	return service.UpdateProduct(cmd)
}

type DeleteProductCommand struct {
	Id int64 `json:"id"`
}
