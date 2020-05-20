import React from 'react';
import dynamic from 'next/dynamic';

const DynamicComponentWithNoSSR = dynamic(
  () => import('./Dashboard'),
  { ssr: false },
);

function Dashboard(props) {
  return (
    // eslint-disable-next-line react/jsx-props-no-spreading
    <DynamicComponentWithNoSSR {...props} />
  );
}

export default Dashboard;
