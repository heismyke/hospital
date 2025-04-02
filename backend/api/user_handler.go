package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/heismyke/backend_hospital/store"
	"github.com/heismyke/backend_hospital/utils"
)

type UserHandler struct{
	userStore store.UserStore	
	logger *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler{
	return &UserHandler{
		userStore: userStore,
		logger: logger,
	}
}

type signUpRequest struct{
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

//validate request first

func (h *UserHandler) validateSignUpRequest (req *signUpRequest) error {
	if req.Fullname == "" {
		return errors.New("fullname is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email){
		return errors.New("email is not valid")
	}

	return nil
}

func (h *UserHandler) HandleUserFunc(w http.ResponseWriter, r *http.Request){
	var req signUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		h.logger.Printf("Error: decoding signup: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "invalid payload request"})
		return
	}

	err = h.validateSignUpRequest(&req)
	if err != nil{
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": err.Error()})
		 return
	}

	user := &store.User{
		FullName: req.Fullname,
		Email: req.Email,
	}

	err = user.PasswordHash.Set(req.Password)
	if err != nil{
		h.logger.Printf("ERROR: PsswordHash %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}
	err = h.userStore.CreateUser(user)
	if err != nil{
		h.logger.Printf("ERROR: creating user %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"user": user})

}