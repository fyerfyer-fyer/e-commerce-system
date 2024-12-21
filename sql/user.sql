-- Create database for user service
CREATE DATABASE user_service;

-- Use the user_service database
USE user_service;

-- Table to store user information
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT 'Username',
    password VARCHAR(255) NOT NULL COMMENT 'Password (hashed)',
    email VARCHAR(100) NOT NULL UNIQUE COMMENT 'Email address',
    phone VARCHAR(20) NOT NULL UNIQUE COMMENT 'Phone number',
    avatar VARCHAR(255) COMMENT 'User avatar URL',
    role ENUM('user', 'admin', 'super_admin') DEFAULT 'user' COMMENT 'User role',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time'
) COMMENT='User information table';

-- Add indexes to optimize searching
CREATE INDEX idx_username ON users(username);
CREATE INDEX idx_email ON users(email);
CREATE INDEX idx_phone ON users(phone); -- Index for phone login


-- Table to store user addresses
CREATE TABLE user_addresses (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Address ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    address_line1 VARCHAR(255) NOT NULL COMMENT 'Address line 1',
    address_line2 VARCHAR(255) COMMENT 'Address line 2',
    city VARCHAR(100) NOT NULL COMMENT 'City',
    state VARCHAR(100) NOT NULL COMMENT 'State/Province',
    postal_code VARCHAR(20) NOT NULL COMMENT 'Postal code',
    country VARCHAR(100) NOT NULL COMMENT 'Country',
    is_default BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Is default address',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
     INDEX idx_user_id_default (user_id, is_default)
) COMMENT='User address information table';

-- Index for UserID
CREATE INDEX idx_user_id ON user_addresses(user_id);

-- Table to store admin codes
CREATE TABLE admin_codes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Admin Code ID',
    code VARCHAR(255) UNIQUE NOT NULL COMMENT 'The actual admin code',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    used BOOLEAN DEFAULT FALSE COMMENT 'is code used',
    expired_at TIMESTAMP COMMENT 'code expired time'
) COMMENT='Admin code information';


-- Table to store user login history (Reduced information for better performance)
CREATE TABLE login_history (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Login history ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Login time',
        login_ip VARCHAR(20),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) COMMENT='User login history table';

-- Add indexes to optimize searching
CREATE INDEX idx_user_id_login_time ON login_history(user_id, login_time);

-- Table to store user collection list (wishlist)
CREATE TABLE user_collections (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Collection entry ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID (linked to product service)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Time the product was added',
        UNIQUE KEY user_product (user_id, product_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) COMMENT='User collection/wishlist table';