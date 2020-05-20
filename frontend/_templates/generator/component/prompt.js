module.exports = [
  {
    name: 'name',
    type: 'input',
    message: 'Name of your component',
    required: true,
  },
  {
    name: 'type',
    type: 'select',
    message: 'Choose the type of your component',
    choices: ['Component', 'WithHook', 'Stateless'],
    required: true,
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
