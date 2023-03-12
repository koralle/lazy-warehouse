\c lazy_warehouse;

CREATE TABLE lazy_warehouse.users (
    id VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    profile_image_url VARCHAR(255),
    _created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    _updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO lazy_warehouse.users (id, name, email) VALUES ('1', 'Mugicha', 'mugicha@mugicha.com');
INSERT INTO lazy_warehouse.users (id, name, email) VALUES ('2', 'Mugizarashi', 'mugizarashi@mugicha.com');
