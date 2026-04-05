package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"ab-production/internal/model"
)

// TelegramConfig — настройки для Telegram-бота
type TelegramConfig struct {
	BotToken string `json:"bot_token"`
	ChatID   string `json:"chat_id"`
}

// WhatsAppConfig — настройки для WhatsApp
type WhatsAppConfig struct {
	Phone string `json:"phone"` // номер в формате 79991234567
}

// OrderMessage — форматированное сообщение о заказе
type OrderMessage struct {
	RestaurantName string
	OrderID        int64
	Items          []model.OrderItem
	Total          int
	CustomerName   string
	CustomerPhone  string
	Messenger      string
}

// FormatOrderText создаёт текстовое представление заказа
func FormatOrderText(msg *OrderMessage) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("🍽 Новый заказ #%d\n", msg.OrderID))
	b.WriteString(fmt.Sprintf("Ресторан: %s\n\n", msg.RestaurantName))

	for _, item := range msg.Items {
		price := float64(item.Price*item.Quantity) / 100
		b.WriteString(fmt.Sprintf("• %s x%d — %.0f ₽\n", item.Name, item.Quantity, price))
	}

	total := float64(msg.Total) / 100
	b.WriteString(fmt.Sprintf("\n💰 Итого: %.0f ₽\n", total))

	if msg.CustomerName != "" {
		b.WriteString(fmt.Sprintf("👤 %s\n", msg.CustomerName))
	}
	if msg.CustomerPhone != "" {
		b.WriteString(fmt.Sprintf("📞 %s\n", msg.CustomerPhone))
	}

	b.WriteString(fmt.Sprintf("\nМессенджер: %s", msg.Messenger))

	return b.String()
}

// SendTelegram отправляет сообщение через Telegram Bot API
func SendTelegram(cfg TelegramConfig, text string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.BotToken)

	data := url.Values{
		"chat_id":    {cfg.ChatID},
		"text":       {text},
		"parse_mode": {"HTML"},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("telegram request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram error %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// BuildWhatsAppURL создаёт ссылку wa.me с текстом заказа
func BuildWhatsAppURL(cfg WhatsAppConfig, text string) string {
	phone := strings.ReplaceAll(cfg.Phone, "+", "")
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	phone = strings.ReplaceAll(phone, "-", "")

	return fmt.Sprintf("https://wa.me/%s?text=%s", phone, url.QueryEscape(text))
}

// ParseTelegramConfig парсит JSON конфиг Telegram
func ParseTelegramConfig(configJSON string) (TelegramConfig, error) {
	var cfg TelegramConfig
	err := json.Unmarshal([]byte(configJSON), &cfg)
	return cfg, err
}

// ParseWhatsAppConfig парсит JSON конфиг WhatsApp
func ParseWhatsAppConfig(configJSON string) (WhatsAppConfig, error) {
	var cfg WhatsAppConfig
	err := json.Unmarshal([]byte(configJSON), &cfg)
	return cfg, err
}
