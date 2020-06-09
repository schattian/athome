CREATE TABLE IF NOT EXISTS product_attribute_schemas(
    id serial primary key,
    name varchar(50) not null,
    category_id integer not null,
    value_type varchar(10) not null
);

CREATE TABLE IF NOT EXISTS product_attribute_datas(
    id serial primary key,
    schema_id integer not null,
    user_id integer not null,
    entity_id integer not null,
    entity_table varchar(30) not null,

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

ALTER TABLE ONLY product_attribute_datas ADD CONSTRAINT uq_product_attribute_datas UNIQUE(schema_id, entity_id, entity_table);
