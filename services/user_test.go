package services

import (
	"errors"
	"testing"

	mocks "github.com/rajajamal/golang-framework/mocks/repositories"
	"github.com/rajajamal/golang-framework/models"
	"github.com/rajajamal/golang-framework/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_User_Get_Paginated_Error(t *testing.T) {
	paginator := utils.Paginator{}
	repository := mocks.UserRepository{}

	repository.On("FindPaginated", paginator).Return(paginator, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.GetPaginated(paginator)

	repository.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, err.Code)
}

func Test_User_Get_Paginated(t *testing.T) {
	paginator := utils.Paginator{}
	repository := mocks.UserRepository{}

	repository.On("FindPaginated", paginator).Return(paginator, nil).Once()

	service := User{Repository: &repository}
	_, err := service.GetPaginated(paginator)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Validate_Login_User_Not_Found(t *testing.T) {
	form := models.Login{}
	form.Username = "admin"
	form.Password = "12345"

	repository := mocks.UserRepository{}
	repository.On("FindByUsername", form.Username).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.ValidateLogin(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusBadRequest)
}

func Test_User_Validate_Login_Password_Not_Match(t *testing.T) {
	form := models.Login{}
	form.Username = "admin"
	form.Password = "12345"

	user := models.User{
		Username: form.Username,
		Password: utils.EncodePassword("abcde"),
	}

	repository := mocks.UserRepository{}
	repository.On("FindByUsername", form.Username).Return(user, nil).Once()

	service := User{Repository: &repository}
	_, err := service.ValidateLogin(form)

	repository.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, err.Code, fiber.StatusBadRequest)
}

func Test_User_Validate_Login_Valid(t *testing.T) {
	form := models.Login{}
	form.Username = "admin"
	form.Password = "12345"

	user := models.User{
		Username: form.Username,
		Password: utils.EncodePassword(form.Password),
	}

	repository := mocks.UserRepository{}
	repository.On("FindByUsername", form.Username).Return(user, nil).Once()

	service := User{Repository: &repository}
	_, err := service.ValidateLogin(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Get_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", 1).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Get(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_User_Get_Found(t *testing.T) {
	user := models.User{
		ID:       1,
		Username: "admin",
	}

	repository := mocks.UserRepository{}
	repository.On("Find", user.ID).Return(user, nil).Once()

	service := User{Repository: &repository}
	result, err := service.Get(user.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, user.Username, result.Username)
}

func Test_User_Create_Error(t *testing.T) {
	form := models.CreateUser{
		Username: "user",
		Password: "12345",
	}

	repository := mocks.UserRepository{}
	repository.On("Save", mock.Anything).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}

func Test_User_Create_Success(t *testing.T) {
	form := models.CreateUser{
		Username: "user",
		Password: "12345",
	}

	repository := mocks.UserRepository{}
	repository.On("Save", mock.Anything).Return(nil).Once()

	service := User{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Delete_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", 1).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	err := service.Delete(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_User_Delete_Error(t *testing.T) {
	user := models.User{
		ID:       1,
		Username: "admin",
	}

	repository := mocks.UserRepository{}
	repository.On("Find", user.ID).Return(user, nil).Once()
	repository.On("Delete", &user).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	err := service.Delete(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}

func Test_User_Delete_Found(t *testing.T) {
	user := models.User{
		ID:       1,
		Username: "admin",
	}

	repository := mocks.UserRepository{}
	repository.On("Find", user.ID).Return(user, nil).Once()
	repository.On("Delete", &user).Return(nil).Once()

	service := User{Repository: &repository}
	err := service.Delete(user.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Update_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", 0).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Update(models.UpdateUser{})

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_User_Update_Error(t *testing.T) {
	form := models.UpdateUser{
		ID:       1,
		Username: "admin",
	}
	user := models.User{
		ID:       form.ID,
		Username: form.Username,
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", &user).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Update(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}

func Test_User_Update_Success(t *testing.T) {
	form := models.UpdateUser{
		ID:       1,
		Username: "admin",
	}
	user := models.User{
		ID:       form.ID,
		Username: form.Username,
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", &user).Return(nil).Once()

	service := User{Repository: &repository}
	result, err := service.Update(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)
}

func Test_User_Update_Password_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", 0).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(models.UpdatePassword{})

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_User_Update_Password_Old_Password_Not_Match(t *testing.T) {
	form := models.UpdatePassword{
		ID:          1,
		OldPassword: "old",
		Password:    "new",
	}
	user := models.User{
		ID:       form.ID,
		Username: "admin",
		Password: form.Password,
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusBadRequest)
}

func Test_User_Update_Password_Error(t *testing.T) {
	form := models.UpdatePassword{
		ID:          1,
		OldPassword: "old",
		Password:    "new",
	}
	user := models.User{
		ID:       form.ID,
		Username: "admin",
		Password: utils.EncodePassword(form.OldPassword),
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", mock.Anything).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}

func Test_User_Update_Password_Success(t *testing.T) {
	form := models.UpdatePassword{
		ID:          1,
		OldPassword: "old",
		Password:    "new",
	}
	user := models.User{
		ID:       form.ID,
		Username: "admin",
		Password: utils.EncodePassword(form.OldPassword),
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", mock.Anything).Return(nil).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}
