local users = import './users.jsonnet';
{
  foo: {
    local sender = users.consumers,
    local receiver = users.merchants,
 
    id: 53478538,
    body: 'fooBodyMessage',
    sender_id: sender.foo.id,
    receiver_id: receiver.foo.id,
    created_at: '2010-01-01T15:04:05Z',
    received_at: '2010-01-02T15:05:05Z',
    seen_at: '2010-01-03T15:04:05Z',
  },

  bar: {
    local sender = users.consumers,
    local receiver = users.service_providers.medic,
  
    id: 4385239,
    body: 'barBodyMessage',
    sender_id: sender.foo.id,
    receiver_id: receiver.foo.id,
    created_at: '2020-01-02T15:04:05Z',
    received_at: '2020-01-02T15:05:05Z',
    seen_at: '2020-01-03T15:04:05Z',
  },
}
