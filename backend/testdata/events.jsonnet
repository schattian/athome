local availabilities = import './availabilities.jsonnet';
local users = import './users.jsonnet';

{
  foo: {

    medic: {
      first: {
        local avs = availabilities.foo.medic.first,
        a: avs.a {
          id: 8923,
          claimant_id: users.consumers.foo.id,
          start_hour: avs.a.start_hour + 1,
          end_hour: avs.a.end_hour - 2,
        },
        b: avs.b {
          id: 94423,
          claimant_id: users.consumers.foo.id,
          start_hour: avs.a.start_hour + 2,
          end_hour: avs.a.end_hour - 3,
          end_minute: avs.a.end_minute - 30,
        },
        c: avs.c {
          id: 2124,
          claimant_id: users.consumers.foo.id,
          start_hour: avs.a.start_hour + 2,
          end_hour: avs.a.end_hour - 3,
        },
      },
    },

    delivery: {
      first: {
        local avs = availabilities.foo.delivery.first,
        a: avs.a {
          id: 1211,
          claimant_id: users.consumers.foo.id,
          start_hour: avs.a.start_hour + 1,
          end_hour: avs.a.end_hour - 2,
        },
        b: avs.b {
          id: 111,
          claimant_id: users.consumers.foo.id,
          start_hour: avs.a.start_hour + 2,
          end_hour: avs.a.end_hour - 3,
          end_minute: avs.a.end_minute - 30,
        },
        c: avs.c {
          id: 99,
          claimant_id: users.consumers.foo.id,
          start_hour: avs.a.start_hour + 2,
          end_hour: avs.a.end_hour - 3,
        },
      },
    },


  },


}
