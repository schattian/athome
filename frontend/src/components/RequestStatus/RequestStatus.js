
// eslint-disable-next-line no-unused-vars
import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import {
  getIsLoadingByType, getErrorByType, getStatusByType, getIsEndByType,
} from 'src/redux/loaders';

function RequestStatus({
  isLoading,
  isEnd,
  status,
  error,
  children,
}) {
  return children({
    isLoading,
    isEnd,
    status,
    error,
  });
}

function mapStateToProps(store, props) {
  if (!props.loaderType) {
    throw new Error('src/components/RequestStatus missing loaderType');
  }
  return {
    isLoading: getIsLoadingByType(store, props.loaderType),
    isEnd: getIsEndByType(store, props.loaderType),
    status: getStatusByType(store, props.loaderType),
    error: getErrorByType(store, props.loaderType),
  };
}

export default connect(mapStateToProps, null)(RequestStatus);
RequestStatus.defaultProps = {
  isLoading: false,
  isEnd: false,
  error: null,
  status: null,
};

RequestStatus.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  loaderType: PropTypes.string.isRequired,
  isLoading: PropTypes.bool,
  isEnd: PropTypes.bool,
  // eslint-disable-next-line react/forbid-prop-types
  error: PropTypes.any,
  status: PropTypes.string,
};
