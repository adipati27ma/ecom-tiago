package product

import (
	"database/sql"
	"ecom-tiago/types"
	"fmt"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

// docs: codeium says, this is example of data access method
func (s *Store) GetProducts() ([]types.Product, error) {
	// docs: implementasi query untuk mendapatkan produk
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // menutup rows apapun yang terjadi

	products := make([]types.Product, 0) // membuat slice kosong dgn size 0
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p) // dereferencing p, when want to work with the value that pointer points to
	}
	fmt.Println(products)

	return products, nil
}

func (s *Store) GetProductsByIDs(ids []string) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(ids)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)
	fmt.Println("ids placeholders", placeholders)

	// docs: Convert ids to []interface{}
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil

	// docs: bisa langsung dengan kode berikut,
	// tapi query-nya akan banyak dan berkali-kali (berat, performance issue)
	// products := make([]types.Product, 0)
	// for _, id := range ids {
	// 	product, err := s.GetProductByID(id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	products = append(products, *product)
	// }
	// return products, nil
}

func (s *Store) GetProductByID(id string) (*types.Product, error) {
	// docs: implementasi query untuk mendapatkan produk berdasarkan ID
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product := new(types.Product)
	for rows.Next() {
		product, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("this is ID", id)

	return product, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	// docs: implementasi query untuk membuat produk
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		p.Name, p.Description, p.ImageURL, p.Price, p.Quantity)
	return err
}

func (s *Store) UpdateProduct(p types.Product) error {
	_, err := s.db.Exec(`UPDATE products SET name = ?, description = ?,
	image = ?, price = ?, quantity = ? WHERE id = ?`,
		p.Name, p.Description, p.ImageURL, p.Price, p.Quantity, p.ID)
	return err
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product) // make new pointer to Product (return an alocated memory of Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.ImageURL,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
