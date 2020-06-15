local addresses = import './addresses.jsonnet';
local calendars = import './calendars.jsonnet';
local users = import './users.jsonnet';

{
  foo: {
    id: 1543934,
    user_id: users.notifier_providers.medic.foo.id,
    calendar_id: calendars.foo.medic.a.id,
    address_id: addresses.notifier_providers.medic.foo.id,

    title: 'foonotifierName',
    duration_in_minutes: 120,
    price_min: 20000,
    price_max: 100000,
  },

  bar: {
    id: 143934,
    user_id: users.notifier_providers.medic.bar.id,
    calendar_id: calendars.bar.medic.a.id,
    address_id: addresses.notifier_providers.medic.bar.id,

    title: 'barnotifierName',
    duration_in_minutes: 120,
    price_min: 20000,
    price_max: 100000,
  },
}
