package main

import (
	"sync"
	"sync/atomic"

	"github.com/go-openapi/errors"

	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
)

// User manager
type UserManager struct {
	users []*pb.User
}
var usersLock = &sync.Mutex{}
var lastUserId int64 = 2

//CreateUsersInDB: add test users
func (userManager *UserManager) CreateUsersInDB() (error) {
 
	user1 := &pb.User{
		FirstName:    	              	"name_user1",
		LastName:  	                     "lastname_user1",
		EmailAddress:  "user1@gmail.com",
	}
	admin1 := &pb.User{
		FirstName:    	              	"name_admin1",
		LastName:  	                     "lastname_admin1",
		EmailAddress:  "admin1@gmail.com",
	}

	userManager.addUser(user1)
	userManager.addUser(admin1)
	return nil
}

func newUserId() int64 {
	return atomic.AddInt64(&lastUserId, 1)
}


// returns a new User manager
func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userManager *UserManager)  addUser(user *pb.User) {

	usersLock.Lock()
	defer usersLock.Unlock()
	user.UserId = newUserId()
	userManager.users = append(userManager.users, user)
}

// return user by Id
func (userManager *UserManager) userById(Id int64) (*pb.User, error) {

	for _, user := range userManager.users {
		if user.UserId == Id {
			return user, nil
		}
	}
	return nil, errors.NotFound("not found: user %d", Id)
}