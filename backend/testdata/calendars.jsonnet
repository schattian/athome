local users = import './users.jsonnet';

{
  local groups = { foo: 2313, bar: 77723 },
  local medics = users.service_providers.medic,
  local lawyers = users.service_providers.lawyer,

  foo: {
    medic: {
      a: {
        id: 15934,
        user_id: medics.foo.id,
        name: 'fooMedicCalendarz',
        group_id: groups.foo,
      },
      b: {
        id: 4334,
        user_id: medics.foo.id,
        name: 'fooMedicCalendary',
        group_id: groups.foo,
      },
      c: {
        id: 5437134,
        user_id: medics.foo.id,
        name: 'fooMedicCalendarx',
        group_id: groups.foo,
      },
    },
    lawyer: {
      a: {
        id: 13554,
        user_id: lawyers.foo.id,
        name: 'fooLawyerCalendary',
        group_id: groups.foo,
      },
      b: {
        id: 43333334,
        user_id: lawyers.foo.id,
        name: 'fooLawyerCalendarz',
        group_id: groups.foo,
      },
      c: {
        id: 54399134,
        user_id: lawyers.foo.id,
        name: 'fooLawyerCalendarx',
        group_id: groups.foo,
      },
    },
  },

  bar: {
    medic: {
      a: {
        id: 1511934,
        user_id: medics.bar.id,
        name: 'barMedicCalendarz',
        group_id: groups.bar,
      },
      b: {
        id: 4335434,
        user_id: medics.bar.id,
        name: 'barMedicCalendary',
        group_id: groups.bar,
      },
      c: {
        id: 54334,
        user_id: medics.bar.id,
        name: 'barMedicCalendarx',
        group_id: groups.bar,
      },
    },
    lawyer: {
      a: {
        id: 115934,
        user_id: lawyers.bar.id,
        name: 'barLawyerCalendary',
        group_id: groups.bar,
      },
      b: {
        id: 4315134,
        user_id: lawyers.bar.id,
        name: 'barLawyerCalendarz',
        group_id: groups.bar,
      },
      c: {
        id: 5437134,
        user_id: lawyers.bar.id,
        name: 'barLawyerCalendarx',
        group_id: groups.bar,
      },
    },
  },
}
