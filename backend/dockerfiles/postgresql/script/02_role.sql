\c lazy_warehouse;

CREATE TABLE lazy_warehouse.roles (
    id VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    _created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    _updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- まず１つ目の関数を実行
CREATE TRIGGER update_roles_updated_at_step1
    BEFORE UPDATE ON lazy_warehouse.roles FOR EACH ROW
    EXECUTE PROCEDURE lazy_warehouse.trigger_update_timestamp_none();

-- _updated_atカラムが更新された時、２つ目の関数を実行
CREATE TRIGGER update_roles_updated_at_step2
    BEFORE UPDATE OF _updated_at ON lazy_warehouse.roles FOR EACH ROW
    EXECUTE PROCEDURE lazy_warehouse.trigger_update_timestamp_same();

-- 最後に３つ目の関数を実行
CREATE TRIGGER update_roles_updated_at_step3
    BEFORE UPDATE ON lazy_warehouse.roles FOR EACH ROW
    EXECUTE PROCEDURE lazy_warehouse.trigger_update_timestamp_current();

GRANT ALL PRIVILEGES ON lazy_warehouse.roles TO app_user_for_read;

INSERT INTO lazy_warehouse.roles (id, name) VALUES ('1', upper('Owner'));
INSERT INTO lazy_warehouse.roles (id, name) VALUES ('2', upper('Editor'));
INSERT INTO lazy_warehouse.roles (id, name) VALUES ('3', upper('Reader'));
INSERT INTO lazy_warehouse.roles (id, name) VALUES ('4', upper('None'));
