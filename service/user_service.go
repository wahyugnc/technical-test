package service

import (
	"database/sql"
	"usedeall/exception"
	"usedeall/mapper"
	"usedeall/model/entity"
	"usedeall/model/request"
	"usedeall/model/response"
	"usedeall/repository"
	"usedeall/utils"
	"usedeall/validation"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll(ctx *fiber.Ctx, pageNumber int, pageSize int) (*response.UserDataResponse, error)
	FindById(ctx *fiber.Ctx, userId int) (*response.User, error)
	Save(ctx *fiber.Ctx, params *entity.User) (*response.User, error)
	Update(ctx *fiber.Ctx, params *entity.User) (*response.User, error)
	Delete(ctx *fiber.Ctx, userId int) (*response.GeneralResponse, error)
	UserLogin(ctx *fiber.Ctx, params *request.LoginRequest) (*response.LoginResponse, error)
	RefreshToken(ctx *fiber.Ctx, params *request.RefreshTokenRequest) (*response.RefreshTokenResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Mapper         mapper.Mapper
}

func NewUserService(userRepository repository.UserRepository, mapper mapper.Mapper) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Mapper:         mapper,
	}
}

func (service *UserServiceImpl) FindAll(ctx *fiber.Ctx, pageNumber int, pageSize int) (*response.UserDataResponse, error) {
	claims := utils.JwtVerificationWithClaim(ctx)

	if claims.Role != utils.Assign_Role {
		panic(exception.ForbiddenError{
			Message: utils.Message_Forbidden,
		})
	}

	var number int
	if pageNumber == 1 {
		number = 0
	} else {
		number = pageNumber - 1
	}

	res, err := service.UserRepository.FindAll(ctx.Context(), number, pageSize)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	count, err := service.UserRepository.Count(ctx.Context())
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	dto := service.Mapper.UsersDTO(res)

	result := &response.UserDataResponse{
		Data:  dto,
		Total: count,
	}

	return result, nil
}

func (service *UserServiceImpl) Save(ctx *fiber.Ctx, params *entity.User) (*response.User, error) {
	claims := utils.JwtVerificationWithClaim(ctx)

	if claims.Role != utils.Assign_Role {
		panic(exception.ForbiddenError{
			Message: utils.Message_Forbidden,
		})
	}

	validation.UserSaveValidate(ctx, *params)

	usernameRequest := &request.LoginRequest{
		Username: params.Username,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	userRequest := &entity.User{
		Username:  params.Username,
		Password:  string(hash),
		Role:      params.Role,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	user, err := service.UserRepository.FindByUsername(ctx.Context(), usernameRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			userId, err := service.UserRepository.Save(ctx.Context(), userRequest)
			if err != nil {
				panic(exception.GeneralError{
					Message: err.Error(),
				})
			}

			res, err := service.UserRepository.FindById(ctx.Context(), userId)
			if err != nil {
				panic(exception.DataNotFoundError{
					Message: utils.Message_Data_Notfound,
				})
			}

			dto := service.Mapper.UserDTO(res)

			return &dto, nil
		} else {
			panic(exception.GeneralError{
				Message: err.Error(),
			})
		}
	}

	if user.Username != "" {
		panic(exception.DataExsistError{
			Message: utils.Message_Duplicate_Username,
		})
	}

	userId, err := service.UserRepository.Save(ctx.Context(), userRequest)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	res, err := service.UserRepository.FindById(ctx.Context(), userId)
	if err != nil {
		panic(exception.DataNotFoundError{
			Message: utils.Message_Data_Notfound,
		})
	}

	dto := service.Mapper.UserDTO(res)

	return &dto, nil
}

func (service *UserServiceImpl) Update(ctx *fiber.Ctx, params *entity.User) (*response.User, error) {
	claims := utils.JwtVerificationWithClaim(ctx)

	if claims.Role != utils.Assign_Role {
		panic(exception.ForbiddenError{
			Message: utils.Message_Forbidden,
		})
	}

	validation.UserUpdateValidate(ctx, *params)

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	userRequest := &entity.User{
		Username:  params.Username,
		Password:  string(hash),
		Role:      params.Role,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	_, err = service.UserRepository.FindById(ctx.Context(), params.Id)
	if err != nil {
		if err != sql.ErrNoRows {
			err = service.UserRepository.Update(ctx.Context(), userRequest)
			if err != nil {
				panic(exception.GeneralError{
					Message: err.Error(),
				})
			}

			res, err := service.UserRepository.FindById(ctx.Context(), params.Id)
			if err != nil {
				panic(exception.GeneralError{
					Message: err.Error(),
				})
			}

			dto := service.Mapper.UserDTO(res)

			return &dto, nil
		} else {
			panic(exception.DataNotFoundError{
				Message: utils.Message_Data_Notfound,
			})
		}
	}

	err = service.UserRepository.Update(ctx.Context(), userRequest)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	res, err := service.UserRepository.FindById(ctx.Context(), params.Id)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	dto := service.Mapper.UserDTO(res)

	return &dto, nil
}

func (service *UserServiceImpl) FindById(ctx *fiber.Ctx, userId int) (*response.User, error) {
	claims := utils.JwtVerificationWithClaim(ctx)

	if claims.UserId != userId {
		if claims.Role != utils.Assign_Role {
			panic(exception.ForbiddenError{
				Message: utils.Message_Forbidden,
			})
		}
	}

	res, err := service.UserRepository.FindById(ctx.Context(), userId)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(exception.DataNotFoundError{
				Message: utils.Message_Data_Notfound,
			})
		} else {
			panic(exception.GeneralError{
				Message: err.Error(),
			})
		}
	}

	dto := service.Mapper.UserDTO(res)
	return &dto, nil
}

func (service *UserServiceImpl) Delete(ctx *fiber.Ctx, userId int) (*response.GeneralResponse, error) {
	claims := utils.JwtVerificationWithClaim(ctx)

	if claims.Role != utils.Assign_Role {
		panic(exception.ForbiddenError{
			Message: utils.Message_Forbidden,
		})
	}

	_, err := service.UserRepository.FindById(ctx.Context(), userId)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(exception.DataNotFoundError{
				Message: utils.Message_Data_Notfound,
			})
		} else {
			panic(exception.GeneralError{
				Message: err.Error(),
			})
		}
	}

	err = service.UserRepository.Delete(ctx.Context(), userId)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	result := &response.GeneralResponse{
		Message:    utils.Message_Success,
		StatusCode: fiber.StatusOK,
	}

	return result, nil
}

func (service *UserServiceImpl) UserLogin(ctx *fiber.Ctx, params *request.LoginRequest) (*response.LoginResponse, error) {
	validation.UserLoginValidate(ctx, *params)

	res, err := service.UserRepository.FindByUsername(ctx.Context(), params)
	if err != nil {
		panic(exception.GeneralError{
			Message: utils.Message_Invalid_Login,
		})
	}

	byteHash := []byte(res.Password)

	err = bcrypt.CompareHashAndPassword(byteHash, []byte(params.Password))
	if err != nil {
		panic(exception.GeneralError{
			Message: utils.Message_Invalid_Login,
		})
	}

	accessToken, err := utils.GenerateNewAccessToken(res.Username, res.Role, res.Id)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	refreshToken, err := utils.GenerateNewRefreshToken(res.Username, res.Role, res.Id)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	result := &response.LoginResponse{
		StatusCode:   fiber.StatusOK,
		Message:      utils.Message_Success,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return result, nil
}

func (service *UserServiceImpl) RefreshToken(ctx *fiber.Ctx, params *request.RefreshTokenRequest) (*response.RefreshTokenResponse, error) {
	validation.RefreshTokenValidate(ctx, *params)

	claims := utils.JwtVerificationWithClaimEmail(ctx, params.RefreshToken)
	params.Username = claims.Username

	accessToken, err := utils.GenerateNewAccessToken(claims.Username, claims.Role, claims.UserId)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	result := &response.RefreshTokenResponse{
		StatusCode:  fiber.StatusOK,
		Message:     utils.Message_Success,
		AccessToken: accessToken,
	}

	return result, nil
}
