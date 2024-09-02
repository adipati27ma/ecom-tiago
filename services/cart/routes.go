package cart

import (
	"ecom-tiago/types"
	"ecom-tiago/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// injecting dependencies
type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore // use for checking product stock
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store, productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/cart", h.handleGetCart).Methods(http.MethodGet)
	// router.HandleFunc("/cart", h.handleAddToCart).Methods(http.MethodPost)
	// router.HandleFunc("/cart/{productID}", h.handleRemoveFromCart).Methods(http.MethodDelete)

	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := 0 // for now, we don't have user authentication

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSONRes(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check the items/products quantity available or not
	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get the products (on cart) from the database
	ps, err := h.productStore.GetProductsByIDs(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// create the order
	orderID, price, err := h.createOrder(ps, cart.Items, userID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"orderId":    orderID,
		"totalPrice": price,
	})
}
