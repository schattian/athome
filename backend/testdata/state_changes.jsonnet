local payments = import './payments.jsonnet';
local purchases = import './purchases.jsonnet';
local shippings = import './shippings.jsonnet';
{
  payments: {
    cancelled: {
      id: 3232,
      name: 'cancelled',
      stage: -1,
      entity_id: payments.purchases.foo.id,
      entity_table: 'payments',
    },
    created: {
      id: 329,
      name: 'created',
      stage: 1,
      entity_id: payments.purchases.foo.id,
      entity_table: 'payments',
    },
    finished: {
      id: 39399,
      name: 'finished',
      stage: 2,
      entity_id: payments.purchases.foo.id,
      entity_table: 'payments',
    },
    rejected: {
      id: 18282,
      name: 'rejected',
      stage: -2,
      entity_id: payments.purchases.foo.id,
      entity_table: 'payments',
    },
  },
  shippings: {
    cancelled: {
      id: 3824832,
      name: 'cancelled',
      stage: -1,
      entity_id: shippings.foo.id,
      entity_table: 'shippings',
    },
    created: {
      id: 382222,
      name: 'created',
      stage: 1,
      entity_id: shippings.foo.id,
      entity_table: 'shippings',
    },
    taken: {
      id: 3993,
      name: 'taken',
      stage: 2,
      entity_id: shippings.foo.id,
      entity_table: 'shippings',
    },
    finished: {
      id: 99932,
      name: 'finished',
      stage: 3,
      entity_id: shippings.foo.id,
      entity_table: 'shippings',
    },
  },
  purchases: {
    cancelled: {
      id: 9329,
      name: 'cancelled',
      stage: -1,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
    created: {
      id: 234,
      name: 'purchase:created',
      stage: 1,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
    addressed: {
      id: 134,
      name: 'purchase:addressed',
      stage: 2,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
    shipping_method_selected: {
      id: 214,
      name: 'purchase:shipping_method_selected',
      stage: 3,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
    paid: {
      id: 1114,
      name: 'purchase:paid',
      stage: 4,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
    confirmed: {
      id: 2341119,
      name: 'purchase:confirmed',
      stage: 5,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
    finished: {
      id: 328,
      name: 'purchase:finished',
      stage: 6,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
    },
  },
}
