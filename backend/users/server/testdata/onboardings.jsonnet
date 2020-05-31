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

  service_providers:
    {
      local role = 'service-provider',
      foo: $.base.foo { id: 241, role: role, category: "medic" },
    },
}
