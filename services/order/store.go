package cart

import (
	"database/sql"
	"ecom-tiago/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	// docs: implementasi query untuk membuat order
	result, err := s.db.Exec(`INSERT INTO orders
  (userId, total, status, address) VALUES (?, ?, ?, ?)`,
		order.UserID, order.Total, order.Status, order.Address)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	// docs: implementasi query untuk membuat order item
	_, err := s.db.Exec(`INSERT INTO order_items
  (orderId, productId, quantity, price) VALUES (?, ?, ?, ?)`,
		orderItem.OrderID, orderItem.ProductID,
		orderItem.Quantity, orderItem.Price)
	return err
}
