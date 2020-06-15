local onboardings = import './onboardings.jsonnet';
{
  base: {
    notifier_providers: {

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

  notifier_providers:
    {
      local role = 'notifier-provider',
      medic: {
        foo: $.base.notifier_providers.foo { id: 234, license: 123123, onboarding_id: onboardings.notifier_providers.medic.foo.id },
      },
      lawyer: {
        foo: $.base.notifier_providers.bar { id: 854,  tome: 123, folio: 456, onboarding_id: onboardings.notifier_providers.lawyer.foo.id },
      },
    },

}
