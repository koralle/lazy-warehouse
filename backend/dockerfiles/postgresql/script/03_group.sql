\c lazy_warehouse;

CREATE TABLE lazy_warehouse.groups (
    id VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    profile_image_url VARCHAR(255),
    _created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    _updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

insert into lazy_warehouse.groups (id, name, description, profile_image_url) VALUES ('1', 'koralle''s house', 'koralleの家', 'https://example.com');
insert into lazy_warehouse.groups (id, name, description, profile_image_url) VALUES ('2', 'mugicha''s house', 'mugichaの家', 'https://example.com');
