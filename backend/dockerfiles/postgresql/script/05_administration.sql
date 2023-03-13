\c lazy_warehouse;

CREATE TABLE lazy_warehouse.administration (
    id VARCHAR(255),
    user_id VARCHAR(255) REFERENCES lazy_warehouse.users(id) NOT NULL,
    group_id VARCHAR(255) REFERENCES lazy_warehouse.groups(id) NOT NULL,
    role_id VARCHAR(255) REFERENCES lazy_warehouse.roles(id) NOT NULL,
    _created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    _updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT user_and_group UNIQUE NULLS NOT DISTINCT (user_id, group_id) 
);

