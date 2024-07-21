package network

import (
	"crud_server/service"
	"crud_server/types"
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
	var req types.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("binding error", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("create error", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("It's success", 1, nil),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("It's success", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("binding error", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("update error", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("It's success", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("binding error", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("delete error", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("It's success", 1, nil),
		})
	}
}
