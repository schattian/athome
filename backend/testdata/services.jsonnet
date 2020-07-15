local addresses = import './addresses.jsonnet';
local calendars = import './calendars.jsonnet';
local users = import './users.jsonnet';

{
  foo: {
    id: 1543934,
    user_id: users.service_providers.medic.foo.id,
    calendar_id: calendars.foo.medic.a.id,
    address_id: addresses.service_providers.medic.foo.id,

    title: 'foomessengerName',
    duration_in_minutes: 120,
    price_min: 0,
    price_max: 10,
  },

  bar: {
    id: 143934,
    user_id: users.service_providers.medic.bar.id,
    calendar_id: calendars.bar.medic.a.id,
    address_id: addresses.service_providers.medic.bar.id,

    title: 'barServiceName',
    duration_in_minutes: 120,
    price_min: 0,
    price_max: 50,
  },

  delivery: {
    id: 1543934,
    user_id: users.service_providers.medic.foo.id,
    calendar_id: calendars.foo.delivery.a.id,
    address_id: addresses.service_providers.medic.foo.id,

    title: 'foomessengerName',
    duration_in_minutes: 120,
    price_min: 0,
    price_max: 500,
  },


}
