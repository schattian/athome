local onboardings = import './onboardings.jsonnet';
{
  consumers:
    {
      local role = 'consumer',
      foo: onboardings.consumers.foo { id: 4 },
      bar: onboardings.consumers.bar { id: 214 },
    },

  merchants:
    {
      local role = 'merchant',
      foo: onboardings.merchants.foo { id: 3446 },
      bar: onboardings.merchants.bar { id: 3426 },
    },


  service_providers:
    {
      local role = 'service_provider',
      medic: {
        foo: onboardings.service_providers.medic.foo { id: 30 },
        bar: onboardings.service_providers.medic.bar { id: 39 },
      },
      delivery: {
        foo: onboardings.service_providers.delivery.foo { id: 34 },
        bar: onboardings.service_providers.delivery.bar { id: 24 },
      },

      lawyer: {
        foo: onboardings.service_providers.lawyer.foo { id: 45 },
        bar: onboardings.service_providers.lawyer.bar { id: 48 },
      },
    },

}
