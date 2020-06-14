local drafts = import './drafts.jsonnet';
local products = import './products.jsonnet';

{
  foo: {
    a: products.foo.a { user_id: '', id: 8329, draft_id: drafts.foo.id },
    b: products.foo { user_id: '', id: 5438, draft_id: drafts.foo.id },
  },
}
