--liquibase formatted sql

--changeset Elliot-BAXUS:1
--rollback DROP TABLE TestTable
CREATE TABLE TestTable (
    id int primary key,
    name varchar(50) not null
)


