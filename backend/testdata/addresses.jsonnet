local users = import './users.jsonnet';

{
  consumers:
    {
      local user = users.consumers,
      foo: {
        id: 83284,
        user_id: user.foo.id,
        country: 'ARGENTINA',
        province: 'BUENOS AIRES',
        zipcode: '2700',
        street: 'Alsina',
        number: 3,
        floor: 7,
        department: 'A',
        latitude: 3245344,
        longitude: 325435,
        alias: 'casa',
      },
    },
  merchants:
    {
      local users = user.merchants,
      foo: {
        id: 84,
        user_id: user.foo.id,
        country: 'ARGENTINA',
        province: 'BUENOS AIRES',
        zipcode: '700',
        street: 'Ala',
        number: 2,
        floor: 2,
        department: '',
        latitude: 35344,
        longitude: 32535,
        alias: 'casaej',
      },
    },
  service_providers:
    {
      medic: {
        local user = users.service_providers.medic,
        foo: {
          id: 22284,
          user_id: user.foo.id,
          country: 'PARAGUAY',
          province: 'UENOS AIRES',
          zipcode: '99700',
          street: 'La la',
          number: 23,
          floor: 0,
          department: '',
          latitude: 25344,
          longitude: 22535,
          alias: 'casita',
        },
      },
      lawyer: {
        local user = users.service_providers.lawyer,
        foo: {
          id: 3484,
          user_id: user.foo.id,
          country: 'URUGUAY',
          province: 'AIRES',
          zipcode: '100',
          street: 'sdaa',
          number: 13243,
          floor: 1,
          department: '',
          latitude: 131313,
          longitude: 313131,
          alias: 'caseta',
        },
      },
    },
}
