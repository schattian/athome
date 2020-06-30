local messenger = import './messenger.jsonnet';

{
  foo: messenger.foo { id: 4325 },
  bar: messenger.foo { id: 342423 },
}
