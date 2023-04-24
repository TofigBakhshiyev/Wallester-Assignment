package service

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"wallester.com/assignment/models"
)

type Customer struct {
	Firstname *string `json:"firstname"`

	Lastname *string `json:"lastname"`

	Birthdate *string `json:"birthdate"`

	Gender *string `json:"gender"`

	Email *string `json:"email"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateCustomer(context *fiber.Ctx) error {

	customer := Customer{}

	err := context.BodyParser(&customer)

	if err != nil {

		context.Status(http.StatusUnprocessableEntity).JSON(

			&fiber.Map{"message": "request failed"})

		return err

	}

	validator := validator.New()

	err = validator.Struct(Customer{})

	if err != nil {

		context.Status(http.StatusUnprocessableEntity).JSON(

			&fiber.Map{"message": err},
		)

		return err

	}

	err = r.DB.Create(&customer).Error

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not create customer"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "customer has been successfully added",
	})

	return nil

}

func (r *Repository) UpdateCustomer(context *fiber.Ctx) error {

	id := context.Params("id")

	if id == "" {

		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{

			"message": "id cannot be empty",
		})

		return nil

	}

	customerModel := &models.Customers{}

	customer := Customer{}

	err := context.BodyParser(&customer)

	if err != nil {

		context.Status(http.StatusUnprocessableEntity).JSON(

			&fiber.Map{"message": "request failed"})

		return err

	}

	err = r.DB.Model(customerModel).Where("id = ?", id).Updates(customer).Error

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(&fiber.Map{

			"message": "could not update customer",
		})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "customer has been successfully updated",
	})

	return nil

}

func (r *Repository) GetCustomer(context *fiber.Ctx) error {

	customerModels := &[]models.Customers{}

	err := r.DB.Find(customerModels).Error

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not get customers"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "customers gotten successfully",

		"data": customerModels,
	})

	return nil

}

func (r *Repository) GetCustomerByID(context *fiber.Ctx) error {

	name := context.Params("name")

	customerModel := &models.Customers{}

	if name == "" {

		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{

			"message": "name cannot be empty",
		})

		return nil

	}

	err := r.DB.Where("firstname = ? or lastname = ?", name, name).First(customerModel).Error

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not get customer"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "customer id gotten successfully",

		"data": customerModel,
	})

	return nil

}

func (r *Repository) SetupRoutes(app *fiber.App) {

	api := app.Group("/wallester")

	api.Post("/create_customer", r.CreateCustomer)

	api.Patch("/update_customer/:id", r.UpdateCustomer)

	api.Get("/get_customer/:name", r.GetCustomerByID)

	api.Get("/customer", r.GetCustomer)

}
