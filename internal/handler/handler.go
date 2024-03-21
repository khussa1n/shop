package handler

import (
	"fmt"
	"github.com/khussa1n/shop/internal/service"
	"os"
	"strings"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) InitHandler() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Никаких аргументов не было представлено.")
		return
	}

	orders := strings.Split(args[0], ",")

	fmt.Println()
	fmt.Println("=+=+=+=")
	fmt.Println("Страница сборки заказов :", orders)

	goodsByOrders, err := h.service.Good.GetAllByOrders(orders...)
	if err != nil {
		fmt.Println("Ошибка при получении товаров по заказам:", err)
		return
	}

	for key, value := range goodsByOrders {
		fmt.Println()
		fmt.Println("===Стеллаж", key)
		for _, goodWithOrders := range value {
			fmt.Println()
			fmt.Printf("%s (id=%d)\n", goodWithOrders.Good.Name, goodWithOrders.Good.ID)
			fmt.Printf("заказ %d, %d шт\n", goodWithOrders.OrderNumber, goodWithOrders.GoodsCount)
			if goodWithOrders.AdditionalShelves != "{NULL}" {
				fmt.Println("доп стеллаж:", goodWithOrders.AdditionalShelves)
			}
		}
	}

}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
