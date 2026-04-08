ALTER TABLE menu_items
    ADD COLUMN ingredients TEXT,
    ADD COLUMN weight VARCHAR(50),
    ADD COLUMN calories INT,
    ADD COLUMN proteins NUMERIC(5,1),
    ADD COLUMN fats NUMERIC(5,1),
    ADD COLUMN carbs NUMERIC(5,1),
    ADD COLUMN cook_time VARCHAR(50);
