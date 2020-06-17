local messager = import './messager.jsonnet';

{
  foo: messager.foo { id: 4325 },
  bar: messager.foo { id: 342423 },
}
