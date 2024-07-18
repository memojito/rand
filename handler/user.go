package handler

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/memojito/igapi/storage"
	"github.com/memojito/igapi/types"
	"github.com/memojito/igapi/utils"
)

type UserHandler struct {
	storage storage.Storage
}

func New(s storage.Storage) *UserHandler {
	return &UserHandler{
		storage: s,
	}
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return err
	}
	user := types.User{
		ID:    id,
		Name:  "masha",
		Email: "masha@gmail.com",
	}
	log.Println(user.ID)
	b, _ := json.Marshal(user)
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	payload := &types.ListUsersRequest{}
	if err := utils.ParseJSON(r, payload); err != nil {
		return err
	}

	users, err := h.storage.List(r.Context(), payload.IDs)
	if err != nil {
		return err
	}

	err = utils.WriteJSON(w, 200, users)
	if err != nil {
		return err
	}
	return nil
}

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) error {
	user1 := types.User{
		ID:    rand.Int(),
		Name:  "masha",
		Email: "masha@gmail.com",
	}
	user2 := types.User{
		ID:    rand.Int(),
		Name:  "sasha",
		Email: "sasha@gmail.com",
	}
	users := []types.User{user1, user2}
	b, _ := json.Marshal(users)
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (h UserHandler) AddUser(w http.ResponseWriter, r *http.Request) error {
	payload := &types.CreateUserRequest{}
	if err := utils.ParseJSON(r, payload); err != nil {
		return err
	}
	if err := types.Validate(payload); err != nil {
		return err
	}

	log.Printf("name: %s, email: %s", payload.Name, payload.Email)

	if err := h.storage.Store(r.Context(), payload); err != nil {
		return err
	}

	return nil
}
