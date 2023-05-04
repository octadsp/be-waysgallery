package handlers

import (
	"net/http"
	"strconv"
	orderdto "waysgallery/dto/order"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	repository "waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerOrder struct {
	OrderRepository repository.OrderRepository
}

func HandlerOrder(OrderRepository repository.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) FindOrders(c echo.Context) error {
	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: orders})
}

func (h *handlerOrder) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: order})
}

func (h *handlerOrder) AddOrder(c echo.Context) error {
	request := new(orderdto.OrderRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	price, _ := strconv.Atoi(c.FormValue("price"))
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	order_to_id, _ := strconv.Atoi(c.Param("order_to_id"))

	order := models.Order{
		Title:       request.Title,
		Description: request.Description,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Price:       price,
		UserID:      user_id,
		OrderToID:   order_to_id,
	}

	data, err := h.OrderRepository.AddOrder(order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerOrder) UpdateOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(orderdto.OrderRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	order, err := h.OrderRepository.GetOrder(id)
	StartDate := request.StartDate
	EndDate := request.EndDate
	formatStartDate := StartDate.Format("02/01/2006")
	formatEndDate := EndDate.Format("02/01/2006")

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		order.Title = request.Title
	}

	if request.Description != "" {
		order.Description = request.Description
	}

	if formatStartDate != "" {
		order.StartDate = request.StartDate
	}

	if formatEndDate != "" {
		order.UserID = request.UserID
	}

	if request.Price != 0 {
		order.Price = request.Price
	}

	data, err := h.OrderRepository.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerOrder) DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.OrderRepository.DeleteOrder(order, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
