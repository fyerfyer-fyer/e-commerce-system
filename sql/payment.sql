-- Create database for payment service
CREATE DATABASE payment_service;

-- Use the payment_service database
USE payment_service;

-- Table to store payment transactions
CREATE TABLE payments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Payment ID',
    order_id BIGINT NOT NULL COMMENT 'Order ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    amount DECIMAL(10, 2) NOT NULL COMMENT 'Payment amount',
    payment_method ENUM('credit_card', 'debit_card', 'paypal', 'bank_transfer', 'wallet') NOT NULL COMMENT 'Payment method',
    payment_status ENUM('pending', 'completed', 'failed', 'refunded') DEFAULT 'pending' COMMENT 'Payment status',
    transaction_id VARCHAR(100) COMMENT 'Transaction ID from the payment gateway',
    payment_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Time of payment',
        expired_at TIMESTAMP COMMENT 'payment expired time, which used for payment timeout',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
        FOREIGN KEY (order_id) REFERENCES order_service.orders(id),
    FOREIGN KEY (user_id) REFERENCES user_service.users(id)
) COMMENT='Payment transactions table';

-- Add indexes to optimize searching
CREATE INDEX idx_order_id ON payments(order_id);
CREATE INDEX idx_user_id ON payments(user_id);
CREATE INDEX idx_payment_status ON payments(payment_status);
CREATE INDEX idx_transaction_id ON payments(transaction_id);

-- Table to store refund information (optional, can be moved to a separate Refund service later if needed)
CREATE TABLE refunds (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Refund ID',
    payment_id BIGINT NOT NULL COMMENT 'Payment ID',
    refund_amount DECIMAL(10, 2) NOT NULL COMMENT 'Refund amount',
    refund_status ENUM('pending', 'processed', 'failed') DEFAULT 'pending' COMMENT 'Refund status',
    refund_reason VARCHAR(255) COMMENT 'Reason for refund',
    refund_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Time of refund',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (payment_id) REFERENCES payments(id) ON DELETE CASCADE
) COMMENT='Refund transactions table';

-- Add indexes to optimize searching
CREATE INDEX idx_payment_id ON refunds(payment_id);
CREATE INDEX idx_refund_status ON refunds(refund_status);