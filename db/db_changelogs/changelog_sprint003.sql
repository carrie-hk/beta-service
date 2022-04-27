--liquibase formatted sql

--changeset Elliot:1
--rollback DROP PRIMARY KEY
ALTER TABLE asset_view_table
ADD PRIMARY KEY(axu_id)

--changeset Elliot:2
ALTER TABLE sprt_class
CHANGE COLUMN distillery producer TEXT;

--changeset Elliot:3
ALTER TABLE wine_class
CHANGE COLUMN winery producer TEXT;

--changeset Elliot:4
--rollback DROP COLUMN axu.holaplex_link
ALTER TABLE axu
ADD holaplex_link TEXT;

--changeset Elliot:5
DROP TABLE asset_view_table;

--changeset Elliot:6
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

--changeset Elliot:7
--rollback DROP PRIMARY KEY
ALTER TABLE asset_view_table
ADD PRIMARY KEY(axu_id);

--changeset Elliot:8 
--rollback DROP COLUMN country
ALTER TABLE kyc
ADD COLUMN ship_country TEXT AFTER ship_zip;

--changeset Elliot:9
ALTER TABLE kyc
DROP COLUMN title

--changeset Elliot:10
ALTER TABLE kyc
MODIFY COLUMN ship_zip TEXT NOT NULL;

--changeset Elliot:11
ALTER TABLE kyc
MODIFY COLUMN ship_state TEXT;

--changeset Elliot:12
ALTER TABLE kyc
DROP COLUMN dob_day;

--changeset Elliot:13
ALTER TABLE kyc
DROP COLUMN dob_month;

--changeset Elliot:14
ALTER TABLE kyc
DROP COLUMN dob_year;

--changeset Elliot:15
ALTER TABLE kyc
ADD COLUMN dob_unix_ms BIGINT NOT NULL;

--changeset Elliot:16
ALTER TABLE kyc
DROP COLUMN dob_unix_ms;

--changeset Elliot:17
ALTER TABLE kyc
ADD COLUMN dob TEXT NOT NULL;