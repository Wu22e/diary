package network

import (
	"crud_server/service"
	"crud_server/types"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once // only one callable
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerPOST("/", userRouterInstance.create)
		router.registerPUT("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
		router.registerGET("/", userRouterInstance.get)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("It's create")

	u.userService.Create(nil)

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("It's success", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("It's get")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("It's success", 1),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("It's update")

	u.userService.Update(nil, nil)

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("It's success", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("It's delete")

	u.userService.Delete(nil)

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("It's success", 1),
	})
}
