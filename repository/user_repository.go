package repository

import (
	"context"
	"usedeall/model/entity"
	"usedeall/model/request"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, params *request.LoginRequest) (*entity.User, error)
	FindById(ctx context.Context, id int) (*entity.User, error)
	Save(ctx context.Context, params *entity.User) (int, error)
	Update(ctx context.Context, params *entity.User) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) ([]entity.User, error)
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, id int) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, params *request.LoginRequest) (*entity.User, error) {
	var data entity.User
	row := repository.DB.WithContext(ctx).Raw(`select id, username, passwd, role, first_name, last_name, nvl(created_at,'') created_at, nvl(updated_at,'') updated_at from users where username=?`, params.Username).Row()
	err := row.Scan(
		&data.Id,
		&data.Username,
		&data.Password,
		&data.Role,
		&data.FirstName,
		&data.LastName,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, id int) (*entity.User, error) {
	var data entity.User
	row := repository.DB.WithContext(ctx).Raw(`select id, username, passwd, role, first_name, last_name, nvl(created_at,'') created_at, nvl(updated_at,'') updated_at from users where id=?`, id).Row()
	err := row.Scan(
		&data.Id,
		&data.Username,
		&data.Password,
		&data.Role,
		&data.FirstName,
		&data.LastName,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, params *entity.User) (int, error) {
	var Id int
	err := repository.DB.WithContext(ctx).Raw(`insert into users(username, passwd, role, first_name, last_name, created_at) values(?, ?, ?, ?, ?, now()) returning id`, params.Username, params.Password, params.Role, params.FirstName, params.LastName).Scan(&Id).Error
	if err != nil {
		return 0, err
	}

	return Id, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, params *entity.User) error {
	err := repository.DB.WithContext(ctx).Raw(`update users set username=?, passwd=?, role=?, first_name=?, last_name=?, updated_at=now() where id=?`, params.Username, params.Password, params.Role, params.FirstName, params.LastName, params.Id).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) ([]entity.User, error) {
	rows, err := repository.DB.WithContext(ctx).Raw(`select id, username, passwd, role, first_name, last_name, nvl(created_at,'') created_at, nvl(updated_at,'') updated_at from users limit ?, ?`, pageNumber, pageSize).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	datas := make([]entity.User, 0)
	for rows.Next() {
		data := entity.User{}
		err := rows.Scan(
			&data.Id,
			&data.Username,
			&data.Password,
			&data.Role,
			&data.FirstName,
			&data.LastName,
			&data.CreatedAt,
			&data.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (repository *UserRepositoryImpl) Count(ctx context.Context) (int, error) {
	var count int
	err := repository.DB.WithContext(ctx).Raw(`select count(1) from users`).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, id int) error {
	err := repository.DB.WithContext(ctx).Exec(`delete from users where id = ?`, id).Error
	if err != nil {
		return err
	}

	return nil
}
