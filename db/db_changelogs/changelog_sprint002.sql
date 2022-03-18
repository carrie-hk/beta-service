--liquibase formatted sql

--changeset Elliot:1
--rollback DROP TABLE axu
CREATE TABLE axu
(
    axu_id INT NOT NULL,
    bax_id TEXT NOT NULL,
    asc_num INT NOT NULL,
    asset_type INT NOT NULL,
    time_created TIMESTAMP NOT NULL,
    asset_status TEXT NOT NULL,
    token_addr TEXT,
    mint_addr TEXT,
    update_addr TEXT,
    featured BOOLEAN NOT NULL,
    shelf_loc TEXT,
    price FLOAT,
    PRIMARY KEY (axu_id),
    UNIQUE (axu_id)
);

--changeset Elliot:2
--rollback DROP INDEX idx_asc_num
CREATE INDEX idx_asc_num 
ON axu(asc_num)

--changeset Elliot:3
--rollback DROP TABLE user
CREATE TABLE user
(
    username VARCHAR(50) NOT NULL,
    passwd TEXT NOT NULL,
    two_fa TEXT,
    PRIMARY KEY (username)
);

--changeset Elliot:4
--rollback DROP TABLE bttl_class
CREATE TABLE bttl_class
(
    class_id INT NOT NULL,
    class_name TEXT NOT NULL,
    PRIMARY KEY (class_id),
    UNIQUE (class_id)
);

--changeset Elliot:5
--rollback DROP TABLE wine_class
CREATE TABLE wine_class
(
    name TEXT NOT NULL,
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
    UNIQUE (class_id),
    FOREIGN KEY (class_id)
        REFERENCES bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE 
);

--changeset Elliot:6
--rollback DROP TABLE sprt_class
CREATE TABLE sprt_class
(
    name TEXT NOT NULL,
    age INT,
    desc_short TEXT,
    desc_long TEXT NOT NULL,
    year_distilled YEAR,
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
    UNIQUE (class_id),
    FOREIGN KEY (class_id)
        REFERENCES bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE 
);

--changeset Elliot:7
--rollback DROP TABLE wine_bttl
CREATE TABLE wine_bttl
(
    axu_id INT NOT NULL,
    bottle_num INT,
    serial_num TEXT,
    barcode TEXT,
    grade TEXT NOT NULL,
    packaging_desc TEXT NOT NULL,
    class_id INT NOT NULL,
    html5 TEXT,
    mp4 TEXT,
    cover TEXT,
    front TEXT,
    back TEXT,
    PRIMARY KEY (axu_id),
    FOREIGN KEY (axu_id)
        REFERENCES axu(axu_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE,
    FOREIGN KEY (class_id)
        REFERENCES bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE
);

--changeset Elliot:8
--rollback DROP TABLE winery
CREATE TABLE winery
(
    id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    country TEXT NOT NULL,
    region TEXT,
    PRIMARY KEY (id)
);

--changeset Elliot:9
--rollback DROP TABLE sprt_bttl
CREATE TABLE sprt_bttl
(
    axu_id INT NOT NULL,
    bottle_num INT,
    serial_num TEXT,
    barcode TEXT,
    grade TEXT NOT NULL,
    packaging_desc TEXT NOT NULL,
    class_id INT NOT NULL,
    html5 TEXT,
    mp4 TEXT,
    cover TEXT,
    front TEXT,
    back TEXT,
    PRIMARY KEY (axu_id),
    FOREIGN KEY (axu_id)
        REFERENCES axu(axu_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE,
    FOREIGN KEY (class_id)
        REFERENCES bttl_class(class_id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE
);

--changeset Elliot:10
--rollback DROP TABLE distillery
CREATE TABLE distillery
(
    id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    country TEXT NOT NULL,
    region TEXT,
    smws TEXT,
    PRIMARY KEY (id)
);

--changeset Elliot:11
--rollback DROP TABLE kyc
CREATE TABLE kyc
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

--changeset Elliot:12
--rollback DROP TABLE asc
CREATE TABLE asc_token
(
    asc_num INT NOT NULL,
    wallet_pk VARCHAR(50) NOT NULL,
    PRIMARY KEY (asc_num),
    FOREIGN KEY (asc_num)
        REFERENCES axu(asc_num)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

--changeset Elliot:13
--rollback DROP INDEX idx_asc_num
CREATE INDEX idx_wallet_pk
ON asc_token(wallet_pk)

--changeset Elliot:14
--rollback DROP TABLE wallet
CREATE TABLE wallet
(
    wallet_pk VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    PRIMARY KEY (wallet_pk),
    FOREIGN KEY (wallet_pk)
        REFERENCES asc_token(wallet_pk)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

--changeset Elliot:15
--rollback DROP TABLE asset_view_table
CREATE TABLE asset_view_table
(
SELECT axu.axu_id,
       asc_num,
       asset_status,
       token_addr,
       mint_addr,
       update_addr,
       price,
       featured,
       html5,
       cover,
       name,
       age,
       desc_short,
       desc_long,
       abv
FROM axu
         inner join (SELECT sprt_bttl.axu_id,
                            html5,
                            cover,
                            name,
                            age,
                            desc_short,
                            desc_long,
                            abv
                     from sprt_bttl
                              inner join sprt_class on sprt_bttl.class_id = sprt_class.class_id) as A
                    on axu.axu_id = A.axu_id
UNION

SELECT axu.axu_id,
       asc_num,
       asset_status,
       token_addr,
       mint_addr,
       update_addr,
       price,
       featured,
       html5,
       cover,
       name,
       age,
       desc_short,
       desc_long,
       abv
from axu
         inner join (SELECT wine_bttl.axu_id,
                            html5,
                            cover,
                            name,
                            age,
                            desc_short,
                            desc_long,
                            abv
                     from wine_bttl
                              inner join wine_class on wine_bttl.class_id = wine_class.class_id) as B
                    on axu.axu_id = B.axu_id
);


