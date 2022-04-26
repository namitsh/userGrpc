package server

import (
	"context"
	pb "corpuser/user"
	"log"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// it contains the services of the grpc Server.. it interacts with the

type UserService struct {
	pb.UnimplementedUserMethodServer
	Users []*pb.User
}

// temporary database
func Seed() []*pb.User {

	var s []*pb.User

	var i int64 = 1

	for ; i <= 10; i++ {
		usr := &pb.User{
			Id:      i,
			Fname:   "Hello" + strconv.FormatInt(i, 10),
			City:    "Dubai" + strconv.FormatInt(i, 10),
			Phone:   123456789,
			Height:  5.8,
			Married: true,
		}
		s = append(s, usr)
	}
	return s
}

// create a new instance of UserService
func NewService() *UserService {
	s := Seed()
	return &UserService{
		Users: s,
	}
}

//

func (s *UserService) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.User, error) {
	log.Println("Received request: Get User By Id")
	log.Println("Checking the validation of request")
	// validation of request
	if in == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument\n")
	}
	if in.Id == 0 {
		log.Println("Id is not provided.")
		return nil, status.Errorf(codes.InvalidArgument, "User Id is not provided\n")
	}
	log.Printf("Successful validation for user ID: %d", in.Id)
	for _, val := range s.Users {
		if val.Id == in.Id {
			log.Printf("User with id: %d found with data: %v \n", in.Id, val)
			return val, nil
		}
	}
	log.Printf("User with Id: %d  not found\n", in.Id)
	return nil, status.Errorf(codes.NotFound, "User Id: `%d` not found.\n", in.Id)
}

func (s *UserService) GetUsersById(ctx context.Context, in *pb.UsersRequest) (*pb.UserResponse, error) {
	log.Println("Received request: Get Users based on Ids")
	log.Println("Checking the validation of request")
	if in == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument\n")
	}
	// no id provided in the list , return all the users.
	if len(in.Id) == 0 {
		log.Printf("%v\n", s.Users)
		return &pb.UserResponse{Users: s.Users}, nil
	}

	log.Printf("Successful validation for user IDs: %d", in.Id)

	var res []*pb.User

	// TODO : using nested loop tempororily, try to come up with good waay, if any

	for _, val := range s.Users {
		for _, id := range in.Id {
			if val.Id == id {
				res = append(res, val)
				break
			}
		}
	}
	log.Printf("%v\n", res)
	return &pb.UserResponse{Users: res}, nil
}
