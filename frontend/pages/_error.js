import React from 'react';
import Error from 'next/error';
import PropTypes from 'prop-types';

const ErrorPage = ({ errorCode }) => <Error statusCode={errorCode || 404} />;

ErrorPage.getInitialProps = async () => ({
  namespacesRequired: ['common'],
});

export default ErrorPage;

ErrorPage.defaultProps = {
  errorCode: 404,
};
ErrorPage.propTypes = {
  errorCode: PropTypes.number,
};
