/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable react/require-default-props */
import React from 'react';
// nodejs library that concatenates classes
import clsx from 'clsx';
// nodejs library to set properties for components
import PropTypes from 'prop-types';

// material-ui components
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import getStyle from './buttonStyle';

const useStyles = makeStyles(getStyle);

export default function RegularButton(props) {
  const classes = useStyles();
  const {
    color,
    round,
    children,
    disabled,
    simple,
    size,
    block,
    link,
    justIcon,
    className,
    muiClasses,
    ...rest
  } = props;
  const btnClasses = clsx({
    [classes.button]: true,
    [classes[size]]: size,
    [classes[color]]: color,
    [classes.round]: round,
    [classes.disabled]: disabled,
    [classes.simple]: simple,
    [classes.block]: block,
    [classes.link]: link,
    [classes.justIcon]: justIcon,
    [className]: className,
  });
  return (
    <Button {...rest} classes={muiClasses} className={btnClasses}>
      {children}
    </Button>
  );
}

RegularButton.propTypes = {
  color: PropTypes.oneOf([
    'primary',
    'info',
    'success',
    'warning',
    'danger',
    'rose',
    'white',
    'transparent',
  ]),
  size: PropTypes.oneOf(['sm', 'lg']),
  simple: PropTypes.bool,
  round: PropTypes.bool,
  disabled: PropTypes.bool,
  block: PropTypes.bool,
  link: PropTypes.bool,
  justIcon: PropTypes.bool,
  className: PropTypes.string,
  // use this to pass the classes props from Material-UI
  // eslint-disable-next-line react/forbid-prop-types
  muiClasses: PropTypes.object,
  children: PropTypes.node,
};
