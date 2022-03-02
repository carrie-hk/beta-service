--liquibase formatted sql

--changeset Elliot-BAXUS:1
/*

In order to avoid collision of changesets, and to keep track of who made which changes,
each changeset uses the following attribute nomenclature:

[Liquibase ID]:[Changeset ID]

Your Liquibase ID will be unique, thus ensuring that you won't conflict with another developer
It is up to each developer to label each changeset with the correct Changeset ID! Violators will be thrown in the rancor pit

Additionally, the DATABASECHANGELOG, the tracking table in the database that keeps the record of all changes made with Liquibase,
keeps track of changesets using these attributes

Each changeset should be atomic - i.e. one single change to the database schema - this makes debugging easier

Every changeset should have a --rollback [ACTION] [Liquibase ID] command directly under the changeset, where the ACTION is the same as the one you're
adding in the changeset - this is how Liquibase knows what to undo if the 'rollback' command is executed
Example from the Liquibase Github repo:
--changeset yourname:yourname1
--rollback DROP TABLE yourname;
CREATE TABLE yourname (
    id int primary key,
    name varchar(50) not null
)

*/

--changeset Elliot-BAXUS:2
--rollback DROP TABLE TestTable
CREATE TABLE TestTable (
    id int primary key,
    name varchar(50) not null
)


