
CREATE TABLE playground (
    equip_id serial PRIMARY KEY,
    type varchar (50) NOT NULL,
    color varchar (25) NOT NULL,
    location varchar(25) check (location in ('north', 'south', 'west', 'east', 'northeast', 'southeast', 'southwest', 'northwest')),
    install_date date
);

INSERT INTO public.playground(
        type, color, location, install_date)
        VALUES ('pump', 'blue', 'south', '2017-11-03');
    
INSERT INTO public.playground(
        type, color, location, install_date)
        VALUES ('heater', 'yellow', 'northwest', '2017-11-04');