import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Layout from 'src/layouts/home';

export default Page({
  loginRequired: false,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['home', 'common'],
  showSwitchLangBtn: false,
  displayName: 'omePage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Home</title>
    </Head>
    <Layout />
  </React.Fragment>
));
