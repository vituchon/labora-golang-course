-- EJECUTAR DENTRO DE LA BASE interspecies_bonds_db (logeado como postgres o interspecies_bonds_role (preferentemente) )

DO $$
  DECLARE
    cat_names text[] =  ARRAY['Igui','Salem','Agatha','Luna','Kika','Zondo','Pampita'];
  	cat_name text;
	  cat_kind_value int = 0;

	  person_names text[] =  ARRAY['Ale','Aru','Jeral','Ludmi','Mary','Poly','Rai','Sabri','Thali','Titita','Valen','Vitu'];
	  person_name text;

    sabri_person_id int;
    cat_id int;
  BEGIN

  	FOREACH person_name IN ARRAY person_names
	  LOOP
		  INSERT INTO person (name) VALUES (person_name);
	  END LOOP;

    select id from person where name = 'Sabri' into sabri_person_id;

    --raise notice 'sabri_person_id: %', sabri_person_id;

	  FOREACH cat_name IN ARRAY cat_names
	  LOOP
		  INSERT INTO animal (name,kind) VALUES (cat_name,cat_kind_value) returning id into cat_id;
		  INSERT INTO bond(person_id,animal_id) VALUES (sabri_person_id, cat_id);
	  END LOOP;

  END
$$;