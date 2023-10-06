package controllers

import (
	models "myBE/app/models/users"
	"myBE/config"
	"myBE/helper"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c echo.Context) error {

	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(500, map[string]string{
			"message": "Failed to retrieve users",
		})
	}

	return c.JSON(200, users)
}

func RegisterUser(c echo.Context) (err error) {

	u := new(models.RegisterUserInput)

	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user := models.User{
		User_uuid: uuid.Must(uuid.NewRandom()).String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  string(passwordHash),
	}

	_, err = models.Save(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func Login(c echo.Context) (err error) {
	u := new(models.LoginInput)
	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	user, err := models.FindByEmail(u.Email)

	if err != nil {
		return echo.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		response := helper.APIResponse("Unauthorized", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// Create token with claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.User_uuid
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(time.Hour * 2))

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	formatter := models.FormatUser(user, t)

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	return c.JSON(http.StatusOK, response)

}

func GetProfile(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	uuid := claims.Uuid
	userdata, err := models.FindByID(uuid)

	if err != nil {
		return err
	}

	if userdata.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "No user found on with that ID")
	}

	formatter := models.FormatProfile(userdata)

	response := helper.APIResponse("Success get profile", http.StatusOK, "success", formatter)

	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) (err error) {
	id := c.Param("id")
	user, err := models.FindByID(id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "No user found on with that ID")
	}

	u := new(models.UpdateUserInput)

	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user.Name = u.Name
	user.Email = u.Email

	_, err = models.Update(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)

}
