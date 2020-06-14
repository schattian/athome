{
  foo: {
    root: {
      id: 9123,
      name: 'fooPRootName',
    },

    leaves: {
      first: {
        local parent = $.foo.branches.a,
        a: {
          id: 8534,
          name: 'fooPBranchNameALeafNameA',
          parent_id: parent.id,
        },
        b: {
          id: 8533424,
          name: 'fooPBranchNameALeafNameB',
          parent_id: parent.id,
        },
      },
      second: {
        local parent = $.foo.branches.b,
        a: {
          id: 855494,
          name: 'fooPBranchNameBLeafNameA',
          parent_id: parent.id,
        },
        b: {
          id: 38533424,
          name: 'fooPBranchNameBLeafNameB',
          parent_id: parent.id,
        },
      },
    },

    branches: {
      local parent = $.foo.root,
      a: {
        id: 5438,
        name: 'fooPBranchNameA',
        parent_id: parent.id,
      },
      b: {
        id: 213,
        name: 'fooPBranchName',
        parent_id: parent.id,
      },
    },
  },

  bar: {
    root: {
      id: 912321231,
      name: 'barPRootName',
    },

    leaves: {
      first: {
        local parent = $.bar.branches.a,
        a: {
          id: 8534211,
          name: 'barPBranchNameALeafNameA',
          parent_id: parent.id,
        },
        b: {
          id: 424,
          name: 'barPBranchNameALeafNameB',
          parent_id: parent.id,
        },
      },
      second: {
        local parent = $.bar.branches.b,
        a: {
          id: 852112194,
          name: 'barPBranchNameBLeafNameA',
          parent_id: parent.id,
        },
        b: {
          id: 39329924,
          name: 'barPBranchNameBLeafNameB',
          parent_id: parent.id,
        },
      },
    },

    branches: {
      local parent = $.bar.root,
      a: {
        id: 1115438,
        name: 'barPBranchNameA',
        parent_id: parent.id,
      },
      b: {
        id: 213111111,
        name: 'barPBranchName',
        parent_id: parent.id,
      },
    },
  },


}
