local events = import './events.jsonnet';
local purchases = import './purchases.jsonnet';
local services = import './services.jsonnet';
local users = import './users.jsonnet';
{
  foo: {
    local service = services.delivery,
    local deliverer = users.service_providers.delivery.foo,
    local event = events.foo.delivery.first.a,
    local purchase = purchases.foo,

    id: 432213,
    user_id: deliverer.id,
    event_id: event.id,
    shipping_method_id: service.id,

    order_price: ((service.price_max + service.price_min) / 2) * (purchase.distance_in_kilometers),
    order_duration_in_minutes:  purchase.distance_in_kilometers * service.duration_in_minutes,

    // real_price: ,
    // real_duration_in_minutes: u.RealDurationInMinutes,
  },
}
