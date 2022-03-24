--liquibase formatted sql

--changeset Elliot:1
--rollback DROP PRIMARY KEY
ALTER TABLE asset_view_table
ADD PRIMARY KEY(axu_id)

--changeset Elliot:2
ALTER TABLE sprt_class
ADD producer INT;

--changeset Elliot:3
ALTER TABLE sprt_class
ADD FOREIGN KEY (producer)
    REFERENCES distillery(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

--changeset Elliot:4
ALTER TABLE wine_class
ADD producer INT;

--changeset Elliot:5
ALTER TABLE wine_class
ADD FOREIGN KEY (producer)
    REFERENCES winery(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE;

--changeset Elliot:6
--rollback DROP COLUMN axu.holaplex_link
ALTER TABLE axu
ADD holaplex_link TEXT;

--changeset Elliot:7
DROP TABLE asset_view_table;

--changeset Elliot:8
--rollback DROP TABLE asset_view_table
CREATE TABLE asset_view_table
(
SELECT axu.axu_id,
       asc_num,
       time_created,
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
       abv,
       bttl_size,
       producer,
       holaplex_link
FROM axu
         inner join (SELECT sprt_bttl.axu_id,
                            html5,
                            cover,
                            name,
                            age,
                            desc_short,
                            desc_long,
                            abv,
                            bttl_size,
                            producer
                     from sprt_bttl
                              inner join sprt_class on sprt_bttl.class_id = sprt_class.class_id) as A
                    on axu.axu_id = A.axu_id
UNION

SELECT axu.axu_id,
       asc_num,
       time_created,
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
       abv,
       bttl_size,
       producer,
       holaplex_link
from axu
         inner join (SELECT wine_bttl.axu_id,
                            html5,
                            cover,
                            name,
                            age,
                            desc_short,
                            desc_long,
                            abv,
                            bttl_size,
                            producer
                     from wine_bttl
                              inner join wine_class on wine_bttl.class_id = wine_class.class_id) as B
                    on axu.axu_id = B.axu_id
); 

--changeset Elliot:9
--rollback DROP PRIMARY KEY
ALTER TABLE asset_view_table
ADD PRIMARY KEY(axu_id);

--changeset Elliot:10
--rollback DROP TABLE redemption_program_info
CREATE TABLE redemption_info
(
    id INT NOT NULL AUTO_INCREMENT,
    wallet_pk TEXT NOT NULL,
    redemption_info_accnt_addr TEXT NOT NULL,
    baxus_escrow_addr TEXT NOT NULL,
    mint_addr TEXT NOT NULL,
    PRIMARY KEY (id)
);
