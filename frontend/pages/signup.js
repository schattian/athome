import React from 'react';
import Screen from 'src/screens/signup';
import Page from 'src/components/Page';
import Head from 'next/head';

export default Page({
  loginRequired: false,
  logoutRequired: true,
  adminRequired: false,
  i18n: ['signup'],
  showSwitchLangBtn: true,
  displayName: 'SignupPage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Sign up</title>
    </Head>
    <Screen />
  </React.Fragment>
));
