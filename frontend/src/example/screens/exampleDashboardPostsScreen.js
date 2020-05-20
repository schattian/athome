import React from 'react';
import { useTranslation } from 'src/i18n';
import { LayoutProvider } from 'src/components/DashboardLayout/LayoutContext';
import DashboardLayout from 'src/components/DashboardLayout';
import DashboardSideBar from '../components/DashboardSideBar';

function exampleDashboardScreenScreen() {
  const { t } = useTranslation('common');
  return (
    <LayoutProvider>
      <DashboardLayout
        renderMain={() => (
          <div>
            <p>{t('screenName')}</p>
            Here is all the posts
          </div>
        )}
        renderSidebarBody={DashboardSideBar}
      />
    </LayoutProvider>
  );
}

export default exampleDashboardScreenScreen;
