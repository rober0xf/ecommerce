CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    category VARCHAR(100),
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT check_price_positive CHECK (price >= 0),
    CONSTRAINT check_stock_quantity_positive CHECK (stock_quantity >= 0)
);

-- index for faster queries
CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_products_status ON products(status);
