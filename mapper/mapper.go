package mapper

import (
	"usedeall/model/entity"
	"usedeall/model/response"
)

type Mapper interface {
	UsersDTO(params []entity.User) []response.User
	UserDTO(params *entity.User) response.User
}

type MapperImpl struct {
}

func NewMapper() Mapper {
	return &MapperImpl{}
}

func (mapper *MapperImpl) UsersDTO(params []entity.User) []response.User {
	datas := make([]response.User, 0)
	for _, item := range params {
		data := response.User{
			Id:        item.Id,
			Username:  item.Username,
			Role:      item.Role,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		datas = append(datas, data)
	}
	return datas
}

func (mapper *MapperImpl) UserDTO(params *entity.User) response.User {
	data := response.User{
		Id:        params.Id,
		Username:  params.Username,
		Role:      params.Role,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}
	return data
}
