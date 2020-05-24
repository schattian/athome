import React from 'react';
// import { useTranslation } from 'src/i18n';
import { makeStyles } from '@material-ui/core/styles';
import { LayoutProvider } from 'src/components/DashboardLayout/LayoutContext';
import DashboardLayout from 'src/components/DashboardLayout';
import DashboardSideBar from 'src/components/DashboardSideBar';
import Account from 'src/components/Account';

const useStyles = makeStyles({
  card: {
    minWidth: 275,
    margin: 15,
  },
});

function dashboardScreenScreen() {
  // const { t } = useTranslation('common');
  const classes = useStyles();
  return (
    <LayoutProvider>
      <DashboardLayout
        renderMain={() => (
          <div className={classes.card}>
            <Account />
          </div>
        )}
        renderSidebarBody={DashboardSideBar}
      />
    </LayoutProvider>
  );
}

export default dashboardScreenScreen;
