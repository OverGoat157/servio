package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"menulink/internal/model"
	"menulink/internal/repository"
	"menulink/internal/service"
)

type PublicHandler struct {
	restaurants *repository.RestaurantRepo
	categories  *repository.CategoryRepo
	items       *repository.MenuItemRepo
	orders      *repository.OrderRepo
	messengers  *repository.MessengerRepo
}

func NewPublicHandler(
	restaurants *repository.RestaurantRepo,
	categories *repository.CategoryRepo,
	items *repository.MenuItemRepo,
	orders *repository.OrderRepo,
	messengers *repository.MessengerRepo,
) *PublicHandler {
	return &PublicHandler{
		restaurants: restaurants,
		categories:  categories,
		items:       items,
		orders:      orders,
		messengers:  messengers,
	}
}

// MenuCategory — категория с позициями для публичного API
type MenuCategory struct {
	ID    int64            `json:"id"`
	Name  string           `json:"name"`
	Items []PublicMenuItem `json:"items"`
}

type PublicMenuItem struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
	Image       *string `json:"image"`
}

type PublicRestaurant struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Description  *string        `json:"description"`
	Logo         *string        `json:"logo"`
	Phone        *string        `json:"phone"`
	Address      *string        `json:"address"`
	WorkingHours json.RawMessage `json:"working_hours"`
	Theme        string         `json:"theme"`
	Categories   []MenuCategory `json:"categories"`
	Messengers   []string       `json:"messengers"`
}

// GetMenu возвращает полное меню ресторана по slug
func (h *PublicHandler) GetMenu(c *gin.Context) {
	slug := c.Param("slug")

	rest, err := h.restaurants.GetBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	cats, err := h.categories.ListByRestaurant(rest.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load categories"})
		return
	}

	var menuCats []MenuCategory
	for _, cat := range cats {
		items, err := h.items.ListByCategory(cat.ID)
		if err != nil {
			continue
		}
		var pubItems []PublicMenuItem
		for _, item := range items {
			if !item.Available {
				continue
			}
			pubItems = append(pubItems, PublicMenuItem{
				ID:          item.ID,
				Name:        item.Name,
				Description: item.Description,
				Price:       item.Price,
				Image:       item.Image,
			})
		}
		menuCats = append(menuCats, MenuCategory{
			ID:    cat.ID,
			Name:  cat.Name,
			Items: pubItems,
		})
	}

	// Получаем доступные мессенджеры
	msgConfigs, _ := h.messengers.ListByRestaurant(rest.ID)
	var messengers []string
	for _, mc := range msgConfigs {
		if mc.Enabled {
			messengers = append(messengers, mc.Type)
		}
	}

	pub := PublicRestaurant{
		ID:          rest.ID,
		Name:        rest.Name,
		Slug:        rest.Slug,
		Description: rest.Description,
		Logo:        rest.Logo,
		Phone:       rest.Phone,
		Address:     rest.Address,
		Theme:       rest.Theme,
		Categories:  menuCats,
		Messengers:  messengers,
	}

	if rest.WorkingHours != nil {
		pub.WorkingHours = json.RawMessage(*rest.WorkingHours)
	}

	c.JSON(http.StatusOK, pub)
}

// CreateOrder создаёт заказ от клиента
func (h *PublicHandler) CreateOrder(c *gin.Context) {
	slug := c.Param("slug")

	rest, err := h.restaurants.GetBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Подсчитываем итого и обогащаем items
	var total int
	var enrichedItems []model.OrderItem
	for _, oi := range req.Items {
		item, err := h.items.GetByID(oi.MenuItemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid menu item"})
			return
		}
		enrichedItems = append(enrichedItems, model.OrderItem{
			MenuItemID: item.ID,
			Name:       item.Name,
			Price:      item.Price,
			Quantity:   oi.Quantity,
		})
		total += item.Price * oi.Quantity
	}

	itemsJSON, _ := json.Marshal(enrichedItems)

	order, err := h.orders.Create(rest.ID, string(itemsJSON), total, req.Messenger, req.CustomerName, req.CustomerPhone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	// Формируем сообщение о заказе
	msg := &service.OrderMessage{
		RestaurantName: rest.Name,
		OrderID:        order.ID,
		Items:          enrichedItems,
		Total:          total,
		CustomerName:   req.CustomerName,
		CustomerPhone:  req.CustomerPhone,
		Messenger:      req.Messenger,
	}
	orderText := service.FormatOrderText(msg)

	response := gin.H{
		"order": order,
	}

	// Обрабатываем в зависимости от мессенджера
	switch req.Messenger {
	case "telegram":
		// Отправляем в Telegram от имени бота
		msgCfg, err := h.messengers.GetEnabled(rest.ID, "telegram")
		if err == nil {
			tgCfg, err := service.ParseTelegramConfig(msgCfg.Config)
			if err == nil {
				go func() {
					if err := service.SendTelegram(tgCfg, orderText); err != nil {
						log.Printf("telegram send error: %v", err)
					}
				}()
			}
		}
		response["messenger_sent"] = true

	case "whatsapp":
		// Генерируем ссылку wa.me
		msgCfg, err := h.messengers.GetEnabled(rest.ID, "whatsapp")
		if err == nil {
			waCfg, err := service.ParseWhatsAppConfig(msgCfg.Config)
			if err == nil {
				response["whatsapp_url"] = service.BuildWhatsAppURL(waCfg, orderText)
			}
		}
	}

	c.JSON(http.StatusCreated, response)
}
