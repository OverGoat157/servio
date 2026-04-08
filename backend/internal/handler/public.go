package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"ab-team/internal/model"
	"ab-team/internal/repository"
	"ab-team/internal/service"
)

type PublicHandler struct {
	restaurants *repository.RestaurantRepo
	categories  *repository.CategoryRepo
	items       *repository.MenuItemRepo
	orders      *repository.OrderRepo
	messengers  *repository.MessengerRepo
	combos      *repository.ComboRepo
}

func NewPublicHandler(
	restaurants *repository.RestaurantRepo,
	categories *repository.CategoryRepo,
	items *repository.MenuItemRepo,
	orders *repository.OrderRepo,
	messengers *repository.MessengerRepo,
	combos *repository.ComboRepo,
) *PublicHandler {
	return &PublicHandler{
		restaurants: restaurants,
		categories:  categories,
		items:       items,
		orders:      orders,
		messengers:  messengers,
		combos:      combos,
	}
}

// MenuCategory — категория с позициями для публичного API
type MenuCategory struct {
	ID    int64            `json:"id"`
	Name  string           `json:"name"`
	Items []PublicMenuItem `json:"items"`
}

type PublicMenuItem struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Price       int      `json:"price"`
	Image       *string  `json:"image"`
	Available   bool     `json:"available"`
	Ingredients *string  `json:"ingredients"`
	Weight      *string  `json:"weight"`
	Calories    *int     `json:"calories"`
	Proteins    *float64 `json:"proteins"`
	Fats        *float64 `json:"fats"`
	Carbs       *float64 `json:"carbs"`
	CookTime    *string  `json:"cook_time"`
}

type PublicRestaurant struct {
	ID               int64           `json:"id"`
	Name             string          `json:"name"`
	Slug             string          `json:"slug"`
	Description      *string         `json:"description"`
	Logo             *string         `json:"logo"`
	CoverImage       *string         `json:"cover_image"`
	Phone            *string         `json:"phone"`
	Address          *string         `json:"address"`
	WorkingHours     json.RawMessage `json:"working_hours"`
	SocialLinks      json.RawMessage `json:"social_links"`
	Theme            string          `json:"theme"`
	PromoTitle       *string         `json:"promo_title"`
	PromoDescription *string         `json:"promo_description"`
	Categories       []MenuCategory  `json:"categories"`
	Combos           []PublicCombo   `json:"combos"`
	Messengers       []string        `json:"messengers"`
}

type PublicCombo struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	Image       *string          `json:"image"`
	Price       int              `json:"price"`
	Available   bool             `json:"available"`
	Items       []PublicComboItem `json:"items"`
}

type PublicComboItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
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
			pubItems = append(pubItems, PublicMenuItem{
				ID:          item.ID,
				Name:        item.Name,
				Description: item.Description,
				Price:       item.Price,
				Image:       item.Image,
				Available:   item.Available,
				Ingredients: item.Ingredients,
				Weight:      item.Weight,
				Calories:    item.Calories,
				Proteins:    item.Proteins,
				Fats:        item.Fats,
				Carbs:       item.Carbs,
				CookTime:    item.CookTime,
			})
		}
		menuCats = append(menuCats, MenuCategory{
			ID:    cat.ID,
			Name:  cat.Name,
			Items: pubItems,
		})
	}

	// Получаем комбо-наборы
	comboList, _ := h.combos.ListByRestaurant(rest.ID)
	var pubCombos []PublicCombo
	for _, combo := range comboList {
		items, _ := h.combos.GetItems(combo.ID)
		var pubItems []PublicComboItem
		for _, ci := range items {
			pubItems = append(pubItems, PublicComboItem{
				Name:     ci.Name,
				Quantity: ci.Quantity,
			})
		}
		if pubItems == nil {
			pubItems = []PublicComboItem{}
		}
		pubCombos = append(pubCombos, PublicCombo{
			ID:          combo.ID,
			Name:        combo.Name,
			Description: combo.Description,
			Image:       combo.Image,
			Price:       combo.Price,
			Available:   combo.Available,
			Items:       pubItems,
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
		ID:               rest.ID,
		Name:             rest.Name,
		Slug:             rest.Slug,
		Description:      rest.Description,
		Logo:             rest.Logo,
		CoverImage:       rest.CoverImage,
		Phone:            rest.Phone,
		Address:          rest.Address,
		Theme:            rest.Theme,
		PromoTitle:       rest.PromoTitle,
		PromoDescription: rest.PromoDescription,
		Categories:       menuCats,
		Combos:           pubCombos,
		Messengers:       messengers,
	}

	if rest.WorkingHours != nil {
		pub.WorkingHours = json.RawMessage(*rest.WorkingHours)
	}
	if rest.SocialLinks != nil {
		pub.SocialLinks = json.RawMessage(*rest.SocialLinks)
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
