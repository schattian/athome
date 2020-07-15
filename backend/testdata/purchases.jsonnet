local addresses = import './addresses.jsonnet';
local products = import './products.jsonnet';
local shippings = import './shippings.jsonnet';
local users = import './users.jsonnet';

{
  foo: {
    id: 432,
    user_id: users.consumers.foo.id,
    distance_in_kilometers: 3,
    dest_address_id: addresses.consumers.foo.id,
    src_address_id: addresses.merchants.foo.id,
    shipping_id: shippings.foo.id,
    merchant_id: users.merchants.foo.id,
    items: { [std.toString(products.foo.a.id)]: 2, [std.toString(products.foo.b.id)]: 4 },
  },
}
