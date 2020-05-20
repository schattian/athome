import React from 'react';
import Page from 'src/components/Page';
import Head from 'next/head';
import Screen from 'src/example/screens/exampleDashboardAccountScreen';

export default Page({
  loginRequired: true,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['common'],
  showSwitchLangBtn: true,
  displayName: 'exampleDashboardAccountScreen',
})(() => (
  <React.Fragment>
    <Head>
      <title>Example Dashboard Account Page</title>
    </Head>
    <Screen />
  </React.Fragment>
));
