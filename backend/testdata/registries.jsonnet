local notifier = import './notifier.jsonnet';

{
  foo: notifier.foo { id: 4325 },
  bar: notifier.foo { id: 342423 },
}
