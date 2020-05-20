import React from 'react';
import PropTypes from 'prop-types';
import SidebarLink from 'src/components/DashboardLayout/Sidebar/SidebarLink';
import {
  TableChart as TableChartIcon,
} from '@material-ui/icons';
import { useRouter } from 'next/router';

function DashboardSideBar({ isSidebarOpened }) {
  const router = useRouter();
  return (
    <React.Fragment>
      <SidebarLink
        key="Posts"
        id="Posts"
        label="Posts"
        link="/example/dashboard/dashboard-posts"
        isSidebarOpened={isSidebarOpened}
        icon={<TableChartIcon />}
        isLinkActive={router.pathname === '/example/dashboard/dashboard-posts'}
      />
      <SidebarLink
        key="Users"
        id="Users"
        label="Users"
        link="/example/dashboard/dashboard-users"
        isSidebarOpened={isSidebarOpened}
        icon={<TableChartIcon />}
        isLinkActive={router.pathname === '/example/dashboard/dashboard-users'}
      />
      <SidebarLink
        key="Account"
        id="Account"
        label="Account"
        link="/example/dashboard/dashboard-account"
        isSidebarOpened={isSidebarOpened}
        icon={<TableChartIcon />}
        isLinkActive={router.pathname === '/example/dashboard/dashboard-users'}
      />
    </React.Fragment>
  );
}

export default DashboardSideBar;

DashboardSideBar.propTypes = {
  isSidebarOpened: PropTypes.bool.isRequired,
};
