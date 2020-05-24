import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Dashboard from 'src/layouts/dashboardAccount';

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
