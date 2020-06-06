CREATE TABLE IF NOT EXISTS product_attributes_schema(
    id serial primary key,
    name varchar(50) not null,
    category_id integer not null,
    value_type varchar(10) not null,
    -- 
    bool_value boolean,
    string_value varchar(100),
    int_64_value integer,
    float_64_value float8,
    -- 
    sl_string_value varchar(100) [],
    sl_int_64_value int [],
    sl_float_64_value float8
);

CREATE TABLE IF NOT EXISTS product_attributes_data(
    id serial primary key,
    schema_id integer not null,
    entity_id integer not null,
    entity_table string not null,

    --    
    bool_value boolean,
    string_value varchar(100),
    int_64_value integer,
    float_64_value float8,
    -- 
    sl_string_value varchar(100) [],
    sl_int_64_value int [],
    sl_float_64_value float8
);

ALTER TABLE ONLY product_attributes_data ADD CONSTRAINT uq_product_attributes_data UNIQUE(schema_id, entity_id, entity_table);
