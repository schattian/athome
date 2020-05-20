import React from 'react';
import Screen from 'src/screens/verifyAccount';
import Page from 'src/components/Page';
import Head from 'next/head';

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
