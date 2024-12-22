-- Create database for reply service
CREATE DATABASE reply_service;

-- Use the reply_service database
USE reply_service;

-- Table to store replies to comments
CREATE TABLE replies (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Reply ID',
    comment_id BIGINT NOT NULL COMMENT 'ID of the comment being replied to',
    user_id BIGINT NOT NULL COMMENT 'User ID of the person replying',
    reply_content TEXT NOT NULL COMMENT 'Content of the reply',
    parent_reply_id BIGINT NULL COMMENT 'ID of the parent reply (NULL if it is a direct reply to the comment)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Time the reply was created',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    FOREIGN KEY (user_id) REFERENCES user_service.users(id),
    FOREIGN KEY (comment_id) REFERENCES comment_service.comments(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_reply_id) REFERENCES replies(id) ON DELETE SET NULL -- Allow cascading delete but set parent_reply_id to NULL
) COMMENT='Replies to comments table';

-- Add indexes to optimize searching
CREATE INDEX idx_comment_id ON replies(comment_id);
CREATE INDEX idx_user_id ON replies(user_id);
CREATE INDEX idx_parent_reply_id ON replies(parent_reply_id); -- Index for parent_reply_id