# Backend 

## Naming

For instance, suppose the svc created to serve entity-related things (an
entity named `entity`).

The naming convention of the svc will be:

- name of the svc: `entities` (plural)
- name of the folder containing the svc: `entities` (plural)
- name of the db it uses: `entity_<env>` (to avoid confusion with tables, just in case it holds a `Entity` model)
- name of the compilable .proto: `entity.proto`
- name of the exported go pkg: `pbentity`