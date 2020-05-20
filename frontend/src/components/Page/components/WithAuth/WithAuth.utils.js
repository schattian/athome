
import React from 'react';
import Router from 'next/router';
import { ROUTES } from 'src/enums';
import logger from 'src/services/logger';
import Button from '@material-ui/core/Button';
import { reAuthenticate } from 'src/redux/auth/auth.actions';

export const onRedirectToSignIn = () => {
  setTimeout(() => {
    Router.push(`${ROUTES.SIGNIN_ROUTE}?next=${Router.router.pathname}`);
    /*
     This delay is important,
     for some reason Router.push is not always navigate without this delay
    */
  }, 1000);
  logger.debug('User is not authenticate, navigate to signin page');
};
export const onRedirectToHome = () => {
  setTimeout(() => {
    Router.push('/');
    /*
     This delay is important,
     for some reason Router.push is not always navigate without this delay
    */
  }, 1000);
  logger.debug('User is authenticate, navigate home page');
};

export const DisplayLoader = () => 'Loading...';
export const DisplayRedirect = () => 'Redirect...';

const OFF_LINE_STYLE = {
  margin: 15,
  padding: '8px 15px',
  border: '1px solid #91d5ff',
};

// eslint-disable-next-line react/prop-types
export const DisplayOffLine = ({ dispatch, tokenFromStorage }) => (
  <div style={OFF_LINE_STYLE}>
    <p>Sorry, you are off line
      This page is available for viewing only with a network connection
    </p>
    <Button onClick={() => reAuthenticate(dispatch, tokenFromStorage)}>Try again</Button>
  </div>
);
