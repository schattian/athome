import React from 'react';
import Button from 'src/components/Button';
import Favorite from '@material-ui/icons/Favorite';

const classes = {
  icons: {
    width: '17px',
    height: '17px',
    color: '#FFFFFF',
  },
};

const ButtonExample = () => (
  <div id="buttons">
    <p>style</p>
    <Button color="primary">Default</Button>
    <Button color="primary" round>
      round
    </Button>
    <Button color="primary" round>
      <Favorite className={classes.icons} /> with icon
    </Button>
    <Button justIcon round color="primary">
      <Favorite className={classes.icons} />
    </Button>
    <Button color="primary" simple>
      simple
    </Button>
    <p>size</p>
    <Button color="primary" size="sm">
      Small
    </Button>
    <Button color="primary">Regular</Button>
    <Button color="primary" size="lg">
      Large
    </Button>
    <p>color</p>
    <Button>Default</Button>
    <Button color="primary">Primary</Button>
    <Button color="info">Info</Button>
    <Button color="success">Success</Button>
    <Button color="warning">Warning</Button>
    <Button color="danger">Danger</Button>
    <Button color="rose">Rose</Button>
  </div>
);

export default ButtonExample;
