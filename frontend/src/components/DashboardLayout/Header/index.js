/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
import React, { useState } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import PropTypes from 'prop-types';
import {
  AppBar,
  Toolbar,
  IconButton,
  // InputBase,
  Menu,
  MenuItem,
  // Fab,
} from '@material-ui/core';
import {
  Menu as MenuIcon,
  MailOutline as MailIcon,
  NotificationsNone as NotificationsIcon,
  Person as AccountIcon,
  // Search as SearchIcon,
  // Send as SendIcon,
  ArrowBack as ArrowBackIcon,
} from '@material-ui/icons';
import classNames from 'classnames';
import SwitchLang from 'src/components/SwitchLang';

// styles
import { logout } from 'src/redux/auth/auth.actions';
import { getUser } from 'src/redux/user/user.selectors';
import { useRouter } from 'next/router';
import { useTranslation } from 'src/i18n';
import Logo from 'src/components/Logo';
import useStyles from './styles';

// components
import { Badge, Typography } from '../Wrappers';
import Notification from './components/Notification';
import UserAvatar from './components/UserAvatar';

// context
import {
  useLayoutState,
  useLayoutDispatch,
  toggleSidebar,
} from '../LayoutContext';

const messages = [
  // {
  //   id: 0,
  //   variant: 'warning',
  //   name: 'Jane Hew',
  //   message: 'Hey! How is it going?',
  //   time: '9:32',
  // },
  // {
  //   id: 1,
  //   variant: 'success',
  //   name: 'Lloyd Brown',
  //   message: 'Check out my new Dashboard',
  //   time: '9:18',
  // },
  // {
  //   id: 2,
  //   variant: 'primary',
  //   name: 'Mark Winstein',
  //   message: 'I want rearrange the appointment',
  //   time: '9:15',
  // },
  // {
  //   id: 3,
  //   variant: 'secondary',
  //   name: 'Liana Dutti',
  //   message: 'Good news from sale department',
  //   time: '9:09',
  // },
];

const notifications = [
  // { id: 0, color: 'warning', message: 'Check out this awesome ticket' },
  // {
  //   id: 1,
  //   color: 'success',
  //   type: 'info',
  //   message: 'What is the best way to get ...',
  // },
  // {
  //   id: 2,
  //   color: 'secondary',
  //   type: 'notification',
  //   message: 'This is just a simple notification',
  // },
  // {
  //   id: 3,
  //   color: 'primary',
  //   type: 'e-commerce',
  //   message: '12 new orders has arrived today',
  // },
];

function Header({ user = {}, actions }) {
  const userData = user || {};

  const classes = useStyles();
  const router = useRouter();
  const { t } = useTranslation('common');
  // global
  const layoutState = useLayoutState();
  const layoutDispatch = useLayoutDispatch();

  // local
  const [mailMenu, setMailMenu] = useState(null);
  const [isMailsUnread, setIsMailsUnread] = useState(true);
  const [notificationsMenu, setNotificationsMenu] = useState(null);
  const [isNotificationsUnread, setIsNotificationsUnread] = useState(true);
  const [profileMenu, setProfileMenu] = useState(null);
  // const [isSearchOpen, setSearchOpen] = useState(false);

  return (
    <AppBar position="fixed" className={classes.appBar}>
      <Toolbar className={classes.toolbar}>
        <IconButton
          color="inherit"
          onClick={() => toggleSidebar(layoutDispatch)}
          className={classNames(
            classes.headerMenuButton,
            classes.headerMenuButtonCollapse,
          )}
        >
          {layoutState.isSidebarOpened ? (
            <ArrowBackIcon
              classes={{
                root: classNames(
                  classes.headerIcon,
                  classes.headerIconCollapse,
                ),
              }}
            />
          ) : (
            <MenuIcon
              classes={{
                root: classNames(
                  classes.headerIcon,
                  classes.headerIconCollapse,
                ),
              }}
            />
          )}
        </IconButton>
        <Typography variant="h6" weight="medium" className={classes.logotype}>
          {t('dashboardLayout.title')}
        </Typography>
        <div className={classes.grow}>
          <Logo appName />
        </div>
        {/*
          <div className={classes.grow} />
        <div
          className={classNames(classes.search, {
            [classes.searchFocused]: isSearchOpen,
          })}
        >
          <div
            className={classNames(classes.searchIcon, {
              [classes.searchIconOpened]: isSearchOpen,
            })}
            onClick={() => setSearchOpen(!isSearchOpen)}
          >
            <SearchIcon classes={{ root: classes.headerIcon }} />
          </div>
          <InputBase
            placeholder="Search…"
            classes={{
              root: classes.inputRoot,
              input: classes.inputInput,
            }}
          />
          </div>
        */}
        <IconButton
          color="inherit"
          aria-haspopup="true"
          aria-controls="mail-menu"
          onClick={(e) => {
            setNotificationsMenu(e.currentTarget);
            setIsNotificationsUnread(false);
          }}
          className={classes.headerMenuButton}
        >
          <Badge
            badgeContent={isNotificationsUnread ? notifications.length : null}
            color="warning"
          >
            <NotificationsIcon classes={{ root: classes.headerIcon }} />
          </Badge>
        </IconButton>
        <IconButton
          color="inherit"
          aria-haspopup="true"
          aria-controls="mail-menu"
          onClick={(e) => {
            setMailMenu(e.currentTarget);
            setIsMailsUnread(false);
          }}
          className={classes.headerMenuButton}
        >
          <Badge
            badgeContent={isMailsUnread ? messages.length : null}
            color="secondary"
          >
            <MailIcon classes={{ root: classes.headerIcon }} />
          </Badge>
        </IconButton>
        <IconButton
          aria-haspopup="true"
          color="inherit"
          className={classes.headerMenuButton}
          aria-controls="profile-menu"
          onClick={(e) => setProfileMenu(e.currentTarget)}
        >
          <AccountIcon classes={{ root: classes.headerIcon }} />
        </IconButton>
        <Menu
          id="mail-menu"
          open={Boolean(mailMenu)}
          anchorEl={mailMenu}
          onClose={() => setMailMenu(null)}
          MenuListProps={{ className: classes.headerMenuList }}
          className={classes.headerMenu}
          classes={{ paper: classes.profileMenu }}
          disableAutoFocusItem
        >
          <div className={classes.profileMenuUser}>
            <Typography variant="h4" weight="medium">
              {t('dashboardLayout.new_messages')}
            </Typography>
            <Typography
              className={classes.profileMenuLink}
              component="a"
              color="secondary"
            >
              {messages.length} {t('dashboardLayout.new_messages')}
            </Typography>
          </div>
          {messages.map((message) => (
            <MenuItem key={message.id} className={classes.messageNotification}>
              <div className={classes.messageNotificationSide}>
                <UserAvatar color={message.variant} name={message.name} />
                <Typography size="sm" color="text" colorBrightness="secondary">
                  {message.time}
                </Typography>
              </div>
              <div
                className={classNames(
                  classes.messageNotificationSide,
                  classes.messageNotificationBodySide,
                )}
              >
                <Typography weight="medium" gutterBottom>
                  {message.name}
                </Typography>
                <Typography color="text" colorBrightness="secondary">
                  {message.message}
                </Typography>
              </div>
            </MenuItem>
          ))}
          {/*
            <Fab
            variant="extended"
            color="primary"
            aria-label="Add"
            className={classes.sendMessageButton}
          >
            Send New Message
            <SendIcon className={classes.sendButtonIcon} />
          </Fab>
          */}
        </Menu>
        <Menu
          id="notifications-menu"
          open={Boolean(notificationsMenu)}
          anchorEl={notificationsMenu}
          onClose={() => setNotificationsMenu(null)}
          className={classes.headerMenu}
          disableAutoFocusItem
        >
          {notifications.map((notification) => (
            <MenuItem
              key={notification.id}
              onClick={() => setNotificationsMenu(null)}
              className={classes.headerMenuItem}
            >
              <Notification {...notification} typographyVariant="inherit" />
            </MenuItem>
          ))}
          {
            (!notifications || !notifications.length)
            && (
              <div className={classes.profileMenuUser}>
                <Typography variant="h4" weight="medium">
                  {t('dashboardLayout.empty_notifications')}
                </Typography>
              </div>
            )
          }
        </Menu>
        <Menu
          id="profile-menu"
          open={Boolean(profileMenu)}
          anchorEl={profileMenu}
          onClose={() => setProfileMenu(null)}
          className={classes.headerMenu}
          classes={{ paper: classes.profileMenu }}
          disableAutoFocusItem
        >
          <div className={classes.profileMenuUser}>
            <Typography variant="h4" weight="medium">
              {userData.email}
            </Typography>
            <Typography
              component="span"
              color="primary"
            >
              {`${userData.firstName || ''} ${userData.lastName || ''}`}
            </Typography>
          </div>
          <MenuItem
            className={classNames(
              classes.profileMenuItem,
              classes.headerMenuItem,
            )}
            onClick={() => router.push('/dashboard/account')}
          >
            <AccountIcon className={classes.profileMenuIcon} /> {t('dashboardLayout.profile_link')}
          </MenuItem>
          <SwitchLang />
          <div className={classes.profileMenuUser}>
            <Typography
              className={classes.profileMenuLink}
              color="primary"
              onClick={actions.logout}
            >
              {t('dashboardLayout.sign_out_link')}
            </Typography>
          </div>
        </Menu>
      </Toolbar>
    </AppBar>
  );
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ logout }, dispatch),
  };
}

function mapStateToProps(store) {
  return {
    user: getUser(store),
  };
}

Header.defaultProps = {
  user: {
    email: '',
    firstName: '',
    lastName: '',
  },
};

Header.propTypes = {
  user: PropTypes.shape({
    email: PropTypes.string,
    firstName: PropTypes.string,
    lastName: PropTypes.string,
  }),
  actions: PropTypes.shape({
    logout: PropTypes.func.isRequired,
  }).isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(Header);
