{
  base: {
    foo: {
      email: 'foo@foodomain.fooext',
      name: 'fooName',
      surname: 'fooSurname',
    },
    bar: {
      email: 'bar@bardomain.barext',
      name: 'barName',
      surname: 'barSurname',
    },
  },

  consumers:
    {
      local role = 'consumer',
      foo: $.base.foo { id: 324, role: role },
    },

  merchants:
    {
      local role = 'merchant',
      foo: $.base.foo { id: 994, role: role, category_id: 8 },
      bar: $.base.bar { id: 323254, role: role, category_id: 3 },
    },


  notifier_providers:
    {
      local role = 'notifier-provider',
      medic:
        {
          foo: $.base.foo { id: 241, role: role, category_id: 4 },
          bar: $.base.bar { id: 231, role: role, category_id: 4 },
        },

      lawyer:
        {
          foo: $.base.foo { id: 341, role: role, category_id: 5 },
          bar: $.base.bar { id: 311, role: role, category_id: 5 },
        },
    },
}
