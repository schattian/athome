import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Dashboard from 'src/layouts/dashboard';

export default Page({
  loginRequired: true,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['dashboard', 'common'],
  showSwitchLangBtn: false,
  displayName: 'DashboardPage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Dashboard</title>
    </Head>
    <Dashboard />
  </React.Fragment>
));
