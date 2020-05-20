import React from 'react';
import Screen from 'src/screens/home';
import Page from 'src/components/Page';
import Head from 'next/head';

export default Page({
  loginRequired: false,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['home', 'common'],
  showSwitchLangBtn: true,
  displayName: 'omePage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Home</title>
    </Head>
    <Screen />
  </React.Fragment>
));
