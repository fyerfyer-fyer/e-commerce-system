-- Create database for comment service
CREATE DATABASE comment_service;

-- Use the comment_service database
USE comment_service;

-- Table to store comments on products
CREATE TABLE comments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Comment ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID',
    user_id BIGINT NOT NULL COMMENT 'User ID of the commenter',
    rating TINYINT CHECK (rating BETWEEN 1 AND 5) COMMENT 'Rating given to the product (1-5)',
    comment_content TEXT NOT NULL COMMENT 'Content of the comment',
    status ENUM('visible', 'hidden', 'deleted') DEFAULT 'visible' COMMENT 'Comment status',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Time the comment was created',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
        FOREIGN KEY (user_id) REFERENCES user_service.users(id),
        FOREIGN KEY (product_id) REFERENCES product_service.products(id)
) COMMENT='Product comments table';

-- Add indexes to optimize searching
CREATE INDEX idx_product_id ON comments(product_id);
CREATE INDEX idx_user_id ON comments(user_id);
CREATE INDEX idx_status ON comments(status);
CREATE INDEX idx_rating ON comments(rating);