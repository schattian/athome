# Backend 

## Ports

Only internal:
 - **9900**: auth
 - **9901**: mailer 
 - **9902**: identifier


Exposed svcs:
 - **9990**: users
 - **9991**: semantic 
 - **9992**: products 
 - **9993**: images
 - **9994**: services 
 - **9995**: address
 - **9996**: notifier 
 - **9997**: messenger 
 - **9998**: checkout 
 - **9999**: agreement

## Naming

For instance, suppose the svc created to serve entity-related things (an
entity named `entity`).

The naming convention of the svc will be:

- name of the svc: `entities`
- name of the folder containing the svc: `entities` 
- name of the db it uses: `entity_<env>` (to avoid confusion with tables, just in case it holds a `Entity` model)
- name of the compilable .proto: `entities.proto`
- name of the exported go pkg: `pbentities
