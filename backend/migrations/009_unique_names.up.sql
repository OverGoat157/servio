-- Удаляем дубликаты блюд (оставляем запись с наименьшим id)
DELETE FROM menu_items a USING menu_items b
WHERE a.category_id = b.category_id AND a.name = b.name AND a.id > b.id;

-- Удаляем дубликаты комбо
DELETE FROM combos a USING combos b
WHERE a.restaurant_id = b.restaurant_id AND a.name = b.name AND a.id > b.id;

-- Удаляем дубликаты категорий
DELETE FROM categories a USING categories b
WHERE a.restaurant_id = b.restaurant_id AND a.name = b.name AND a.id > b.id;

CREATE UNIQUE INDEX IF NOT EXISTS idx_menu_items_cat_name ON menu_items(category_id, name);
CREATE UNIQUE INDEX IF NOT EXISTS idx_combos_rest_name ON combos(restaurant_id, name);
CREATE UNIQUE INDEX IF NOT EXISTS idx_categories_rest_name ON categories(restaurant_id, name);
