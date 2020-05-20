import React from 'react';
import Page from 'src/components/Page';
import Head from 'next/head';
import Screen from 'src/example/components/nextjs-material-kit-master/components';

export default Page({
  loginRequired: false,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['common'],
  showSwitchLangBtn: false,
  displayName: 'nextjs-material-kit-master',
})(() => (
  <React.Fragment>
    <Head>
      <title>Example Private Page</title>
    </Head>
    <h1>nextjs-material-kit-master</h1>
    <Screen />
  </React.Fragment>
));
