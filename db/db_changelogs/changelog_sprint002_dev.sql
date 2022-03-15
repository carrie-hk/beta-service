--liquibase formatted sql

--changeset Elliot:1
USE dev;

--changeset Elliot:2
--rollback DROP TABLE dev.axu
CREATE TABLE dev.axu
(
    axu_id INT NOT NULL,
    bax_id TEXT NOT NULL,
    asc_num INT NOT NULL,
    asset_type INT NOT NULL,
    time_created TIMESTAMP NOT NULL,
    asset_status TEXT NOT NULL,
    mint_addr TEXT,
    update_addr TEXT,
    featured BOOLEAN NOT NULL,
    shelf_loc TEXT,
    PRIMARY KEY (axu_id),
    UNIQUE (axu_id)
);

--changeset Elliot:3
--rollback DROP INDEX idx_asc_num
CREATE INDEX idx_asc_num 
ON dev.axu(asc_num)

--changeset Elliot:4
--rollback DROP TABLE dev.visual_content
CREATE TABLE dev.visual_content
(
    axu_id INT NOT NULL,
    html5 TEXT,
    mp4 TEXT,
    cover TEXT,
    front TEXT,
    back TEXT,
    PRIMARY KEY (axu_id),
    FOREIGN KEY (axu_id)
        REFERENCES dev.axu(axu_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE 
);

--changeset Elliot:5
--rollback DROP TABLE dev.user
CREATE TABLE dev.user
(
    username VARCHAR(50) NOT NULL,
    passwd TEXT NOT NULL,
    two_fa TEXT,
    PRIMARY KEY (username)
);

--changeset Elliot:6
--rollback DROP TABLE dev.bttl_class
CREATE TABLE dev.bttl_class
(
    class_id INT NOT NULL,
    class_name TEXT NOT NULL,
    PRIMARY KEY (class_id),
    UNIQUE (class_id)
);

--changeset Elliot:7
--rollback DROP TABLE dev.wine_class
CREATE TABLE dev.wine_class
(
    class_name TEXT NOT NULL,
    age INT NOT NULL,
    desc_short TEXT,
    desc_long TEXT NOT NULL,
    vintage INT NOT NULL,
    cask_type TEXT,
    bttl_size INT NOT NULL,
    series TEXT,
    varietal_1 TEXT NOT NULL,
    varietal_1_pct FLOAT,
    varietal_2 TEXT,
    varietal_2_pct FLOAT,
    varietal_3 TEXT,
    varietal_3_pct FLOAT,
    varietal_4 TEXT,
    varietal_4_pct FLOAT,
    harvest TEXT,
    appellation TEXT,
    original_release_qnty INT,
    abv FLOAT,
    winery TEXT NOT NULL,
    bottler TEXT NOT NULL,
    class_id INT NOT NULL,
    PRIMARY KEY (class_id),
    FOREIGN KEY (class_id)
        REFERENCES dev.bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE 
);

--changeset Elliot:8
--rollback DROP TABLE dev.sprt_class
CREATE TABLE dev.sprt_class
(
    class_name TEXT NOT NULL,
    age INT,
    desc_short TEXT,
    desc_long TEXT NOT NULL,
    year_distilled INT,
    year_bottled INT,
    cask_type TEXT,
    cask_num TEXT,
    single_cask BOOLEAN,
    bttl_size INT,
    series TEXT,
    spirit_type TEXT NOT NULL,
    original_cask_yield INT,
    abv FLOAT,
    distillery TEXT,
    bottler TEXT,
    class_id INT NOT NULL,
    PRIMARY KEY (class_id),
    FOREIGN KEY (class_id)
        REFERENCES dev.bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE 
);

--changeset Elliot:9
--rollback DROP TABLE dev.bttl
CREATE TABLE dev.bttl
(
    axu_id INT NOT NULL,
    bottle_num INT,
    serial_num TEXT,
    barcode TEXT,
    grade TEXT NOT NULL,
    packaging_desc TEXT NOT NULL,
    class_id INT NOT NULL,
    PRIMARY KEY (axu_id),
    FOREIGN KEY (axu_id)
        REFERENCES dev.axu(axu_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE,
    FOREIGN KEY (class_id)
        REFERENCES dev.bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE
);

--changeset Elliot:10
--rollback DROP TABLE dev.winery
CREATE TABLE dev.winery
(
    name VARCHAR(50) NOT NULL,
    country TEXT NOT NULL,
    region TEXT,
    PRIMARY KEY (name)
);

--changeset Elliot:11
--rollback DROP TABLE dev.distillery
CREATE TABLE dev.distillery
(
    name VARCHAR(50) NOT NULL,
    country TEXT NOT NULL,
    region TEXT,
    smws TEXT,
    PRIMARY KEY (name)
);

--changeset Elliot:12
--rollback DROP TABLE dev.kyc
CREATE TABLE dev.kyc
(
    wallet_pk VARCHAR(50) NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone_num TEXT NOT NULL,
    email TEXT NOT NULL,
    ship_addr_a TEXT NOT NULL,
    ship_addr_b TEXT NOT NULL,
    ship_city TEXT NOT NULL,
    ship_state TEXT NOT NULL,
    ship_zip INT NOT NULL,
    dob_day INT NOT NULL,
    dob_month INT NOT NULL,
    dob_year INT NOT NULL,
    title TEXT,
    PRIMARY KEY (wallet_pk)
);

--changeset Elliot:13
--rollback DROP TABLE dev.asc
CREATE TABLE dev.asc
(
    asc_num INT NOT NULL,
    wallet_pk VARCHAR(50) NOT NULL,
    PRIMARY KEY (asc_num),
    FOREIGN KEY (asc_num)
        REFERENCES dev.axu(asc_num)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

--changeset Elliot:14
--rollback DROP INDEX idx_asc_num
CREATE INDEX idx_wallet_pk
ON dev.asc(wallet_pk)

--changeset Elliot:15
--rollback DROP TABLE dev.wallet
CREATE TABLE dev.wallet
(
    wallet_pk VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    PRIMARY KEY (wallet_pk),
    FOREIGN KEY (wallet_pk)
        REFERENCES dev.asc(wallet_pk)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

--changeset Elliot:16
--rollback DROP TABLE dev.primary_sale
CREATE TABLE dev.primary_sale
(
    axu_id INT NOT NULL,
    price FLOAT NOT NULL,
    date_listed TIMESTAMP NOT NULL,
    PRIMARY KEY (axu_id),
    FOREIGN KEY (axu_id)
        REFERENCES dev.axu(axu_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE 
);



