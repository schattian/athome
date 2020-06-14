{
  foo: {
    id: 1111,
    name: 'fooMRootName',
    branches: {
      a: {
        id: 11438,
        name: 'fooMBranchNameA',
        parent_id: sp.ParentId,
        childs: {
          a: {
            id: 342,
            name: 'fooMBranchNameALeafNameA',
            parent_id: $.id,
          },
          b: {
            id: 3424,
            name: 'fooMBranchNameALeafNameB',
            parent_id: $.id,
          },
        },
      },
      b: {
        id: 20913,
        name: 'fooMBranchName',
        parent_id: sp.ParentId,
        childs: {
          a: {
            id: 9548,
            name: 'fooMBranchNameBLeafNameA',
            parent_id: $.id,
          },
          b: {
            id: 4532,
            name: 'fooMBranchNameBLeafNameB',
            parent_id: $.id,
          },
        },

      },
    },
  },
}
