{
  foo: {
    id: 21311,
    name: 'fooSPRootName',
    parent_id: sp.ParentId,
    branches: {
      a: {
        id: 5348,
        name: 'fooSPBranchNameA',
        parent_id: sp.ParentId,
        childs: {
          a: {
            id: 113429,
            name: 'fooSPBranchNameALeafNameA',
            parent_id: $.id,
          },
          b: {
            id: 342498,
            name: 'fooSPBranchNameALeafNameB',
            parent_id: $.id,
          },
        },
      },
      b: {
        id: 984,
        name: 'fooSPBranchName',
        parent_id: sp.ParentId,
        childs: {
          a: {
            id: 54324328,
            name: 'fooSPBranchNameBLeafNameA',
            parent_id: $.id,
          },
          b: {
            id: 43424332,
            name: 'fooSPBranchNameBLeafNameB',
            parent_id: $.id,
          },
        },

      },
    },
  },
}
