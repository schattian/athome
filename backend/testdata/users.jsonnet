local onboardings = import './onboardings.jsonnet';
{
  consumers:
    {
      local role = 'consumer',
      foo: onboardings.consumers.foo { id: 4 },
    },

  merchants:
    {
      local role = 'merchant',
      foo: onboardings.merchants.foo { id: 3446 },
      bar: onboardings.merchants.bar { id: 3426 },
    },


  service_providers:
    {
      local role = 'service-provider',
      medic: {
        foo: onboardings.service_providers.medic.foo { id: 30 },
        bar: onboardings.service_providers.medic.bar { id: 39 },
      },
      lawyer: {
        foo: onboardings.service_providers.lawyer.foo { id: 45 },
        bar: onboardings.service_providers.lawyer.bar { id: 48 },
      },
    },

}
