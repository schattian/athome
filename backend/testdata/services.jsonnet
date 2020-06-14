local calendars = import './calendars.jsonnet';
local users = import './users.jsonnet';

{
  foo: {
    id: 1543934,
    user_id: users.service_providers.medic.foo.id,
    calendar_id: calendars.foo.medic.a.id,
    // "address_id":  ,

    title: 'fooServiceName',
    duration_in_minutes: 120,
    price_min: 20000,
    price_max: 100000,
  },

  bar: {
    id: 143934,
    user_id: users.service_providers.medic.bar.id,
    calendar_id: calendars.bar.medic.a.id,
    // "address_id":  ,

    title: 'barServiceName',
    duration_in_minutes: 120,
    price_min: 20000,
    price_max: 100000,
  },
}
