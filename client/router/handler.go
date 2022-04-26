package router

import (
	"context"
	grpcclient "corpuser/client/grpcClient"
	pb "corpuser/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var c = grpcclient.Get()

type CustomError struct {
	Err    string `json:"message"`
	Status int    `json:"status"`
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	ids := r.URL.Query().Get("id")
	// no Id provided, then list all users
	uId := &pb.UsersRequest{Id: []int64{}}
	if ids != "" {
		idListSlice := strings.Split(strings.Trim(ids, ","), ",")
		for _, val := range idListSlice {
			if val, err := strconv.Atoi(val); err == nil {
				uId.Id = append(uId.Id, int64(val))
			} else {
				log.Printf("Invalid id provided(IGNORING): %v\n", val)
			}
		}
		// checking if it contains at least 1 valid id
		if len(uId.Id) == 0 {
			w.WriteHeader(400)
			log.Printf("No valid ids provided\n")
			err_msg := fmt.Sprintf("No valid ids provided: %v", ids)
			json.NewEncoder(w).Encode(CustomError{Status: 400, Err: err_msg})
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	respSlice, err := c.UserClient.GetUsersById(ctx, uId)
	if err != nil {
		log.Printf("Error occurred: %v\n", err)
		e := checkGrpcError(err)
		json.NewEncoder(w).Encode(e)
		return
	}
	var response []*pb.User
	response = append(response, respSlice.Users...)
	log.Printf("%v\n", response)
	json.NewEncoder(w).Encode(response)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userId int32
	uId := mux.Vars(r)["id"]
	if id, err := strconv.Atoi(uId); err != nil {
		log.Printf("Invalid User Id: %v", err)
		w.WriteHeader(400)
		e := CustomError{
			Status: 400,
			Err:    fmt.Sprintf("Invalid Id: %v", uId),
		}
		json.NewEncoder(w).Encode(e)
		return
	} else {
		userId = int32(id)
	}
	log.Printf("Requested user id: %d\n", userId)
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := c.UserClient.GetUser(ctx, &pb.UserRequest{Id: int64(userId)})
	if err != nil {
		log.Printf("Error occurred: %v", err)
		e := checkGrpcError(err)
		w.WriteHeader(e.Status)
		json.NewEncoder(w).Encode(e)
		return
	}
	log.Printf("User: %v", resp)
	json.NewEncoder(w).Encode(resp)
}

func checkGrpcError(err error) CustomError {

	er := CustomError{
		Status: 500,
	}

	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.InvalidArgument:
			er.Status = 400
			er.Err = e.Message()
		case codes.Internal:
			er.Err = e.Message()
		case codes.NotFound:
			er.Status = 400
			er.Err = e.Message()
		default:
			er.Err = e.Message()
		}
	} else {
		fmt.Printf("not able to parse error returned %v", err)
	}
	return er
}
