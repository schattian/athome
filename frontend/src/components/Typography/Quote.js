import React from 'react';
import PropTypes from 'prop-types';
// @material-ui/core components
import { makeStyles } from '@material-ui/core/styles';
// @material-ui/core components
import getStyles from './typographyStyle';

const useStyles = makeStyles((theme) => getStyles(theme));

export default function Quote(props) {
  const classes = useStyles();
  const { text, author } = props;
  return (
    <blockquote className={`${classes.defaultFontStyle} ${classes.quote}`}>
      <p className={classes.quoteText}>{text}</p>
      <small className={classes.quoteAuthor}>{author}</small>
    </blockquote>
  );
}

Quote.propTypes = {
  // eslint-disable-next-line react/require-default-props
  text: PropTypes.node,
  // eslint-disable-next-line react/require-default-props
  author: PropTypes.node,
};
