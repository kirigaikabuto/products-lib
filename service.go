package products_lib

type ProductService interface {
	ListProducts() ([]Product, error)
	CreateProduct(cmd *CreateProductCommand) (*Product, error)
	GetProductById(cmd *GetProductByIdCommand) (*Product, error)
	UpdateProduct(cmd *UpdateProductCommand) (*Product, error)
	DeleteProduct(cmd *DeleteProductCommand) error
}

type productService struct {
	productStore ProductStore
}

func NewProductService(productStore ProductStore) ProductService {
	return &productService{productStore: productStore}
}

func (ps *productService) ListProducts() ([]Product, error) {
	products, err := ps.productStore.List()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) CreateProduct(cmd *CreateProductCommand) (*Product, error) {
	product := &Product{
		Name:     cmd.Name,
		Price:    cmd.Price,
		ImageUrl: cmd.ImageUrl,
	}
	newProduct, err := ps.productStore.Create(product)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (ps *productService) GetProductById(cmd *GetProductByIdCommand) (*Product, error) {
	product, err := ps.productStore.GetById(cmd.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *productService) UpdateProduct(cmd *UpdateProductCommand) (*Product, error) {
	updateProduct := &ProductUpdate{}
	updateProduct.Id = cmd.Id
	if cmd.Price != nil {
		updateProduct.Price = cmd.Price
	}
	if cmd.Name != nil {
		updateProduct.Name = cmd.Name
	}
	if cmd.ImageUrl != nil {
		updateProduct.ImageUrl = cmd.ImageUrl
	}
	cmdGetProductById := &GetProductByIdCommand{cmd.Id}
	_, err := ps.GetProductById(cmdGetProductById)
	if err != nil {
		return nil, err
	}
	updatedProduct, err := ps.productStore.Update(updateProduct)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (ps *productService) DeleteProduct(cmd *DeleteProductCommand) error {
	cmdGetProductById := &GetProductByIdCommand{cmd.Id}
	_, err := ps.GetProductById(cmdGetProductById)
	if err != nil {
		return err
	}
	err = ps.productStore.Delete(cmd.Id)
	if err != nil {
		return err
	}
	return nil
}
