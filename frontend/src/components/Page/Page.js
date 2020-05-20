/* eslint-disable no-underscore-dangle */
// eslint-disable-next-line no-unused-vars
import React from 'react';
import WithAuth from './components/WithAuth';
import WithNotifier from './components/WithNotifier';
import WithBrowser from './components/WithBrowser';

const Page = (config) => (WrappedComponent) => {
  if (!config) throw new Error('Config is required');
  // eslint-disable-next-line no-unused-vars
  const {
    loginRequired, logoutRequired, adminRequired, i18n, showSwitchLangBtn, displayName,
  } = config;
  const _displayName = displayName || WrappedComponent.displayName || WrappedComponent.name || 'Component';

  const Extended1 = WithAuth(WrappedComponent, { loginRequired, logoutRequired, adminRequired });
  const Extended2 = WithNotifier(Extended1);
  const Extended3 = WithBrowser(Extended2, { showSwitchLangBtn, i18n });

  Extended3.displayName = _displayName;

  return Extended3;
};


export default Page;
