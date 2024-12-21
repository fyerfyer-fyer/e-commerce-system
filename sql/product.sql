-- Create database for product service
CREATE DATABASE product_service;

-- Use the product_service database
USE product_service;

-- Table to store product categories (Create this FIRST)
CREATE TABLE product_categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Category ID',
    name VARCHAR(100) NOT NULL COMMENT 'Category name',
    parent_id BIGINT DEFAULT NULL COMMENT 'Parent category ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (parent_id) REFERENCES product_categories(id) ON DELETE SET NULL
) COMMENT='Product categories table';

-- Add indexes to optimize searching
CREATE INDEX idx_name ON product_categories(name);
CREATE INDEX idx_parent_id ON product_categories(parent_id);


-- Table to store product information (Create this SECOND, after categories)
CREATE TABLE products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Product ID',
    name VARCHAR(255) NOT NULL COMMENT 'Product name',
    description TEXT COMMENT 'Product description',
    price DECIMAL(10, 2) NOT NULL COMMENT 'Product price',
    stock INT NOT NULL COMMENT 'Product stock quantity',
    category_id BIGINT COMMENT 'Category ID',
    brand VARCHAR(100) COMMENT 'Product brand',
    status ENUM('available', 'unavailable', 'discontinued') DEFAULT 'available' COMMENT 'Product status',
        sales_volume INT DEFAULT 0 COMMENT 'Product sales volume',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (category_id) REFERENCES product_categories(id)
) COMMENT='Product information table';

-- Add indexes to optimize searching
CREATE INDEX idx_name ON products(name);
CREATE INDEX idx_category_id ON products(category_id);
CREATE INDEX idx_status ON products(status);

-- Table to store product images
CREATE TABLE product_images (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Image ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID',
    image_url VARCHAR(255) NOT NULL COMMENT 'Image URL',
    is_primary BOOLEAN DEFAULT FALSE COMMENT 'Is primary image',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
) COMMENT='Product images table';

-- Add indexes to optimize searching
CREATE INDEX idx_product_id ON product_images(product_id);
CREATE INDEX idx_is_primary ON product_images(is_primary);

-- Table to store product attributes
CREATE TABLE product_attributes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Attribute ID',
    product_id BIGINT NOT NULL COMMENT 'Product ID',
    attribute_name VARCHAR(100) NOT NULL COMMENT 'Attribute name',
    attribute_value VARCHAR(255) NOT NULL COMMENT 'Attribute value',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
) COMMENT='Product attributes table';

-- Add indexes to optimize searching
CREATE INDEX idx_product_id ON product_attributes(product_id);
CREATE INDEX idx_attribute_name ON product_attributes(attribute_name);