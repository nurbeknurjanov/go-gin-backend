CREATE TABLE products (
                       id serial not null primary key,
                       name varchar not null,
                       description varchar not null,
                       created_at timestamp default CURRENT_TIMESTAMP not null,
                       updated_at timestamp default CURRENT_TIMESTAMP not null
);

CREATE TRIGGER update_products_trigger
    BEFORE UPDATE
    ON
        products
    FOR EACH ROW
    EXECUTE PROCEDURE update_dates();

INSERT INTO products (name, description) values ('Product 1', 'some description');
INSERT INTO products (name, description) values ('Product 2', 'some description');
INSERT INTO products (name, description) values ('Product 3', 'some description');
INSERT INTO products (name, description) values ('Product 4', 'some description');
INSERT INTO products (name, description) values ('Product 5', 'some description');
INSERT INTO products (name, description) values ('Product 6', 'some description');
INSERT INTO products (name, description) values ('Product 7', 'some description');
INSERT INTO products (name, description) values ('Product 8', 'some description');
INSERT INTO products (name, description) values ('Product 9', 'some description');
INSERT INTO products (name, description) values ('Product 10', 'some description');
INSERT INTO products (name, description) values ('Product 11', 'some description');
INSERT INTO products (name, description) values ('Product 12', 'some description');
INSERT INTO products (name, description) values ('Product 13', 'some description');