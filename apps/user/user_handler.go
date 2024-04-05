package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// type UserHandler struct {
// 	UserRepo UserRepository
// }

type User struct {
	RepoUser UserRepository
}

func (uh *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// userIDStr := r.URL.Query().Get("id")
	userIDStr := chi.URLParam(r, "id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	fmt.Println(userID)

	w.Write([]byte("user/GetUserByIDHandler!"))
}
