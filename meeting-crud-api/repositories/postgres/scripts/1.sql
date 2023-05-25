-- EJECUTAR DENTRO DE LA BASE postgres (logeado como postgres)
CREATE ROLE animal WITH
  LOGIN
  SUPERUSER
  CREATEDB
  CREATEROLE
  INHERIT
  REPLICATION
  CONNECTION LIMIT -1
  PASSWORD '4n1m4l';

CREATE DATABASE animals_db
  WITH
  OWNER = animal
  ENCODING = 'UTF8'
  CONNECTION LIMIT = -1
  IS_TEMPLATE = False;


-- EJECUTAR DENTRO DE LA BASE animals_db (logeado como postgres o animal (preferentemente) )
CREATE TABLE animal (
	id   SERIAL NOT NULL,
	name TEXT NOT NULL,
	kind INT NOT NULL,
	CONSTRAINT id PRIMARY KEY (id)
);
ALTER TABLE animal OWNER TO animal;

INSERT INTO animal(name,kind) VALUES ('Kiki', 2);
INSERT INTO animal(name,kind) VALUES ('Samson', 1);
INSERT INTO animal(name,kind) VALUES ('Kronos', 0);