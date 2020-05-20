import { makeStyles } from '@material-ui/core/styles';
import React from 'react';
import DashboardLayout from 'src/components/DashboardLayout';
import { LayoutProvider } from 'src/components/DashboardLayout/LayoutContext';
import DashboardSideBar from 'src/components/DashboardSideBar';

const useStyles = makeStyles({
  card: {
    minWidth: 275,
    margin: 15,
  },
});

function dashboardScreenScreen() {
  const classes = useStyles();
  return (
    <LayoutProvider>
      <DashboardLayout
        renderMain={() => (
          <div className={classes.card}>
            <h1>Hello</h1>
          </div>
        )}
        renderSidebarBody={DashboardSideBar}
      />
    </LayoutProvider>
  );
}

export default dashboardScreenScreen;
