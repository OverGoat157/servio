CREATE UNIQUE INDEX IF NOT EXISTS idx_menu_items_cat_name ON menu_items(category_id, name);
CREATE UNIQUE INDEX IF NOT EXISTS idx_combos_rest_name ON combos(restaurant_id, name);
CREATE UNIQUE INDEX IF NOT EXISTS idx_categories_rest_name ON categories(restaurant_id, name);
