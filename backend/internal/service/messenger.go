package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"ab-team/internal/model"
)

var reDigits = regexp.MustCompile(`\d+`)

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
	RestaurantName  string
	OrderID         int64
	Items           []model.OrderItem
	Total           int
	CustomerName    string
	CustomerPhone   string
	CustomerAddress string
	Comment         string
	MenuURL         string
	DesiredTime     string // "asap" или "HH:MM"
	EstCookMin      int    // примерное время готовки в минутах
}

// FormatOrderText создаёт текстовое представление заказа
func FormatOrderText(msg *OrderMessage) string {
	var b strings.Builder

	now := time.Now().Format("02.01.2006 15:04")

	b.WriteString("🔔🔔🔔 НОВЫЙ ЗАКАЗ 🔔🔔🔔\n")
	b.WriteString("━━━━━━━━━━━━━━━━━━━━\n")
	b.WriteString(fmt.Sprintf("🆔 Заказ #%d\n", msg.OrderID))
	b.WriteString(fmt.Sprintf("🏪 %s\n", msg.RestaurantName))
	b.WriteString(fmt.Sprintf("🕐 %s\n", now))
	b.WriteString("━━━━━━━━━━━━━━━━━━━━\n\n")

	b.WriteString("📋 СОСТАВ ЗАКАЗА:\n")
	b.WriteString("──────────────────\n")
	var totalQty int
	for i, item := range msg.Items {
		price := float64(item.Price*item.Quantity) / 100
		b.WriteString(fmt.Sprintf("%d. %s × %d — %.0f ₽\n", i+1, item.Name, item.Quantity, price))
		if item.Comment != "" {
			b.WriteString(fmt.Sprintf("   💬 %s\n", item.Comment))
		}
		totalQty += item.Quantity
	}

	total := float64(msg.Total) / 100
	b.WriteString("──────────────────\n")
	b.WriteString(fmt.Sprintf("📦 Позиций: %d (блюд: %d)\n", len(msg.Items), totalQty))
	b.WriteString(fmt.Sprintf("💰 ИТОГО: %.0f ₽\n", total))

	if msg.EstCookMin > 0 {
		b.WriteString(fmt.Sprintf("⏱ Примерное время готовки: ~%d мин\n", msg.EstCookMin))
	}

	if msg.DesiredTime != "" && msg.DesiredTime != "asap" {
		b.WriteString(fmt.Sprintf("\n⏰ ПРИГОТОВИТЬ К: %s\n", msg.DesiredTime))
	} else {
		b.WriteString("\n⚡ КАК МОЖНО БЫСТРЕЕ\n")
	}

	if msg.CustomerName != "" || msg.CustomerPhone != "" || msg.CustomerAddress != "" {
		b.WriteString("\n👤 КЛИЕНТ:\n")
		if msg.CustomerName != "" {
			b.WriteString(fmt.Sprintf("   Имя: %s\n", msg.CustomerName))
		}
		if msg.CustomerPhone != "" {
			b.WriteString(fmt.Sprintf("   Тел: %s\n", msg.CustomerPhone))
		}
		if msg.CustomerAddress != "" {
			b.WriteString(fmt.Sprintf("   📍 Адрес: %s\n", msg.CustomerAddress))
		}
	}

	if msg.Comment != "" {
		b.WriteString(fmt.Sprintf("\n📝 КОММЕНТАРИЙ:\n   %s\n", msg.Comment))
	}

	b.WriteString("\n━━━━━━━━━━━━━━━━━━━━\n")

	if msg.MenuURL != "" {
		b.WriteString(fmt.Sprintf("🔗 Меню: %s\n", msg.MenuURL))
	}

	b.WriteString("⚡ Обработайте заказ как можно скорее!")

	return b.String()
}

// SendTelegram отправляет сообщение через Telegram Bot API
func SendTelegram(cfg TelegramConfig, text string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.BotToken)

	data := url.Values{
		"chat_id": {cfg.ChatID},
		"text":    {text},
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

// FormatOrderTextPlain создаёт текст заказа без эмодзи (для WhatsApp)
func FormatOrderTextPlain(msg *OrderMessage) string {
	var b strings.Builder

	now := time.Now().Format("02.01.2006 15:04")

	b.WriteString("*** НОВЫЙ ЗАКАЗ ***\n")
	b.WriteString("--------------------\n")
	b.WriteString(fmt.Sprintf("Заказ #%d\n", msg.OrderID))
	b.WriteString(fmt.Sprintf("Ресторан: %s\n", msg.RestaurantName))
	b.WriteString(fmt.Sprintf("Время: %s\n", now))
	b.WriteString("--------------------\n\n")

	b.WriteString("СОСТАВ ЗАКАЗА:\n")
	var totalQty int
	for i, item := range msg.Items {
		price := float64(item.Price*item.Quantity) / 100
		b.WriteString(fmt.Sprintf("%d. %s x%d — %.0f руб.\n", i+1, item.Name, item.Quantity, price))
		if item.Comment != "" {
			b.WriteString(fmt.Sprintf("   > %s\n", item.Comment))
		}
		totalQty += item.Quantity
	}

	total := float64(msg.Total) / 100
	b.WriteString(fmt.Sprintf("\nПозиций: %d (блюд: %d)\n", len(msg.Items), totalQty))
	b.WriteString(fmt.Sprintf("ИТОГО: %.0f руб.\n", total))

	if msg.EstCookMin > 0 {
		b.WriteString(fmt.Sprintf("Время готовки: ~%d мин\n", msg.EstCookMin))
	}

	if msg.DesiredTime != "" && msg.DesiredTime != "asap" {
		b.WriteString(fmt.Sprintf("\nПРИГОТОВИТЬ К: %s\n", msg.DesiredTime))
	} else {
		b.WriteString("\nКАК МОЖНО БЫСТРЕЕ\n")
	}

	if msg.CustomerName != "" || msg.CustomerPhone != "" || msg.CustomerAddress != "" {
		b.WriteString("\nКЛИЕНТ:\n")
		if msg.CustomerName != "" {
			b.WriteString(fmt.Sprintf("  Имя: %s\n", msg.CustomerName))
		}
		if msg.CustomerPhone != "" {
			b.WriteString(fmt.Sprintf("  Тел: %s\n", msg.CustomerPhone))
		}
		if msg.CustomerAddress != "" {
			b.WriteString(fmt.Sprintf("  Адрес: %s\n", msg.CustomerAddress))
		}
	}

	if msg.Comment != "" {
		b.WriteString(fmt.Sprintf("\nКОММЕНТАРИЙ:\n  %s\n", msg.Comment))
	}

	if msg.MenuURL != "" {
		b.WriteString(fmt.Sprintf("\nМеню: %s\n", msg.MenuURL))
	}

	return b.String()
}

// BuildWhatsAppURL создаёт ссылку web.whatsapp.com с текстом заказа,
// минуя промежуточный лендинг wa.me/api.whatsapp.com, который на десктопе
// пытается открыть нативное приложение через whatsapp:// протокол.
func BuildWhatsAppURL(cfg WhatsAppConfig, text string) string {
	phone := strings.ReplaceAll(cfg.Phone, "+", "")
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	phone = strings.ReplaceAll(phone, "-", "")

	return fmt.Sprintf("https://web.whatsapp.com/send?phone=%s&text=%s", phone, url.QueryEscape(text))
}

// ParseCookTimeMinutes извлекает число минут из строки вроде "15 мин", "30", "1 час"
func ParseCookTimeMinutes(s string) int {
	if s == "" {
		return 0
	}
	m := reDigits.FindString(s)
	if m == "" {
		return 0
	}
	n, _ := strconv.Atoi(m)
	if strings.Contains(strings.ToLower(s), "час") {
		n *= 60
	}
	return n
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
