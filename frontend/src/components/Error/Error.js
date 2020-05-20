
import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles, createStyles } from '@material-ui/core';
import { connect } from 'react-redux';
import { withTranslation } from 'src/i18n';
import { getErrorByType } from 'src/redux/loaders';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';


const useStyles = makeStyles(() => createStyles({
  fullScreen: {
    position: 'fixed',
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
    width: '100vw',
    height: '100vh',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#00000017',
    pointerEvents: 'none',
  },
  simple: {
    height: '55px',
    display: 'flex',
    padding: '0px 10px',
    alignItems: 'center',
    borderRadius: 4,
    backgroundColor: 'rgb(232, 240, 254) !important',
  },
}));

const getErrorMessage = (error, errorText, defaultErrorMessage) => {
  if (errorText) return errorText;
  if (error && typeof error === 'string') return error;
  if (error
    && error.response
    && error.response.data
    && error.response.data.message) return error.response.data.message;
  if (error && error.data && error.data.message) return error.data.message;
  if (error && error.message) return error.message;
  if (defaultErrorMessage) return defaultErrorMessage;
  return 'oops something went wrong';
};

function Error({
  errorType, error, fullScreen, errorText, t,
}) {
  const classes = useStyles();
  const showErr = (errorType && error);
  const className = fullScreen ? classes.fullScreen : classes.simple;
  const defaultErrorMessage = t('defaultErrorMessage');
  if (showErr) {
    return (
      <Box color="text.primary" className={`${className} error_component`}>
        <Typography color="error">{getErrorMessage(error, errorText, defaultErrorMessage)}</Typography>
      </Box>
    );
  }
  return null;
}

function mapStateToProps(store, props) {
  const propsToPass = {};
  if (props.errorType) {
    propsToPass.error = getErrorByType(store, props.errorType);
  }
  return propsToPass;
}

export default connect(mapStateToProps, null)(withTranslation('common')(Error));

Error.defaultProps = {
  errorText: null,
  error: null,
  fullScreen: false,
};

Error.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  error: PropTypes.any,
  errorText: PropTypes.string,
  errorType: PropTypes.string.isRequired,
  fullScreen: PropTypes.bool,
  t: PropTypes.func.isRequired,
};
