package adminhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	adminusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/admin"
)

type ServiceHandler interface {
	GetServiceList(ctx echo.Context) error
	GetServiceDetail(ctx echo.Context) error
	RegisterService(ctx echo.Context) error
	UpdateService(ctx echo.Context) error
	DeleteService(ctx echo.Context) error
}

type serviceHandler struct {
	usecase adminusecase.ServiceUsecase
}

func NewServiceHandler(usecase adminusecase.ServiceUsecase) ServiceHandler {
	return &serviceHandler{usecase: usecase}
}

func (h *serviceHandler) GetServiceList(ctx echo.Context) error {
	res, errRes := h.usecase.GetServiceList()
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *serviceHandler) GetServiceDetail(ctx echo.Context) error {
	serviceIDParam := ctx.Param("serviceId")
	serviceID, err := strconv.Atoi(serviceIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ServiceID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.GetServiceDetail(serviceID)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *serviceHandler) RegisterService(ctx echo.Context) error {
	body := new(admindto.RegisterServiceRequest)
	if err := ctx.Bind(body); err != nil {
		errRes := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	validator.New()
	if err := ctx.Validate(body); err != nil {
		errRes := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.RegisterService(body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *serviceHandler) UpdateService(ctx echo.Context) error {
	serviceIDParam := ctx.Param("serviceId")
	serviceID, err := strconv.Atoi(serviceIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ServiceID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	body := new(admindto.UpdateServiceRequest)
	if err := ctx.Bind(body); err != nil {
		errRes := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	validator.New()
	if err := ctx.Validate(body); err != nil {
		errRes := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.UpdateService(serviceID, body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *serviceHandler) DeleteService(ctx echo.Context) error {
	serviceIDParam := ctx.Param("serviceId")
	serviceID, err := strconv.Atoi(serviceIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ServiceID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.DeleteService(serviceID)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
