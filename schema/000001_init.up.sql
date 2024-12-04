CREATE DOMAIN sex_type AS integer
    CHECK (
        VALUE IN
        (1, 2)
        );
CREATE TYPE status_type AS ENUM ('0', '1');

CREATE TABLE users (
                       id serial not null primary key,
                       email varchar not null unique,
                       encrypted_password varchar not null,
                       name varchar not null,
                       age integer null,
                       sex sex_type null,
                       status status_type not null default '0',
                       created_at timestamp default CURRENT_TIMESTAMP not null,
                       updated_at timestamp default CURRENT_TIMESTAMP not null
);

CREATE OR REPLACE FUNCTION update_dates()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_trigger
    BEFORE UPDATE
    ON
        users
    FOR EACH ROW
    EXECUTE PROCEDURE update_dates();


INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('nurbek.nurjanov@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Nurbek Nurjanov', 39, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('alan@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Alan Parker', 39, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('bob@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Bob Martin', 39, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('celine@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Celine Dion', 39, 2, '0');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('trump@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Donald Trump', 80, 1, '0');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('biden@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Joe biden', 80, 1, '0');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('hanz@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Hanz Zimmer', 60, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('kamala@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Kamala Harris', 50, 2, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('hilary@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Hilary Clinton', 55, 2, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('bill@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Bill Clinton', 56, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('ben@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Ben Stiller', 56, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('cat@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Cat Mew', 56, 1, '1');
INSERT INTO users (email, encrypted_password, name, age, sex, status)
values ('dog@mail.ru', '$2a$04$TFjzk4E94FrF282cLSNwwuE62WBuZrg2W/b0CNzHFU/L55Lb8.JM6',
        'Hot Dog', 56, 1, '1');