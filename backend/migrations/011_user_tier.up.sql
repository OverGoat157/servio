ALTER TABLE users ADD COLUMN IF NOT EXISTS tier VARCHAR(20) NOT NULL DEFAULT 'basic';
UPDATE users SET tier = 'basic' WHERE tier IS NULL OR tier = '';
