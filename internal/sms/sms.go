package sms

import (
	"encoding/json"
	"errors"
	"golang-training/main/internal/repository"
	"net/http"
	"strconv"
)

type Service interface {
	GetData(w http.ResponseWriter, r *http.Request)
	AddData(w http.ResponseWriter, r *http.Request)
}

//implementer เ
//can keep data 
type service struct {
	Data []string
	userDetailRepo repository.UserDetail
}


func NewService(userDetailRepo repository.UserDetail) Service {
	return &service{
		Data: []string{},
		userDetailRepo:  userDetailRepo,
		
	}
}

//implement 
type GetDataResponse struct{
	UserDetail []repository.UserDetailEntity `json: "userDetail"`
	

}

func (s service) GetData(w http.ResponseWriter, r *http.Request){
	// var tmp []string
	// for _, val := range s.Data{
	// 	tmp = append(tmp, val)
	// }

	data, err := s.userDetailRepo.FindAllData()
	if err != nil {
		panic(err)
	}

	w.Header().Add("content-type", "application/json")
	_  = json.NewEncoder(w).Encode(GetDataResponse{
		UserDetail: data,
	})
}

type AddDataRequest struct {
	FirstName 	string 	`json:"firstName"`
	LastName 	string 	`json: "lastname"`
	Email 		string 	`json: "email"`
	Age 		int64 	`json: "age"`
	IsNew		bool	`json: "isNew"`
}

type AddDataRespond struct {
	Result 	string `json: "status"`
	Error 	string `json: "error"`
}

func (s service) AddData(w http.ResponseWriter, r *http.Request){
	//call request 
	req, err := validateReq(r)

	s.Data = append(s.Data, req.FirstName,req.LastName, req.Email, strconv.FormatInt(req.Age, 10)) //add data to array

	if req.IsNew == false {
		err = s.userDetailRepo.Update((repository.UserDetailEntity{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Age: req.Age,
			Email: req.Email,
		}))
		if err != nil {
			http.Error(w, "There was a problem updating..", http.StatusInternalServerError) 
			return	
		}
	}
	if req.IsNew == true  {
		err = s.userDetailRepo.Insert(repository.UserDetailEntity{
			FirstName: req.FirstName,
			LastName: req.LastName,		
			Email: req.Email,
			Age: req.Age,
		})

		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return 
		}
	}

	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(&AddDataRespond{
		Result: "Sucessfully",
		Error: "nil",
	})
	 
}

func validateReq(r *http.Request) (*AddDataRequest, error){
	var req AddDataRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("invalid request")
	}

	if len(req.FirstName)  == 0 {
		return nil, errors.New("firstname is required")
	}

	return &req, nil
}