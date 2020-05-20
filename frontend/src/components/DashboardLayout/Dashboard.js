import React from 'react';
import classnames from 'classnames';
import { i18n } from 'src/i18n';
// styles
import PropTypes from 'prop-types';
import useStyles from './styles';
// components
import Header from './Header';
import Sidebar from './Sidebar';

// context
import { useLayoutState } from './LayoutContext';

function Layout({ renderMain, renderSidebarBody, ...resProps }) {
  const classes = useStyles();

  // global
  const layoutState = useLayoutState();

  return (
    <div className={classes.root} dir={i18n.dir()}>
      <>
        <Header />
        <Sidebar
          renderBody={renderSidebarBody}
        />
        <div
          className={classnames(classes.content, {
            [classes.contentShift]: layoutState.isSidebarOpened,
          })}
        >
          <div className={classes.fakeToolbar} />
          {renderMain(resProps)}
        </div>
      </>
    </div>
  );
}


Layout.propTypes = {
  renderMain: PropTypes.func.isRequired,
  // eslint-disable-next-line react/forbid-prop-types
  renderSidebarBody: PropTypes.any.isRequired,
};

export default Layout;
