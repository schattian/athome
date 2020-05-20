module.exports = [
  {
    name: 'name',
    type: 'input',
    message: 'Name of your reducer',
    required: true,
  },
  {
    name: 'example',
    type: 'toggle',
    message: 'You want an example of action in your code ?',
    initial: true,
  },
  {
    name: 'saga',
    type: 'toggle',
    message: 'Do you need saga ?',
    initial: true,
  },
];
