package httpHandler

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-demo/internal/grpc/client"
	"grpc-demo/pb/userpb"
	"net/http"
)

type Handler struct {
	userClient    *client.UserClient
	productClient *client.ProductClient
}

func NewHandler(userClient *client.UserClient, productClient *client.ProductClient) *Handler {
	return &Handler{
		userClient:    userClient,
		productClient: productClient,
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// lấy token từ REST request
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "missing authorization header", http.StatusUnauthorized)
		return
	}

	// tạo context có token
	md := metadata.New(map[string]string{
		"Authorization": authHeader,
	})
	ctx := metadata.NewOutgoingContext(r.Context(), md)

	user := userpb.CreateUserRequest{
		Name:  "quochoai",
		Email: "dpquochoai@gmail.com",
	}
	userSaved, err := h.userClient.CreateUser(
		ctx,
		&user,
	)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {

			case codes.Unauthenticated:
				http.Error(w, "unauthenticated", http.StatusUnauthorized)
				return

			case codes.PermissionDenied:
				http.Error(w, "permission denied", http.StatusForbidden)
				return

			case codes.InvalidArgument:
				http.Error(w, st.Message(), http.StatusBadRequest)
				return
			}
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userSaved)
}

func (h *Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	products, err := h.productClient.ListProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
