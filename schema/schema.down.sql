CREATE TABLE users (
                       id UUID  DEFAULT gen_random_uuid() PRIMARY KEY,
                       email  varchar(256) NOT NULL UNIQUE,
);

CREATE TABLE tokens (
                        user_id UUID REFERENCES users(id),
                        token_hash  varchar(256) NOT NULL,
                        ip_address  varchar(256) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        expires_at TIMESTAMP,
);
