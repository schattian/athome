local onboardings = import './onboardings.jsonnet';
{
  consumers:
    {
      local role = 'consumer',
      foo: onboardings.consumers.foo { id: 4 },
    },

  service_providers:
    {
      local role = 'service-provider',
      medic: {
        foo: onboardings.service_providers.medic.foo { id: 30 },
      },
      lawyer: {
        foo: onboardings.service_providers.lawyer.foo { id: 45 },
      },
    },

}
