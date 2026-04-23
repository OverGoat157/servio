ALTER TABLE restaurants  DROP COLUMN IF EXISTS name_en;
ALTER TABLE restaurants  DROP COLUMN IF EXISTS description_en;

ALTER TABLE categories   DROP COLUMN IF EXISTS name_en;

ALTER TABLE menu_items   DROP COLUMN IF EXISTS name_en;
ALTER TABLE menu_items   DROP COLUMN IF EXISTS description_en;
ALTER TABLE menu_items   DROP COLUMN IF EXISTS ingredients_en;

ALTER TABLE combos       DROP COLUMN IF EXISTS name_en;
ALTER TABLE combos       DROP COLUMN IF EXISTS description_en;

ALTER TABLE combo_items  DROP COLUMN IF EXISTS name_en;
