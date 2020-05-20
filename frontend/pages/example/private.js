import React from 'react';
import Page from 'src/components/Page';
import Head from 'next/head';
import Screen from 'src/example/screens/exampleScreen';

export default Page({
  loginRequired: true,
  logoutRequired: false,
  adminRequired: false,
  i18n: ['common'],
  showSwitchLangBtn: true,
  displayName: 'ExamplePrivatePage',
})(() => (
  <React.Fragment>
    <Head>
      <title>Example Private Page</title>
    </Head>
    <h1>Example Private Page</h1>
    <Screen />
  </React.Fragment>
));
