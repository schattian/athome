import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
// @material-ui/core components
import getStyles from './typographyStyle';

const useStyles = makeStyles((theme) => getStyles(theme));


export default function Info(props) {
  const classes = useStyles();
  const { children } = props;
  return (
    <div className={`${classes.defaultFontStyle} ${classes.infoText}`}>
      {children}
    </div>
  );
}

Info.propTypes = {
  // eslint-disable-next-line react/require-default-props
  children: PropTypes.node,
};
