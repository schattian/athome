import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Layout from 'src/layouts/signup';

export default Page({
  loginRequired: false,
  logoutRequired: true,
  adminRequired: false,
  i18n: ['signup'],
  showSwitchLangBtn: false,
  displayName: 'SignupPage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Sign up</title>
    </Head>
    <Layout />
  </React.Fragment>
));
