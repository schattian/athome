import React from 'react';
import Screen from 'src/screens/forgotPassword';
import Page from 'src/components/Page';
import Head from 'next/head';

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
