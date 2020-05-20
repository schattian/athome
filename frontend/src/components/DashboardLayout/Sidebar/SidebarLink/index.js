/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable react/prop-types */
import React from 'react';
import {
  // Collapse,
  Divider,
  // List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Typography,
} from '@material-ui/core';
// import { Inbox as InboxIcon } from '@material-ui/icons';
import Router from 'next/router';

import classnames from 'classnames';

// styles
import useStyles from './styles';

// components
import Dot from '../Dot';


export default function SidebarLink({
  link,
  icon,
  label,
  children,
  isSidebarOpened,
  nested,
  type,
  isLinkActive,
}) {
  const classes = useStyles();

  // local
  // const [isOpen, setIsOpen] = useState(false);
  // function toggleCollapse(e) {
  //   if (isSidebarOpened) {
  //     e.preventDefault();
  //     setIsOpen(!isOpen);
  //   }
  // }


  if (type === 'title') {
    return (
      <Typography
        className={classnames(classes.linkText, classes.sectionTitle, {
          [classes.linkTextHidden]: !isSidebarOpened,
        })}
      >
        {label}
      </Typography>
    );
  }

  if (type === 'divider') return <Divider className={classes.divider} />;

  if (!children) {
    return (
      <ListItem
        button
        to={link}
        onClick={() => {
          if (link) {
            Router.push(link);
          }
        }}
        className={classes.link}
        classes={{
          root: classnames(classes.linkRoot, {
            [classes.linkActive]: isLinkActive && !nested,
            [classes.linkNested]: nested,
          }),
        }}
        disableRipple
      >
        <ListItemIcon
          className={classnames(classes.linkIcon, {
            [classes.linkIconActive]: isLinkActive,
          })}
        >
          {nested ? <Dot color={isLinkActive && 'primary'} /> : icon}
        </ListItemIcon>
        <ListItemText
          classes={{
            primary: classnames(classes.linkText, {
              [classes.linkTextActive]: isLinkActive,
              [classes.linkTextHidden]: !isSidebarOpened,
            }),
          }}
          primary={label}
        />
      </ListItem>
    );
  }
  return children;
}
