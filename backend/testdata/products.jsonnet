local images = import './images.jsonnet';
local categories = import './product_categories.jsonnet';
local users = import './users.jsonnet';

{

  foo: {
    a: {
      id: 2134512,
      user_id: users.merchants.foo.id,

      category_id: categories.foo.leaves.first.a.id,
      title: 'fooAProductTitle',

      price: 10000,
      stock: 23,

      image_ids: [images.foo.id],
    },
    b: {
      id: 432432432,
      user_id: users.merchants.foo.id,

      category_id: categories.foo.leaves.second.a.id,
      title: 'fooBProductTitle',

      price: 3,
      stock: 10,

      image_ids: [images.foo.id],
    },

  },


  bar: {
    a: {
      id: 23423312,
      user_id: users.merchants.bar.id,

      category_id: categories.bar.leaves.first.a.id,
      title: 'barAProductTitle',

      price: 33332,
      stock: 13,

      image_ids: [images.bar.id],
    },
    b: {
      id: 3424,
      user_id: users.merchants.bar.id,

      category_id: categories.bar.leaves.second.a.id,
      title: 'barBProductTitle',

      price: 324299,
      stock: 1000,

      image_ids: [images.bar.id],
    },

  },

}
