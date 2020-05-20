
import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles, createStyles } from '@material-ui/core';
import Avatar from '@material-ui/core/Avatar';
import BlurCircular from '@material-ui/icons/BlurCircular';
import { title } from '../../../siteConfig';

const useStyles = makeStyles((theme) => createStyles({
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
}));

function Logo({ appName }) {
  const classes = useStyles();
  if (appName) {
    return (
      <p>
        {title}
      </p>
    );
  }
  return (
    <Avatar className={classes.avatar}>
      <BlurCircular />
    </Avatar>
  );
}

export default (Logo);

Logo.defaultProps = {
  appName: false,
};
Logo.propTypes = {
  appName: PropTypes.bool,
};
