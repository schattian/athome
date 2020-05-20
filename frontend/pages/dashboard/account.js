import React from 'react';
import Dashboard from 'src/screens/dashboardAccount';

import Page from 'src/components/Page';
import Head from 'next/head';

export default Page({
  loginRequired: true,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['dashboard', 'common'],
  showSwitchLangBtn: false,
  displayName: 'DashboardAccount',
})(() => (
  <React.Fragment>
    <Head>
      <title>Dashboard Account</title>
    </Head>
    <Dashboard />
  </React.Fragment>
));
