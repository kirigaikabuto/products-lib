package products_lib

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS products (
		id serial,
		name text,
		price int,
		image_url text,
		PRIMARY KEY(id)
	);`,
}

type postgreStore struct {
	db *sql.DB
}

func NewPostgreStore(cfg Config) (ProductStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgreStore{db: db}, err
}

func (ps *postgreStore) List() ([]Product, error) {
	var products []Product
	data, err := ps.db.Query("select id,name,price,image_url from products")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		product := Product{}
		err = data.Scan(&product.Id, &product.Name, &product.Price, &product.ImageUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (ps *postgreStore) Create(product *Product) (*Product, error) {
	err := ps.db.QueryRow("insert into products (name,price,image_url) values ($1,$2,$3) RETURNING id", product.Name, product.Price, product.ImageUrl).Scan(&product.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *postgreStore) GetById(id int64) (*Product, error) {
	product := &Product{}
	err := ps.db.QueryRow("select id,name,price,image_url from products where id= $1", id).Scan(&product.Id, &product.Name, &product.Price, &product.ImageUrl)
	if err != nil {
		return nil, err
	}
	if product.Id == 0 {
		fmt.Println("go to error with product")
		return nil, errors.New("no data by id")
	}
	return product, nil
}

func (ps *postgreStore) Update(product *ProductUpdate) (*Product, error) {
	query := "update products set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if product.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, product.Name)
	}
	if product.Price != nil {
		cnt++
		parts = append(parts, "price = $"+strconv.Itoa(cnt))
		values = append(values, product.Price)
	}
	if product.ImageUrl != nil {
		cnt++
		parts = append(parts, "image_url = $"+strconv.Itoa(cnt))
		values = append(values, product.ImageUrl)
	}
	if len(parts) <= 0 {
		return nil, errors.New("nothing to update")
	}
	cnt++
	query = query + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, product.Id)
	result, err := ps.db.Exec(query, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("product not found")
	}

	return ps.GetById(product.Id)
}

func (ps *postgreStore) Delete(id int64) error {
	_, err := ps.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
