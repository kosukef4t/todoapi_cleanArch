-- +migrate Up
CREATE TABLE if NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,        -- ユーザーID (UUIDを想定してCHAR(36)にしています)
    username VARCHAR(255) NOT NULL,     -- ユーザー名
    password VARCHAR(255) NOT NULL,     -- ハッシュ化されたパスワード
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- ユーザー作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新日時
    UNIQUE(username)                 -- ユーザー名は
);

-- +migrate Down
DROP TABLE IF EXISTS users;