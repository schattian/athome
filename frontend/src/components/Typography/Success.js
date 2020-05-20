import React from 'react';
import PropTypes from 'prop-types';
// @material-ui/core components
import { makeStyles } from '@material-ui/core/styles';
// core components
import getStyles from './typographyStyle';

const useStyles = makeStyles((theme) => getStyles(theme));

export default function Success(props) {
  const classes = useStyles();
  const { children } = props;
  return (
    <div className={`${classes.defaultFontStyle} ${classes.successText}`}>
      {children}
    </div>
  );
}

Success.propTypes = {
  // eslint-disable-next-line react/require-default-props
  children: PropTypes.node,
};
