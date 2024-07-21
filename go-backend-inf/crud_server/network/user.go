package network

import (
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
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
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
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("It's get")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("It's update")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("It's delete")
}
