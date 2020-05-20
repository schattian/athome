import React from 'react';
import Screen from 'src/screens/signin';
import Page from 'src/components/Page';
import Head from 'next/head';

export default Page({
  loginRequired: false,
  logoutRequired: true,
  adminRequired: false,
  i18n: ['signin'],
  showSwitchLangBtn: true,
  displayName: 'SigninPage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Sign in</title>
    </Head>
    <Screen />
  </React.Fragment>
));
