package usecase

import (
	"database/sql"
	"errors"
	"gofiber/db/repository"
	"gofiber/db/repository/models"
	"gofiber/pkg/functioncaller"
	"gofiber/pkg/hashing"
	"gofiber/pkg/logruslogger"
	"gofiber/pkg/messages"
	"gofiber/server/requests"
	"gofiber/usecase/viewmodel"
	"time"
)

type UserUseCase struct {
	*UcContract
}

//browse
func (uc UserUseCase) Browse(search, orderBy, sort string, page, limit int) (res []viewmodel.UserVm, pagination viewmodel.PaginationVm, err error) {
	repo := repository.NewUserRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort, models.OrderBy)
	users, count, err := repo.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-browse")
		return res, pagination, err
	}

	for _, user := range users {
		res = append(res, uc.buildBody(user))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, nil
}

//read
func (uc UserUseCase) Read(ID string) (res viewmodel.UserVm, err error) {
	repo := repository.NewUserRepository(uc.DB)
	user, err := repo.Read(ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-read")
		return res, err
	}
	res = uc.buildBody(user)

	return res, nil
}

//edit
func (uc UserUseCase) Edit(input *requests.UserRequest, ID string) (res string, err error) {
	repo := repository.NewUserRepository(uc.DB)
	now := time.Now().UTC()
	var password string

	count, err := uc.countByEmail(ID, input.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmailAlreadyExist")
		return res, err
	}

	if input.Password != "" {
		password, _ = hashing.HashAndSalt(input.Password)
	}
	model := models.User{
		ID:        ID,
		Email:     input.Email,
		Password:  password,
		UpdatedAt: now,
	}
	res, err = repo.Edit(model)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-edit")
		return res, err
	}

	return res, nil
}

//add
func (uc UserUseCase) Add(input *requests.UserRequest) (res string, err error) {
	repo := repository.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	count, err := uc.countByEmail("", input.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-user-countByemailAlreadyExist")
		return res, errors.New(messages.DataAlreadyExist)
	}

	password, _ := hashing.HashAndSalt(input.Password)
	model := models.User{
		Email:     input.Email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
	res, err = repo.Add(model)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-add")
		return res, err
	}

	return res, nil
}

//delete
func(uc UserUseCase) Delete(ID string) (res string,err error){
	repo := repository.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	model := models.User{
		ID: ID,
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: now},
	}
	res,err = repo.Delete(model)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-delete")
		return res, err
	}

	return res,nil
}

//count by email
func (uc UserUseCase) countByEmail(ID, email string) (res int, err error) {
	repo := repository.NewUserRepository(uc.DB)
	res, err = repo.CountByEmail(ID, email)
	if err != nil {
		return res, err
	}

	return res, nil
}

//build body
func (uc UserUseCase) buildBody(model models.User) viewmodel.UserVm {
	return viewmodel.UserVm{
		ID:        model.ID,
		Email:     model.Email,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
	}
}
