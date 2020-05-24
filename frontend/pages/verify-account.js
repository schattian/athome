import Head from 'next/head';
import React from 'react';
import Page from 'src/components/Page';
import Screen from 'src/layouts/verifyAccount';

export default Page({
  loginRequired: false,
  logoutRequired: true,
  adminRequired: false,
  i18n: ['verifyAccount'],
  showSwitchLangBtn: true,
  displayName: 'verifyAccountPage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Verify Account</title>
    </Head>
    <Screen />
  </React.Fragment>
));
