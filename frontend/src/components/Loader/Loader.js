
import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles, createStyles } from '@material-ui/core';
import CircularProgress from '@material-ui/core/CircularProgress';
import { connect } from 'react-redux';
import { getIsLoadingByType } from 'src/redux/loaders';


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
  },
  simple: {
    padding: 5,
  },
}));

function Loader({
  isLoading, loaderType, isLoadingByType, fullScreen,
}) {
  const classes = useStyles();
  const showLoader = isLoading || (loaderType && isLoadingByType);
  const className = fullScreen ? classes.fullScreen : classes.simple;
  if (showLoader) {
    return (
      <div className={`${className} loader_component`}>
        <CircularProgress />
      </div>
    );
  }
  return null;
}

function mapStateToProps(store, props) {
  const propsToPass = {};
  if (props.loaderType) {
    propsToPass.isLoadingByType = getIsLoadingByType(store, props.loaderType);
  }
  return propsToPass;
}

export default connect(mapStateToProps, null)(Loader);
Loader.defaultProps = {
  isLoading: false,
  isLoadingByType: false,
  loaderType: null,
  fullScreen: false,
};

Loader.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  isLoading: PropTypes.bool,
  isLoadingByType: PropTypes.bool,
  loaderType: PropTypes.string,
  fullScreen: PropTypes.bool,
};
