/* eslint-disable react/prop-types */
import React from 'react';
import { useTheme, makeStyles } from '@material-ui/styles';
import { Typography } from '../../../Wrappers';


const useStyles = makeStyles(() => ({
  avatar: {
    width: 30,
    height: 30,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: '50%',
  },
  text: {
    color: 'white',
  },
}));

// components

export default function UserAvatar({ color = 'primary', ...props }) {
  const classes = useStyles();
  const theme = useTheme();

  // eslint-disable-next-line react/destructuring-assignment
  const letters = props.name
    .split(' ')
    .map((word) => word[0])
    .join('');

  return (
    <div
      className={classes.avatar}
      style={{ backgroundColor: theme.palette[color].main }}
    >
      <Typography className={classes.text}>{letters}</Typography>
    </div>
  );
}
