local users = import './users.jsonnet';
{
  foo: {
    id: 234,
    user_id: users.consumers.foo.id,
    cvv_hash: '$2a$10$CDnJ/fppVsZi0XRLSvREj.lSFwIOPlDBcRJdEw9qJJTROoUIbWqCO',  // 123
    last_four_digits: 7500,
    expiry_month: 12,
    expiry_year: 43,
    holder_dni: 42418479,
    holder_name: 'Foo Holder Name',
  },
}
