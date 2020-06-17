local users = import './users.jsonnet';
{
  local max = 4,
  local low = 1,
  local mid = 2,
  local high = 3,
  consumers:
    {
      local user = users.consumers,
      foo: {
        id: 43258,
        priority: high,
        body: "fooNotificationBody",
        user_id: user.foo.id,
        entity_table: 'orders',
        entity_id: 432,
        created_at: '2010-01-01T15:04:05Z',
        received_at: '2010-01-02T15:05:05Z',
        seen_at: '2010-01-03T15:04:05Z',
      },
      bar: {
        id: 324,
        priority: mid,
        body: "barNotificationBody",
        user_id: user.bar.id,
        entity_table: 'orders',
        entity_id: 3322,
        created_at: '2006-01-02T15:04:05Z',
        received_at: '2006-01-02T15:05:05Z',
        seen_at: '2006-01-03T15:04:05Z',
      },

    },
  //   merchants:
  //     {
  //       local role = 'merchant',
  //       foo: onboardings.merchants.foo { id: 3446 },
  //       bar: onboardings.merchants.bar { id: 3426 },
  //     },


  //   service_providers:
  //     {
  //       local role = 'service_provider',
  //       medic: {
  //         foo: onboardings.service_providers.medic.foo { id: 30 },
  //         bar: onboardings.service_providers.medic.bar { id: 39 },
  //       },
  //       lawyer: {
  //         foo: onboardings.service_providers.lawyer.foo { id: 45 },
  //         bar: onboardings.service_providers.lawyer.bar { id: 48 },
  //       },
  //     },

}
