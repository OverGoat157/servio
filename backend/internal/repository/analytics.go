package repository

import "github.com/jmoiron/sqlx"

type AnalyticsRepo struct {
	db *sqlx.DB
}

func NewAnalyticsRepo(db *sqlx.DB) *AnalyticsRepo {
	return &AnalyticsRepo{db: db}
}

// IncrementView увеличивает счётчик просмотров на сегодня
func (r *AnalyticsRepo) IncrementView(restaurantID int64) {
	r.db.Exec(
		`INSERT INTO page_views (restaurant_id, view_date, views) VALUES ($1, CURRENT_DATE, 1)
		 ON CONFLICT (restaurant_id, view_date) DO UPDATE SET views = page_views.views + 1`,
		restaurantID,
	)
}

type DayStat struct {
	Date   string `json:"date" db:"date"`
	Views  int    `json:"views" db:"views"`
	Orders int    `json:"orders" db:"orders"`
	Revenue int   `json:"revenue" db:"revenue"`
}

type AnalyticsSummary struct {
	TotalViews     int       `json:"total_views"`
	TotalOrders    int       `json:"total_orders"`
	TotalRevenue   int       `json:"total_revenue"`
	UniqueClients  int       `json:"unique_clients"`
	AvgCheck       int       `json:"avg_check"`
	DayStats       []DayStat `json:"day_stats"`
}

type PlatformSummary struct {
	TotalUsers       int              `json:"total_users"`
	TotalRestaurants int              `json:"total_restaurants"`
	TotalViews       int              `json:"total_views"`
	TotalOrders      int              `json:"total_orders"`
	TotalRevenue     int              `json:"total_revenue"`
	UniqueClients    int              `json:"unique_clients"`
	DayStats         []DayStat        `json:"day_stats"`
	RestaurantStats  []RestaurantStat `json:"restaurant_stats"`
}

type RestaurantStat struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Owner    string `json:"owner" db:"owner"`
	Views    int    `json:"views" db:"views"`
	Orders   int    `json:"orders" db:"orders"`
	Revenue  int    `json:"revenue" db:"revenue"`
	Clients  int    `json:"clients" db:"clients"`
}

// GetPlatformSummary возвращает аналитику по всей платформе
func (r *AnalyticsRepo) GetPlatformSummary(days int) (*PlatformSummary, error) {
	s := &PlatformSummary{}

	r.db.Get(&s.TotalUsers, `SELECT COUNT(*) FROM users`)
	r.db.Get(&s.TotalRestaurants, `SELECT COUNT(*) FROM restaurants`)

	r.db.Get(&s.TotalViews,
		`SELECT COALESCE(SUM(views), 0) FROM page_views WHERE view_date >= CURRENT_DATE - $1::integer`, days)

	r.db.Get(&s.TotalOrders,
		`SELECT COUNT(*) FROM orders WHERE created_at >= CURRENT_DATE - $1::integer`, days)

	r.db.Get(&s.TotalRevenue,
		`SELECT COALESCE(SUM(total), 0) FROM orders WHERE created_at >= CURRENT_DATE - $1::integer`, days)

	r.db.Get(&s.UniqueClients,
		`SELECT COUNT(DISTINCT customer_phone) FROM orders WHERE customer_phone != '' AND created_at >= CURRENT_DATE - $1::integer`, days)

	// Статистика по дням
	r.db.Select(&s.DayStats,
		`SELECT
			d.date::text AS date,
			COALESCE(pv.views, 0) AS views,
			COALESCE(o.orders, 0) AS orders,
			COALESCE(o.revenue, 0) AS revenue
		FROM generate_series(CURRENT_DATE - $1::integer, CURRENT_DATE, '1 day') AS d(date)
		LEFT JOIN (
			SELECT view_date, SUM(views) AS views FROM page_views GROUP BY view_date
		) pv ON pv.view_date = d.date
		LEFT JOIN (
			SELECT created_at::date AS day, COUNT(*) AS orders, SUM(total) AS revenue
			FROM orders GROUP BY created_at::date
		) o ON o.day = d.date
		ORDER BY d.date`, days)

	// Статистика по ресторанам
	r.db.Select(&s.RestaurantStats,
		`SELECT
			r.id,
			r.name,
			COALESCE(u.name, u.email) AS owner,
			COALESCE(pv.views, 0) AS views,
			COALESCE(o.orders, 0) AS orders,
			COALESCE(o.revenue, 0) AS revenue,
			COALESCE(o.clients, 0) AS clients
		FROM restaurants r
		JOIN users u ON u.id = r.user_id
		LEFT JOIN (
			SELECT restaurant_id, SUM(views) AS views FROM page_views
			WHERE view_date >= CURRENT_DATE - $1::integer GROUP BY restaurant_id
		) pv ON pv.restaurant_id = r.id
		LEFT JOIN (
			SELECT restaurant_id, COUNT(*) AS orders, SUM(total) AS revenue,
				COUNT(DISTINCT NULLIF(customer_phone, '')) AS clients
			FROM orders WHERE created_at >= CURRENT_DATE - $1::integer
			GROUP BY restaurant_id
		) o ON o.restaurant_id = r.id
		ORDER BY COALESCE(o.orders, 0) DESC`, days)

	return s, nil
}

// GetSummary возвращает аналитику за последние N дней
func (r *AnalyticsRepo) GetSummary(restaurantID int64, days int) (*AnalyticsSummary, error) {
	s := &AnalyticsSummary{}

	// Просмотры за период
	r.db.Get(&s.TotalViews,
		`SELECT COALESCE(SUM(views), 0) FROM page_views
		 WHERE restaurant_id = $1 AND view_date >= CURRENT_DATE - $2::integer`,
		restaurantID, days)

	// Заказы, выручка, средний чек
	r.db.Get(&s.TotalOrders,
		`SELECT COUNT(*) FROM orders
		 WHERE restaurant_id = $1 AND created_at >= CURRENT_DATE - $2::integer`,
		restaurantID, days)

	r.db.Get(&s.TotalRevenue,
		`SELECT COALESCE(SUM(total), 0) FROM orders
		 WHERE restaurant_id = $1 AND created_at >= CURRENT_DATE - $2::integer`,
		restaurantID, days)

	if s.TotalOrders > 0 {
		s.AvgCheck = s.TotalRevenue / s.TotalOrders
	}

	// Уникальные клиенты (по телефону)
	r.db.Get(&s.UniqueClients,
		`SELECT COUNT(DISTINCT customer_phone) FROM orders
		 WHERE restaurant_id = $1 AND customer_phone != '' AND created_at >= CURRENT_DATE - $2::integer`,
		restaurantID, days)

	// Статистика по дням
	err := r.db.Select(&s.DayStats,
		`SELECT
			d.date::text AS date,
			COALESCE(pv.views, 0) AS views,
			COALESCE(o.orders, 0) AS orders,
			COALESCE(o.revenue, 0) AS revenue
		FROM generate_series(CURRENT_DATE - $2::integer, CURRENT_DATE, '1 day') AS d(date)
		LEFT JOIN page_views pv ON pv.restaurant_id = $1 AND pv.view_date = d.date
		LEFT JOIN (
			SELECT created_at::date AS day, COUNT(*) AS orders, SUM(total) AS revenue
			FROM orders WHERE restaurant_id = $1
			GROUP BY created_at::date
		) o ON o.day = d.date
		ORDER BY d.date`,
		restaurantID, days)

	return s, err
}
