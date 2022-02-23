--liquibase formatted sql

--changeset Elliot-BAXUS:1
/*

In order to avoid collision of changesets, we will adopt the following nomenclature:

[Liquibase ID]:[Changeset ID]

Your Liquibase ID will be unique, thus ensuring that you won't conflict with another developer
It is up to each developer to label each changeset with the correct Changeset ID!

Violators will be thrown in the rancor pit

Additionally, the DATABASECHANGELOG, the table in the database that keeps the record of all changes made with Liquibase,
keeps track of changesets using this ID scheme 

*/


