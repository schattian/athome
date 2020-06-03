# Backend 

## Ports

Only internal:
 - **9900**: auth
 - **9901**: mailer 
 - **9902**: identifier


Exposed svcs:
 - **9990**: users

## Naming

For instance, suppose the svc created to serve entity-related things (an
entity named `entity`).

The naming convention of the svc will be:

- name of the svc: `entities`
- name of the folder containing the svc: `entities` 
- name of the db it uses: `entity_<env>` (to avoid confusion with tables, just in case it holds a `Entity` model)
- name of the compilable .proto: `entities.proto`
- name of the exported go pkg: `pbentities
