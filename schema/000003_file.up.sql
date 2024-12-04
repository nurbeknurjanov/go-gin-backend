CREATE TABLE files (
                          id serial not null primary key,
                          original_file_name varchar not null,
                          ext varchar not null,
                          data varchar null,
                          model_name varchar null,
                          model_id integer null,
                          uuid varchar not null,
                          created_at timestamp default CURRENT_TIMESTAMP not null,
                          updated_at timestamp default CURRENT_TIMESTAMP not null
);


CREATE TRIGGER update_files_trigger
    BEFORE UPDATE
    ON
        files
    FOR EACH ROW
    EXECUTE PROCEDURE update_dates();