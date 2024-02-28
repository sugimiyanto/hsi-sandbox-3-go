CREATE TABLE item (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100),
    status VARCHAR(10), -- item status, either active or inactive
    amount INTEGER -- how many item left (stock) 
);

CREATE TABLE item_detail (
    id SERIAL PRIMARY KEY,
    item_id VARCHAR(36) REFERENCES item(id),
    name VARCHAR(200) -- more detailed item name (description)
);