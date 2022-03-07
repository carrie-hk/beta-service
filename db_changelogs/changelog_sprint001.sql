--liquibase formatted sql

--changeset Elliot:1
--rollback DROP SCHEMA staging
CREATE SCHEMA staging;

--changeset Elliot:2
--rollback DROP TYPE axu_status
CREATE TYPE axu_status AS ENUM ('Processing','Minted','Listed','Sold', 'Escrow', 'Redeemed')

--changeset Elliot:3
--rollback DROP TABLE staging.axu
CREATE TABLE staging.axu
(
    bax_id text not null,
    asc_id int not null,
    asset_type text not null,
    asset_id serial not null,
    time_created timestamp not null,
    asset_status axu_status not null,
    mint_addr text not null,
    PRIMARY KEY (bax_id)
);

--changeset Elliot:4
--rollback DROP TABLE staging.visual_content
CREATE TABLE staging.visual_content
(
    bax_id text not null,
    html5 text null,
    mp4 text null,
    cover text null,
    front text null,
    back text null,
    PRIMARY KEY (bax_id)
);

--changeset Elliot:5
--rollback DROP TABLE staging.historical_trading
CREATE TABLE staging.historical_trading
(
    collection_id text not null,
    trade_platform text null,
    last_trade_time timestamp null,
    last_trade_usdc int null,
    PRIMARY KEY (collection_id)
);

--changeset Elliot:6
--rollback DROP TYPE kyc_title
CREATE TYPE kyc_title AS ENUM ('Mrs.','Mr.','Sir','Madame')

--changeset Elliot:7
--rollback DROP TABLE staging.kyc
CREATE TABLE staging.kyc
(
    wallet_pk text not null,
    username text not null,
    first_name text not null,
    last_name text not null,
    phone_num text not null,
    email text not null,
    ship_addr_a text not null,
    ship_addr_b text not null,
    ship_city text not null,
    -- Create enum with US states
    ship_state text not null,
    ship_zip int not null,
    dob_day int not null,
    dob_month int not null,
    dob_year int not null,
    title kyc_title null,
    PRIMARY KEY (username)
);

--changeset Elliot:8
--rollback DROP TYPE asset_grade
CREATE TYPE asset_grade AS ENUM ('A+', 'A', 'A-', 'B+', 'B')

--changeset Elliot:9
--rollback DROP TABLE staging.bttl
CREATE TABLE staging.bttl
(
    bax_id text not null,
    bottle_num int not null,
    serial_num int not null,
    grade asset_grade not null,
    packaging_desc text not null,
    collection_id text not null
);

--changeset Elliot:10
--rollback DROP TABLE staging.wine_class
CREATE TABLE staging.wine_class
(
    class_name text not null,
    age int not null,
    desc_short text null,
    desc_long text not null,
    vintage int not null,
    cask_type text null,
    bttl_size int not null,
    series text null,
    varietal_1 text not null,
    varietal_1_pct float null,
    varietal_2 text null,
    varietal_2_pct float null,
    varietal_3 text null,
    varietal_3_pct float null,
    varietal_4 text null,
    varietal_4_pct float null,
    harvest text null,
    country text not null,
    appellation text null,
    original_release_qnty int null,
    region text not null,
    abv float not null,
    winery text not null,
    bottler text not null,
    collection_id text not null,
    PRIMARY KEY (class_name)
);

--changeset Elliot:11
--rollback DROP TABLE staging.sprt_class
CREATE TABLE staging.sprt_class
();



