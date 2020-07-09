local calendars = import './calendars.jsonnet';

local base = {
  foo: {
    a: {
      day_of_week: 1,
      start_hour: 9,
      start_minute: 30,
      end_hour: 17,
      end_minute: 30,
    },
    b: {
      day_of_week: 2,
      start_hour: 1,
      start_minute: 30,
      end_hour: 18,
      end_minute: 30,
    },
    c: {
      day_of_week: 5,
      start_hour: 13,
      start_minute: 30,
      end_hour: 12,
      end_minute: 30,
    },
  },
};

{

  foo: {
    medic: {
      first: {
        local calendar = calendars.foo.medic.a,
        a: base.foo.a {
          id: 324,
          calendar_id: calendar.id,
        },
        b: base.foo.b {
          id: 94,
          calendar_id: calendar.id,
        },
        c: base.foo.c {
          id: 223424,
          calendar_id: calendar.id,
        },
      },
    },

    delivery: {
      first: {
        local calendar = calendars.foo.delivery.a,
        a: base.foo.a {
          id: 3214,
          calendar_id: calendar.id,
        },
        b: base.foo.b {
          id: 121,
          calendar_id: calendar.id,
        },
        c: base.foo.c {
          id: 2234,
          calendar_id: calendar.id,
        },
      },
    },

  },


}
