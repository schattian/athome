local services = import './services.jsonnet';

{
  foo: services.foo { id: 4325 },
  bar: services.foo { id: 342423 },
}
