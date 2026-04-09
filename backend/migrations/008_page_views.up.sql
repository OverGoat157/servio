CREATE TABLE IF NOT EXISTS page_views (
    id SERIAL PRIMARY KEY,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    view_date DATE NOT NULL DEFAULT CURRENT_DATE,
    views INTEGER NOT NULL DEFAULT 0,
    UNIQUE(restaurant_id, view_date)
);

CREATE INDEX IF NOT EXISTS idx_page_views_rest_date ON page_views(restaurant_id, view_date);
