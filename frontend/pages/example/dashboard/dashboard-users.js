import React from 'react';
import Page from 'src/components/Page';
import Head from 'next/head';
import Screen from 'src/example/screens/exampleDashboardUsersScreen';

export default Page({
  loginRequired: true,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['common'],
  showSwitchLangBtn: true,
  displayName: 'ExampleDashboardPage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Example Dashboard Users Page</title>
    </Head>
    <Screen />
  </React.Fragment>
));
