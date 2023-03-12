CREATE DATABASE lazy_warehouse;

\c lazy_warehouse;

CREATE SCHEMA lazy_warehouse;

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA lazy_warehouse VERSION '1.3' CASCADE;


-- PostgreSQLで_updated_atの自動更新を擬似的に実装する
CREATE FUNCTION lazy_warehouse.trigger_update_timestamp_none() RETURNS TRIGGER AS
$$
BEGIN
    IF NEW._updated_at == OLD._updated_at THEN
        NEW._updated_at = NULL;
    END IF;
    RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE FUNCTION lazy_warehouse.trigger_update_timestamp_same() RETURNS TRIGGER AS
$$
BEGIN
    IF NEW._updated_at IS NULL THEN
        NEW._updated_at := OLD._updated_at;
    END IF;
    RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE FUNCTION lazy_warehouse.trigger_update_timestamp_current() RETURNS TRIGGER AS
$$
BEGIN
    IF NEW._updated_at IS NULL THEN
        NEW._updated_at := current_timestamp;
    END IF;
END
$$ LANGUAGE plpgsql;

-- アプリケーションユーザーの作成
CREATE ROLE app_user_for_write WITH LOGIN;
CREATE ROLE app_user_for_read WITH LOGIN;

GRANT CONNECT, TEMPORARY ON DATABASE postgres TO app_user_for_read;
GRANT CONNECT, TEMPORARY ON DATABASE postgres TO app_user_for_write;

GRANT USAGE ON SCHEMA lazy_warehouse TO app_user_for_read;
GRANT USAGE ON SCHEMA lazy_warehouse TO app_user_for_write;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA lazy_warehouse TO app_user_for_write;

GRANT SELECT ON ALL TABLES IN SCHEMA lazy_warehouse TO app_user_for_read;

REVOKE ALL ON SCHEMA public FROM PUBLIC;
