package cart

import (
	"ecom-tiago/types"
	"fmt"
	"strconv"
)

func getCartItemsIDs(cart []types.CartItem) ([]string, error) {
	ids := make([]string, len(cart))
	for _, item := range cart {
		if item.Quantity == 0 {
			return nil, fmt.Errorf("quantity of product %d is 0", item.ProductID)
		}

		ids = append(ids, strconv.Itoa(item.ProductID))
	}
	return ids, nil
}

// Business logic decision (can do in a SQL/database transaction)
// 1. Check if the product is available
// 2. Check if the quantity is available
// 3. Calculate the total price
// 4. Create an order
// 5. Create order items
// 6. Update the product quantity
// 7. Return the order ID
func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (int,
	float64, error) {
	// docs: create a map of products for easy access
	//			 and generaly faster/better performance
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce the quantity of the products
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		err := h.productStore.UpdateProduct(product)
		if err != nil {
			return 0, 0, err
		}
	}

	// create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "dummy address",
	})
	if err != nil {
		return 0, 0, err
	}

	// create order items
	for _, item := range items {
		err := h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})

		if err != nil {
			return 0, 0, err
		}
	}

	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var totalPrice float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		totalPrice += float64(item.Quantity) * product.Price
	}
	return totalPrice
}
