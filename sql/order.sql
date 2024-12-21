-- Create database for order service
CREATE DATABASE order_service;

-- Use the order_service database
USE order_service;

-- Table to store order information
CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Order ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    order_number VARCHAR(50) NOT NULL UNIQUE COMMENT 'Order number',
    total_amount DECIMAL(10, 2) NOT NULL COMMENT 'Total order amount',
    status ENUM('pending', 'processing', 'shipped', 'delivered', 'cancelled', 'refunded') DEFAULT 'pending' COMMENT 'Order status',
        address_id BIGINT COMMENT 'Address ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (user_id) REFERENCES user_service.users(id),
        FOREIGN KEY (address_id) REFERENCES user_service.user_addresses(id)
) COMMENT='Order information table';

-- Add indexes to optimize searching
CREATE INDEX idx_user_id ON orders(user_id);
CREATE INDEX idx_order_number ON orders(order_number);
CREATE INDEX idx_status ON orders(status);

-- Table to store order items
CREATE TABLE order_items (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Order item ID',
    order_id BIGINT NOT NULL COMMENT 'Order ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID',
    quantity INT NOT NULL COMMENT 'Quantity',
    price DECIMAL(10, 2) NOT NULL COMMENT 'Price per unit',
        seckill_id BIGINT COMMENT 'seckill_id',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
) COMMENT='Order items table';

-- Add indexes to optimize searching
CREATE INDEX idx_order_id ON order_items(order_id);
CREATE INDEX idx_product_id ON order_items(product_id);

-- Table to store order payment information (Simplified)
CREATE TABLE order_payments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Payment ID',
    order_id BIGINT NOT NULL COMMENT 'Order ID',
    payment_method ENUM('credit_card', 'debit_card', 'paypal', 'bank_transfer') NOT NULL COMMENT 'Payment method',
    payment_status ENUM('pending', 'completed', 'failed') DEFAULT 'pending' COMMENT 'Payment status',
    transaction_id VARCHAR(100) COMMENT 'Transaction ID',
    payment_amount DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
) COMMENT='Order payment information table';

-- Add indexes to optimize searching
CREATE INDEX idx_order_id_payment ON order_payments(order_id);
CREATE INDEX idx_payment_status ON order_payments(payment_status);
CREATE INDEX idx_transaction_id ON order_payments(transaction_id);