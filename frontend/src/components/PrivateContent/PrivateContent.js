import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles, createStyles } from '@material-ui/core';
import { connect } from 'react-redux';
import { useRouter } from 'next/router';
import { bindActionCreators } from 'redux';
import { useTranslation } from 'src/i18n';
import { isAuthenticated as isAuthenticatedSelector } from 'src/redux/auth/auth.selectors';
import { getUser } from 'src/redux/user/user.selectors';
import Link from 'next/link';
import { ROUTES } from 'src/enums';

const useStyles = makeStyles(() => createStyles({
  root: {
    backgroundColor: '#e6f7ff',
    border: '1px solid #91d5ff',
    padding: '8px 15px',
    marginBottom: 16,
    borderRadius: 4,
  },
}));


function PrivateContent({
  children,
  isAuthenticated,
  checkContext,
  user,
  renderBlockMessage,
  showMessage,
}) {
  const classes = useStyles();
  const { t } = useTranslation('common');
  const router = useRouter();
  const showContent = isAuthenticated && (!checkContext || (user && checkContext(user)));
  if (showContent) return children;
  // User not able to see content
  if (!showMessage) return null;
  if (renderBlockMessage) return renderBlockMessage();
  if (checkContext && isAuthenticated) {
    return (
      <div className={`${classes.root} privateContent_component alert-info`}>
        <p>{t('privateContent.You_are_not_able_to_see_this_content')}</p>
      </div>
    );
  }
  return (
    <div className={`${classes.root} privateContent_component alert-info`}>
      <p>{t('privateContent.You_need_to_logged_in_to_see_this_content')}</p>
      <Link href={`${ROUTES.SIGNIN_ROUTE}?next=${router.pathname}`}><a>SIGNIN</a></Link>
    </div>
  );
}

function mapStateToProps(store) {
  return {
    isAuthenticated: isAuthenticatedSelector(store),
    user: getUser(store),

  };
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({}, dispatch),
  };
}
export default connect(mapStateToProps, mapDispatchToProps)(PrivateContent);
PrivateContent.defaultProps = {
  checkContext: null,
  renderBlockMessage: null,
  user: null,
  showMessage: true,
  isAuthenticated: false,
};
PrivateContent.propTypes = {
  checkContext: PropTypes.func,
  renderBlockMessage: PropTypes.func,
  showMessage: PropTypes.bool,
  isAuthenticated: PropTypes.bool,
  // eslint-disable-next-line react/forbid-prop-types
  children: PropTypes.any.isRequired,
  // eslint-disable-next-line react/forbid-prop-types
  user: PropTypes.object,
};
