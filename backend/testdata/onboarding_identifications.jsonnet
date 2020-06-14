local onboardings = import './onboardings.jsonnet';
{
  base: {
    service_providers: {

      foo: {
        dni: 42418479,
        name: 'fooOnboardingIdName',
        surname: 'fooOnboardingIdSurname',
      },

      bar: {
        dni: 41124535,
        name: 'barOnboardingIdName',
        surname: 'barOnboardingIdSurname',
      },
    },
  },

  service_providers:
    {
      local role = 'service-provider',
      medic: {
        foo: $.base.service_providers.foo { id: 234, license: 123123, onboarding_id: onboardings.service_providers.medic.foo.id },
      },
      lawyer: {
        foo: $.base.service_providers.bar { id: 854,  tome: 123, folio: 456, onboarding_id: onboardings.service_providers.lawyer.foo.id },
      },
    },

}
