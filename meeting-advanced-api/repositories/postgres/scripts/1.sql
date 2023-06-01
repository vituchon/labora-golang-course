-- EJECUTAR DENTRO DE LA BASE postgres (logeado como postgres)
CREATE ROLE interspecies_bonds_role WITH
  LOGIN
  SUPERUSER
  CREATEDB
  CREATEROLE
  INHERIT
  REPLICATION
  CONNECTION LIMIT -1
  PASSWORD 'b0nds';

CREATE DATABASE interspecies_bonds_db
  WITH
  OWNER = interspecies_bonds_role
  ENCODING = 'UTF8'
  CONNECTION LIMIT = -1
  IS_TEMPLATE = false;


-- EJECUTAR DENTRO DE LA BASE interspecies_bonds_db (logeado como postgres o interspecies_bonds_role (preferentemente) )

CREATE TABLE animal (
	id   SERIAL NOT NULL,
	name TEXT NOT NULL,
	kind INT NOT NULL,
  CONSTRAINT animal_pkey PRIMARY KEY (id)
);
ALTER TABLE animal OWNER TO interspecies_bonds_role;

CREATE TABLE person (
	id   SERIAL NOT NULL,
	name TEXT NOT NULL,
	CONSTRAINT person_pkey PRIMARY KEY (id)
);
ALTER TABLE person OWNER TO interspecies_bonds_role;


CREATE TABLE bond (
	id   SERIAL NOT NULL,
	person_id INT NOT NULL,
  animal_id INT NOT NULL,
	CONSTRAINT bond_pkey PRIMARY KEY (id),
  CONSTRAINT bond_person_id_fkey FOREIGN KEY (person_id) REFERENCES person (id),
  CONSTRAINT bond_animal_id_fkey FOREIGN KEY (animal_id) REFERENCES animal (id)
);
ALTER TABLE person OWNER TO interspecies_bonds_role;
