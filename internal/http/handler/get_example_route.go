package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"alex/test/internal/http/response"
	"github.com/alex-dwt/go-course-service-2/pkg/repository/memory"
	"github.com/alex-dwt/go-course-service-2/pkg/user/discount"
	"go.uber.org/zap"
)

func GetExampleRouteFunc(logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDString := r.URL.Query().Get("user_id")
		if userIDString == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			io.WriteString(w, "userID is empty")
			return
		}

		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			io.WriteString(w, fmt.Sprintf("userID is not okay: %s", err))
			return
		}

		service := discount.New(logger, 11.1, memory.UserRepository{})
		balance, err := service.CalculateDiscountForUser(context.Background(), userID)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			io.WriteString(w, fmt.Sprintf("error getting user balance: %s", err))
			return
		}

		encoded, err := json.Marshal(response.GetUserBalanceResponse{
			Balance: balance,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
	}
}
