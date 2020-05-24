import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Screen from 'src/layouts/forgotPassword';

export default Page({
  loginRequired: false,
  logoutRequired: true,
  adminRequired: false,
  i18n: ['forgotPassword'],
  showSwitchLangBtn: true,
  displayName: 'forgotPassword',
})(() => (
  <React.Fragment>
    <Head>
      <title>Forgot Password</title>
    </Head>
    <Screen />
  </React.Fragment>
));
