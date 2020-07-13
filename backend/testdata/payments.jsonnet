local cards = import './cards.jsonnet';
local products = import './products.jsonnet';
local purchases = import './purchases.jsonnet';
local services = import './services.jsonnet';
local shippings = import './shippings.jsonnet';
local users = import './users.jsonnet';
{
  purchases: {
    foo: {
      id: 32981,
      user_id: users.consumers.foo.id,
      payment_method_id: 1,
      card_id: cards.foo.id,
      entity_id: purchases.foo.id,
      entity_table: 'purchases',
      amount: shippings.foo.order_price +
              products.foo.a.price * 2 + products.foo.b.price * 4,
      installments: 12,
    },
  },
}
