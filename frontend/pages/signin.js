import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Screen from 'src/layouts/signin';

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
