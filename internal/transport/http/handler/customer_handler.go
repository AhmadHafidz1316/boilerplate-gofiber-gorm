package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/service"
	"github.com/AhmadHafidz1316/go-fiber-gorm/pkg/response"
)

type CustomerHandler struct {
	svc service.CustomerService
}

func NewCustomerHandler(svc service.CustomerService) *CustomerHandler {
	return &CustomerHandler{svc: svc}
}

// List Customers
// @Summary Get list of customers
// @Description Retrieve all customers from the database
// @Tags customers
// @Produce json
// @Success 200 {array} domain.Customer
// @Failure 500 {object} response.Envelope
// @Router /customers [get]
func (h *CustomerHandler) List(c *fiber.Ctx) error {
	items, err := h.svc.List(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}
	return response.Success(c, fiber.StatusOK, items)
}

// Get Customer
// @Summary Get customer by ID
// @Description Retrieve a single customer by its ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} domain.Customer
// @Failure 404 {object} response.Envelope
// @Failure 500 {object} response.Envelope
// @Router /customers/{id} [get]
func (h *CustomerHandler) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	item, err := h.svc.Get(c.Context(), uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}
	if item == nil {
		return response.Error(c, fiber.StatusNotFound, "NOT_FOUND", "customer not found")
	}
	return response.Success(c, fiber.StatusOK, item)
}

// Create Customer
// @Summary Create a new customer
// @Description Add a new customer to the database
// @Tags customers
// @Accept json
// @Produce json
// @Param data body service.CreateCustomerInput true "Customer data"
// @Success 201 {object} domain.Customer
// @Failure 400 {object} response.Envelope
// @Router /customers [post]
func (h *CustomerHandler) Create(c *fiber.Ctx) error {
	var in service.CreateCustomerInput
	if err := c.BodyParser(&in); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "INVALID_BODY", err.Error())
	}
	item, err := h.svc.Create(c.Context(), in)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, item)
}

// Update Customer
// @Summary Update customer by ID
// @Description Update an existing customer by its ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param data body service.UpdateCustomerInput true "Updated data"
// @Success 200 {object} domain.Customer
// @Failure 400 {object} response.Envelope
// @Failure 404 {object} response.Envelope
// @Router /customers/{id} [put]
func (h *CustomerHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var in service.UpdateCustomerInput
	if err := c.BodyParser(&in); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "INVALID_BODY", err.Error())
	}
	item, err := h.svc.Update(c.Context(), uint(id), in)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", err.Error())
	}
	if item == nil {
		return response.Error(c, fiber.StatusNotFound, "NOT_FOUND", "customer not found")
	}
	return response.Success(c, fiber.StatusOK, item)
}

// Delete Customer
// @Summary Delete customer by ID
// @Description Remove a customer from the database by its ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 204 "No Content"
// @Failure 500 {object} response.Envelope
// @Router /customers/{id} [delete]
func (h *CustomerHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.svc.Delete(c.Context(), uint(id)); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}
	return response.Success(c, fiber.StatusNoContent, nil)
}
