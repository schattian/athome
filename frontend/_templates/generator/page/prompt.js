module.exports = [
  {
    name: 'name',
    type: 'input',
    message: 'Name of your page',
    required: true,
  },
  {
    name: 'isPrivate',
    type: 'toggle',
    message: 'Is private screen ?',
  },
  {
    name: 'scss',
    type: 'toggle',
    message: 'Do you want to add a scss file?',
    default: false,
  },
  {
    name: 'withStyle',
    type: 'toggle',
    message: 'Do you want to add a material-ui withStyle?',
    default: true,
  },
  {
    name: 'redux',
    type: 'toggle',
    message: 'Do you want to connect redux?',
    default: false,
  },
];
