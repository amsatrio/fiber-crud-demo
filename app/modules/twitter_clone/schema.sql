CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    email VARCHAR(255) NOT NULL,
    display_name VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    bio VARCHAR(160) DEFAULT '',
    profile_image_url VARCHAR(512) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE KEY uk_username (username),
    UNIQUE KEY uk_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE tweets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    content VARCHAR(280) NOT NULL,
    parent_tweet_id BIGINT UNSIGNED DEFAULT NULL, -- Points to parent tweet if reply or quote
    likes_count INT UNSIGNED DEFAULT 0,
    retweets_count INT UNSIGNED DEFAULT 0,
    replies_count INT UNSIGNED DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_tweets_user FOREIGN KEY (user_id) 
        REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_tweets_parent FOREIGN KEY (parent_tweet_id) 
        REFERENCES tweets(id) ON DELETE SET NULL,
        
    -- Composite index optimized for user profile timelines ORDER BY created_at DESC
    KEY idx_user_created (user_id, created_at DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE follows (
    follower_id BIGINT UNSIGNED NOT NULL,
    followed_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (follower_id, followed_id),
    
    CONSTRAINT fk_follows_follower FOREIGN KEY (follower_id) 
        REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_follows_followed FOREIGN KEY (followed_id) 
        REFERENCES users(id) ON DELETE CASCADE,
        
    -- Essential for looking up "Who is following User X?"
    KEY idx_followed (followed_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE likes (
    user_id BIGINT UNSIGNED NOT NULL,
    tweet_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (user_id, tweet_id),
    
    CONSTRAINT fk_likes_user FOREIGN KEY (user_id) 
        REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_likes_tweet FOREIGN KEY (tweet_id) 
        REFERENCES tweets(id) ON DELETE CASCADE,
        
    KEY idx_tweet (tweet_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;