local onboardings = import './onboardings.jsonnet';
{
  base: {
    foo: onboardings.base.foo { id: 4 },
    bar: onboardings.base.bar { id: 30 },
  },

  consumers:
    {
      local role = 'consumer',
      foo: $.base.foo { role: role },
    },

  service_providers:
    {
      local role = 'service-provider',
      foo: $.base.foo { role: role, category: "medic" },
    },

}
