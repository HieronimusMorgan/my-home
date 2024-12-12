CREATE TABLE users
(
    user_id         SERIAL PRIMARY KEY,                                                        -- Auto-incremented unique identifier
    uuid_key        VARCHAR(255) UNIQUE NOT NULL,                                              -- Unique UUID for the user
    client_id       VARCHAR(255) UNIQUE NOT NULL,                                              -- Unique Client ID
    username        VARCHAR(255) UNIQUE NOT NULL,                                              -- Unique username
    password        VARCHAR(255)        NOT NULL,                                              -- Hashed password
    first_name      VARCHAR(255)        NOT NULL,                                              -- First name of the user
    last_name       VARCHAR(255)        NOT NULL,                                              -- Last name of the user
    full_name       VARCHAR(255)        NOT NULL,                                              -- Full name (can be a computed column)
    phone_number    VARCHAR(20) UNIQUE  NOT NULL,                                              -- Unique phone number
    profile_picture TEXT,                                                                      -- Optional profile picture URL or path
    role_id         INT                 NOT NULL REFERENCES roles (role_id) ON DELETE CASCADE, -- Foreign key to the roles table
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                                       -- Timestamp of creation
    created_by      VARCHAR(255),                                                              -- Who created this record
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                                       -- Timestamp of last update
    updated_by      VARCHAR(255),                                                              -- Who updated this record
    deleted_at      TIMESTAMP,                                                                 -- Timestamp of soft deletion (nullable for non-deleted)
    deleted_by      VARCHAR(255)                                                               -- Who deleted this record
);

CREATE TABLE roles
(
    role_id    SERIAL PRIMARY KEY,                  -- Auto-incremented unique identifier for roles
    name       VARCHAR(255) UNIQUE NOT NULL,        -- Unique name of the role (e.g., Admin, User)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of creation
    created_by VARCHAR(255),                        -- User or system that created this role
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of last update
    updated_by VARCHAR(255),                        -- User or system that updated this role
    deleted_at TIMESTAMP,                           -- Timestamp for soft deletion (nullable for active records)
    deleted_by VARCHAR(255)                         -- User or system that deleted this role
);

CREATE TABLE tokens
(
    token_id      SERIAL PRIMARY KEY,                                                     -- Auto-incremented unique identifier for tokens
    user_id       INT     NOT NULL,                                                       -- Foreign key referencing the users table
    token         TEXT    NOT NULL,                                                       -- The token string
    refresh_token TEXT    NOT NULL,
    expired       BOOLEAN NOT NULL DEFAULT FALSE,                                         -- Indicates if the token is expired
    created_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,                             -- Timestamp of creation
    created_by    VARCHAR(255),                                                           -- User or system that created the token
    updated_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,                             -- Timestamp of last update
    updated_by    VARCHAR(255),                                                           -- User or system that updated the token
    deleted_at    TIMESTAMP,                                                              -- Timestamp for soft deletion (nullable for active records)
    deleted_by    VARCHAR(255),                                                           -- User or system that deleted the token
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE -- Foreign key constraint
);

CREATE TABLE product_categories
(
    category_id SERIAL PRIMARY KEY,                  -- Auto-incrementing primary key
    name        VARCHAR(255) NOT NULL UNIQUE,        -- Unique name for the category
    description TEXT,                                -- Optional description
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Default to current timestamp
    created_by  VARCHAR(255),                        -- User who created the record
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Default to current timestamp
    updated_by  VARCHAR(255),                        -- User who last updated the record
    deleted_at  TIMESTAMP,                           -- Timestamp for soft deletion
    deleted_by  VARCHAR(255)                         -- User who deleted the record
);

CREATE TABLE products
(
    product_id     SERIAL PRIMARY KEY,                  -- Auto-incrementing primary key
    name           VARCHAR(255)   NOT NULL,             -- Product name
    description    TEXT,                                -- Optional description
    price          NUMERIC(10, 2) NOT NULL,             -- Product price with precision
    stock_quantity INT            NOT NULL,             -- Stock quantity
    category_id    INT            NOT NULL,             -- Foreign key to product_categories
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Default to current timestamp
    created_by     VARCHAR(255),                        -- User who created the record
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Default to current timestamp
    updated_by     VARCHAR(255),                        -- User who last updated the record
    deleted_at     TIMESTAMP,                           -- Timestamp for soft deletion
    deleted_by     VARCHAR(255),                        -- User who deleted the record
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES product_categories (category_id) ON DELETE CASCADE
);

CREATE TABLE financials
(
    financial_id SERIAL PRIMARY KEY,                                                      -- Auto-incrementing primary key
    user_id      INT            NOT NULL,                                                 -- Foreign key reference to the Users table
    balance      NUMERIC(10, 2) NOT NULL,                                                 -- Financial balance with precision
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                                     -- Timestamp for creation
    created_by   VARCHAR(255),                                                            -- User who created the record
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                                     -- Timestamp for last update
    updated_by   VARCHAR(255),                                                            -- User who last updated the record
    deleted_at   TIMESTAMP,                                                               -- Timestamp for soft deletion
    deleted_by   VARCHAR(255),                                                            -- User who deleted the record
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE -- Foreign key constraint
);


CREATE TABLE asset_categories
(
    asset_category_id SERIAL PRIMARY KEY,
    name              VARCHAR(255) NOT NULL UNIQUE,
    description       TEXT,
    created_at        TIMESTAMP,
    created_by        VARCHAR(255),
    updated_at        TIMESTAMP,
    updated_by        VARCHAR(255),
    deleted_at        TIMESTAMP,
    deleted_by        VARCHAR(255)
);

-- Table for Asset
CREATE TABLE assets
(
    asset_id          SERIAL PRIMARY KEY,
    name              VARCHAR(255)   NOT NULL,
    description       TEXT,
    value             NUMERIC(18, 2) NOT NULL,
    acquisition_date  TIMESTAMP,
    asset_category_id INT            NOT NULL,
    created_at        TIMESTAMP,
    created_by        VARCHAR(255),
    updated_at        TIMESTAMP,
    updated_by        VARCHAR(255),
    deleted_at        TIMESTAMP,
    deleted_by        VARCHAR(255),
    FOREIGN KEY (asset_category_id) REFERENCES asset_categories (asset_category_id)
);

-- Table for AssetMaintenance
CREATE TABLE asset_maintenances
(
    asset_maintenance_id  SERIAL PRIMARY KEY,
    asset_id              INT            NOT NULL,
    cost                  NUMERIC(18, 2) NOT NULL,
    notes                 TEXT,
    maintenance_date      TIMESTAMP,
    next_maintenance_date TIMESTAMP,
    created_at            TIMESTAMP,
    created_by            VARCHAR(255),
    updated_at            TIMESTAMP,
    updated_by            VARCHAR(255),
    deleted_at            TIMESTAMP,
    deleted_by            VARCHAR(255),
    FOREIGN KEY (asset_id) REFERENCES assets (asset_id)
);

CREATE
OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at
= CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_asset_category
    BEFORE UPDATE
    ON asset_category
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_asset
    BEFORE UPDATE
    ON asset
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_asset_maintenance
    BEFORE UPDATE
    ON asset_maintenance
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_financial
    BEFORE UPDATE
    ON financial
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_product_category
    BEFORE UPDATE
    ON product_category
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_product
    BEFORE UPDATE
    ON product
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_role
    BEFORE UPDATE
    ON role
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_token
    BEFORE UPDATE
    ON token
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_user
    BEFORE UPDATE
    ON master_data."user"
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();