-- Create database for cart service
CREATE DATABASE cart_service;

-- Use the cart_service database
USE cart_service;

-- Table to store cart items
CREATE TABLE cart_items (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Cart item ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID',
    quantity INT NOT NULL COMMENT 'Quantity of the product in the cart',
    price DECIMAL(10, 2) NOT NULL COMMENT 'Price of the product at the time it was added to the cart',
    selected BOOLEAN DEFAULT TRUE COMMENT 'Whether the item is selected for checkout',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Time when the item was added to the cart',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    UNIQUE KEY unique_user_product (user_id, product_id) COMMENT 'Ensure one product per user in the cart',
    FOREIGN KEY (user_id) REFERENCES user_service.users(id),
    FOREIGN KEY (product_id) REFERENCES product_service.products(id)
) COMMENT='Shopping cart items table';

-- Add indexes to optimize searching
CREATE INDEX idx_user_id ON cart_items(user_id);
CREATE INDEX idx_product_id ON cart_items(product_id);
CREATE INDEX idx_selected ON cart_items(selected);