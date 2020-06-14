local users = import './users.jsonnet';

{
  foo: {
    id: 54385834,
    user_id: users.merchants.foo.id,
  },
}
