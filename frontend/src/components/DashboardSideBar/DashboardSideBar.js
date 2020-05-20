/* eslint-disable react/prop-types */
/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import Router from 'next/router';
import SidebarLink from 'src/components/DashboardLayout/Sidebar/SidebarLink';
import { withStyles, createStyles } from '@material-ui/core/styles';
import {
  TableChart as TableChartIcon,
} from '@material-ui/icons';

class DashboardSideBar extends React.Component {
  renderStaticLinks = () => {
    const { isSidebarOpened } = this.props;
    return (
      <React.Fragment>
        <SidebarLink
          key="dashboard"
          isSidebarOpened={isSidebarOpened}
          id="sales"
          label="Dashboard"
          link="/dashboard"
          icon={<TableChartIcon />}
          isLinkActive={Router.pathname === '/dashboard'}
        />
      </React.Fragment>
    );
  }


  render() {
    return (
      <React.Fragment>
        {this.renderStaticLinks()}
      </React.Fragment>
    );
  }
}


const styles = () => createStyles({
});

const Extended = withStyles(styles)(DashboardSideBar);

export default Extended;
