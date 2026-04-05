# AB Team

Платформа для создания шаблонных мини-приложений с цифровым меню для малого ресторанного бизнеса. Клиенты сканируют QR-код, просматривают меню и отправляют заказ прямо в мессенджер ресторана (Telegram, WhatsApp).

## Архитектура

```
┌─────────────────┐     ┌─────────────────┐     ┌──────────────┐
│  Admin Panel     │────▶│   Go API        │────▶│  PostgreSQL  │
│  (Vue, :3001)    │     │   (Gin, :8080)  │     │  (:5432)     │
└─────────────────┘     └────────┬────────┘     └──────────────┘
                                 │
┌─────────────────┐              │          ┌───────────────────┐
│  Public Menu     │─────────────┘          │  Telegram Bot API │
│  (Vue, :3000)    │                        │  WhatsApp (wa.me) │
└─────────────────┘                        └───────────────────┘

┌─────────────────┐
│  Landing Page    │  — маркетинговый сайт
│  (Vue, :5173)    │
└─────────────────┘
```

## Структура проекта

```
restaurants/
├── backend/                    # Go API сервер
│   ├── cmd/server/main.go      # Точка входа
│   ├── internal/
│   │   ├── config/             # Конфигурация (.env)
│   │   ├── handler/            # HTTP handlers
│   │   ├── middleware/         # JWT auth, CORS
│   │   ├── model/              # Структуры данных
│   │   ├── repository/         # Работа с PostgreSQL
│   │   └── service/            # Бизнес-логика (мессенджеры)
│   ├── migrations/             # SQL миграции
│   └── Dockerfile
├── frontend/
│   ├── menu/                   # Публичное меню для клиентов
│   │   ├── src/views/          # HomePage, MenuPage, CartPage
│   │   ├── src/stores/         # restaurant, cart
│   │   └── Dockerfile
│   └── admin/                  # Панель управления
│       ├── src/views/          # Login, Dashboard, Menu, Orders, Messengers
│       ├── src/stores/         # auth
│       └── Dockerfile
├── landing/                    # Лендинг (маркетинг)
│   └── src/components/         # Hero, Features, Pricing, Demo, CTA, Footer
└── docker-compose.postgres.yml
```

## Технологии

| Компонент | Стек |
|-----------|------|
| Backend | Go, Gin, sqlx, PostgreSQL, JWT |
| Frontend | Vue 3, Vite, Vue Router |
| Мессенджеры | Telegram Bot API, WhatsApp (wa.me) |
| Деплой | Docker, nginx |

## Быстрый старт

### 1. Запуск PostgreSQL

```bash
docker compose -f docker-compose.postgres.yml up -d postgres
```

### 2. Запуск бэкенда

```bash
cd backend
cp .env.example .env  # или отредактируйте .env
go run ./cmd/server/
```

Сервер запустится на `http://localhost:8080`. Миграции применяются автоматически.

### 3. Запуск публичного меню

```bash
cd frontend/menu
npm install
npm run dev
```

Доступно на `http://localhost:5173`. Откройте `http://localhost:5173/{slug-ресторана}`.

### 4. Запуск admin-панели

```bash
cd frontend/admin
npm install
npm run dev
```

Доступно на `http://localhost:5174`.

### 5. Запуск лендинга

```bash
cd landing
npm install
npm run dev
```

## Docker

Каждый сервис имеет свой Dockerfile и запускается независимо.

```bash
# Только PostgreSQL
docker compose -f docker-compose.postgres.yml up postgres

# Только бэкенд
docker compose -f docker-compose.postgres.yml up backend

# Только публичное меню
docker compose -f docker-compose.postgres.yml up menu

# Только admin-панель
docker compose -f docker-compose.postgres.yml up admin

# Всё вместе
docker compose -f docker-compose.postgres.yml up
```

### Порты

| Сервис | Порт | Описание |
|--------|------|----------|
| PostgreSQL | 5432 | База данных |
| Backend | 8080 | REST API |
| Menu | 3000 | Публичное меню (nginx) |
| Admin | 3001 | Панель управления (nginx) |

### Переменные окружения

Создайте `.env` в корне проекта для docker compose:

```env
JWT_SECRET=your-secure-secret-key
VITE_API_URL=http://your-server:8080
```

Переменные бэкенда (`backend/.env`):

| Переменная | По умолчанию | Описание |
|------------|-------------|----------|
| `DB_HOST` | localhost | Хост PostgreSQL |
| `DB_PORT` | 5432 | Порт PostgreSQL |
| `DB_USER` | abteam | Пользователь БД |
| `DB_PASSWORD` | abteam | Пароль БД |
| `DB_NAME` | abteam | Имя базы данных |
| `DB_SSLMODE` | disable | SSL режим |
| `JWT_SECRET` | change-me | Секрет для JWT токенов |
| `SERVER_PORT` | 8080 | Порт API сервера |

## API

### Публичное API (без авторизации)

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/public/menu/:slug` | Полное меню ресторана |
| POST | `/api/public/menu/:slug/order` | Создать заказ |

### Авторизация

| Метод | Путь | Описание |
|-------|------|----------|
| POST | `/api/auth/register` | Регистрация |
| POST | `/api/auth/login` | Вход |

### Защищённое API (требует JWT)

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/me` | Текущий пользователь |
| POST | `/api/restaurants` | Создать ресторан |
| GET | `/api/restaurants` | Список ресторанов |
| GET | `/api/restaurants/:id` | Получить ресторан |
| PUT | `/api/restaurants/:id` | Обновить ресторан |
| DELETE | `/api/restaurants/:id` | Удалить ресторан |
| POST | `/api/restaurants/:id/categories` | Создать категорию |
| GET | `/api/restaurants/:id/categories` | Список категорий |
| PUT | `/api/categories/:id` | Обновить категорию |
| DELETE | `/api/categories/:id` | Удалить категорию |
| POST | `/api/categories/:id/items` | Создать позицию меню |
| GET | `/api/categories/:id/items` | Список позиций |
| PUT | `/api/items/:id` | Обновить позицию |
| DELETE | `/api/items/:id` | Удалить позицию |
| GET | `/api/restaurants/:id/orders` | Список заказов |
| PATCH | `/api/orders/:id/status` | Изменить статус заказа |
| GET | `/api/restaurants/:id/messengers` | Настройки мессенджеров |
| POST | `/api/restaurants/:id/messengers` | Создать/обновить мессенджер |
| DELETE | `/api/restaurants/:id/messengers/:type` | Удалить мессенджер |

## Интеграция с мессенджерами

### Telegram

1. Создайте бота через [@BotFather](https://t.me/BotFather)
2. Добавьте бота в группу или чат
3. В admin-панели укажите `bot_token` и `chat_id`
4. Заказы будут приходить автоматически от бота

### WhatsApp

1. В admin-панели укажите номер WhatsApp (формат: `79991234567`)
2. При заказе клиент переходит по ссылке `wa.me` с готовым текстом
3. Клиент нажимает "Отправить" в WhatsApp

### Формат сообщения о заказе

```
🍽 Новый заказ #42
Ресторан: Ricco

• Том Ям x1 — 890 ₽
• Стейк Рибай x1 — 1 490 ₽

💰 Итого: 2 380 ₽
👤 Иван
📞 +7 999 123-45-67

Мессенджер: telegram
```

## Схема базы данных

```
users
  ├── id, email, password_hash, name, created_at

restaurants
  ├── id, user_id, name, slug (unique), description, logo,
  │   phone, address, working_hours (JSONB), theme, created_at

categories
  ├── id, restaurant_id, name, sort_order

menu_items
  ├── id, category_id, name, description, price (копейки),
  │   image, available, sort_order

messenger_configs
  ├── id, restaurant_id, type, config (JSONB), enabled
  │   UNIQUE(restaurant_id, type)

orders
  ├── id, restaurant_id, items (JSONB), total (копейки),
  │   messenger, customer_name, customer_phone, status, created_at
```

## Пользовательский путь

### Владелец ресторана

1. Регистрация в admin-панели
2. Создание ресторана (название, slug, описание, часы работы)
3. Добавление категорий меню (Супы, Горячее, Десерты)
4. Добавление позиций с ценами
5. Настройка мессенджеров (Telegram-бот / WhatsApp-номер)
6. Печать QR-кода со ссылкой `menu-url/{slug}`

### Клиент ресторана

1. Сканирует QR-код или переходит по ссылке
2. Видит главную страницу ресторана (название, часы работы)
3. Переходит в меню, выбирает блюда
4. Оформляет заказ, выбирает мессенджер
5. Telegram — заказ приходит автоматически
6. WhatsApp — открывается чат с готовым текстом заказа
