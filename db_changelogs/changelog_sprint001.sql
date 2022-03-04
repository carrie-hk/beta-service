--liquibase formatted sql

--changeset Elliot:1
--rollback DROP SCHEMA staging
CREATE SCHEMA staging;

--changeset Elliot:2
--rollback DROP TABLE staging.AXU
CREATE TABLE staging.AXU
(
    bax_id       varchar(10) not null,
    asc_id       int         not null,
    asset_type   varchar(3)  not null,
    asset_id     int         not null,
    time_created timestamp   not null,
    axu_status    not null,
    mint_addr varchar(32) not null,
    PRIMARY KEY (bax_id)
);

--changeset Elliot:3
--rollback DROP TABLE staging.visual_content
CREATE TABLE staging.visual_content
(
    bax_id varchar(10) not null,
    html5 text null,
    mp4 text null,
    cover text null,
    front text null,
    back text null,
    PRIMARY KEY (bax_id)
);

--changeset Elliot:4
--rollback DROP TABLE staging.users
CREATE TABLE staging.users 
(
    username varchar(32) not null,
    password varchar(32) not null,
    wallet_pk text not null,
    two_factor_auth varchar(9) null
    PRIMARY KEY (username, wallet)
);

--changeset Elliot:5
--rollback DROP TABLE staging.historical_trading
CREATE TABLE staging.historical_trading
(
    collection_id text not null,
    trade_platform text null,
    last_trade_time timestamp null,
    last_trade_usdc int null
    PRIMARY KEY (collection_id)
);

--changeset Elliot:6
--rollback DROP TABLE staging.kyc
CREATE TABLE staging.kyc
(
    wallet varchar(32) not null,
    first_name text not null,
    last_name text not null,
    phone_num varchar(9) not null,
    email varchar(32) not null,
    ship_addr_a varchar(32) not null,
    ship_addr_b varchar(32) not null,
);

--changeset Elliot:7
--rollback DROP TYPE asset_grade
CREATE TYPE asset_grade AS ENUM ('A+', 'A', 'A-', 'B+', 'B')

--changeset Elliot:8
--rollback DROP TABLE staging.wine
/*
Should we combine the staging.wine and staging.sprt tables, since they have the same fields?
*/
CREATE TABLE staging.wine
(
    bax_id varchar(10) not null,
    bottle_num int not null,
    serial_num int not null,
    grade asset_grade not null,
    packaging text not null,
    collection_id text not null
);



