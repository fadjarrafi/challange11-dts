package service

import (
	"challange10-dts/entity"
	"challange10-dts/pkg/errs"
	"challange10-dts/pkg/helpers"
	"challange10-dts/repository/product_repository"
	"challange10-dts/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authService struct {
	userRepo    user_repository.UserRepository
	productRepo product_repository.ProductRepository
}

func NewAuthService(userRepo user_repository.UserRepository, productRepo product_repository.ProductRepository) AuthService {
	return &authService{
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (a *authService) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		productId, err := helpers.GetParamId(ctx, "productId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		product, err := a.productRepo.GetProductById(productId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if user.Level == entity.Admin {
			ctx.Next()
			return
		}

		if product.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the product data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
