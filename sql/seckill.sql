-- Create database for seckill service
CREATE DATABASE seckill_service;

-- Use the seckill_service database
USE seckill_service;

-- Table to store seckill events
CREATE TABLE seckill_events (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Seckill Event ID',
    name VARCHAR(255) NOT NULL COMMENT 'Name of the seckill event',
    start_time TIMESTAMP NOT NULL COMMENT 'Start time of the seckill event',
    end_time TIMESTAMP NOT NULL COMMENT 'End time of the seckill event',
    status ENUM('upcoming', 'ongoing', 'ended') DEFAULT 'upcoming' COMMENT 'Status of the event',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Event creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time'
) COMMENT='Seckill event table';

-- Table to store products available in seckill events
CREATE TABLE seckill_products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Seckill Product ID',
    event_id BIGINT NOT NULL COMMENT 'Seckill Event ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID from the product service',
    seckill_price DECIMAL(10, 2) NOT NULL COMMENT 'Discounted price during the seckill event',
    stock INT NOT NULL COMMENT 'Available stock for the product in the event',
        original_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (event_id) REFERENCES seckill_events(id) ON DELETE CASCADE,
        FOREIGN KEY (product_id) REFERENCES product_service.products(id)
) COMMENT='Products participating in seckill events';

-- Add indexes to optimize searching
CREATE INDEX idx_event_id ON seckill_products(event_id);
CREATE INDEX idx_product_id ON seckill_products(product_id);

-- Table to record successful purchases
CREATE TABLE seckill_orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Seckill Order ID',
    event_id BIGINT NOT NULL COMMENT 'Seckill Event ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    quantity INT NOT NULL COMMENT 'Quantity purchased',
    order_status ENUM('pending', 'completed', 'failed') DEFAULT 'pending' COMMENT 'Order status',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Order creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
        FOREIGN KEY (user_id) REFERENCES user_service.users(id),
    FOREIGN KEY (event_id) REFERENCES seckill_events(id) ON DELETE CASCADE
) COMMENT='Seckill order records';

-- Add unique constraint to prevent duplicate purchases by the same user
CREATE UNIQUE INDEX idx_event_user_product ON seckill_orders(event_id, user_id, product_id);

-- Add indexes to optimize searching
CREATE INDEX idx_event_id ON seckill_orders(event_id);
CREATE INDEX idx_user_id ON seckill_orders(user_id);