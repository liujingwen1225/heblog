package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"heblog/internal/heblog/biz"
	"heblog/internal/heblog/store"
	"heblog/internal/pkg/core"
	"heblog/internal/pkg/errno"
	"heblog/internal/pkg/log"
	v1 "heblog/pkg/api/heblog/v1"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}

func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")
	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
