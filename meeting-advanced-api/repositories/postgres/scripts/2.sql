-- EJECUTAR DENTRO DE LA BASE interspecies_bonds_db (logeado como postgres o interspecies_bonds_role (preferentemente) )

DO $$
  DECLARE
    animal_id INT;
    person_id INT;
  BEGIN
    INSERT INTO animal(name,kind) VALUES ('Kiki', 2);
    INSERT INTO animal(name,kind) VALUES ('Samson', 1);
    INSERT INTO animal(name,kind) VALUES ('Kronos', 0) returning id into animal_id;

    INSERT INTO person(name) VALUES ('Sabri') returning id into person_id;

    INSERT INTO bond(person_id, animal_id) VALUES (person_id, animal_id);
  END
$$;

