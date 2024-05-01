package payment

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-order-management/internal/service"
	"github.com/jfelipearaujo-org/ms-order-management/internal/service/order/get"
	"github.com/jfelipearaujo-org/ms-order-management/internal/service/payment/send_to_pay"
	"github.com/jfelipearaujo-org/ms-order-management/internal/shared/custom_error"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	sendToPayService service.SendToPayService[send_to_pay.SendToPayDto]
	getOrderService  service.GetOrderService[get.GetOrderDto]
}

func NewHandler(
	sendToPayService service.SendToPayService[send_to_pay.SendToPayDto],
	getOrderService service.GetOrderService[get.GetOrderDto],
) *Handler {
	return &Handler{
		sendToPayService: sendToPayService,
		getOrderService:  getOrderService,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	var request send_to_pay.SendToPayDto

	if err := ctx.Bind(&request); err != nil {
		return custom_error.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	context := ctx.Request().Context()

	getOrderRequest := get.GetOrderDto{
		OrderId: request.OrderID,
	}

	order, err := h.getOrderService.Handle(context, getOrderRequest)
	if err != nil {
		if custom_error.IsBusinessErr(err) {
			return custom_error.NewHttpAppErrorFromBusinessError(err)
		}

		return custom_error.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	if !order.HasItems() {
		return custom_error.NewHttpAppErrorFromBusinessError(custom_error.ErrOrderHasNoItems)
	}

	if order.HasOnGoingPayments() {
		return custom_error.NewHttpAppErrorFromBusinessError(custom_error.ErrOrderHasOnGoingPayments)
	}

	if err := h.sendToPayService.Handle(context, &order, request); err != nil {
		if custom_error.IsBusinessErr(err) {
			return custom_error.NewHttpAppErrorFromBusinessError(err)
		}

		return custom_error.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	ok := map[string]string{
		"message": "payment sent to be paid",
	}

	return ctx.JSON(http.StatusCreated, ok)
}