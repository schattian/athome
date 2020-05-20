
import React from 'react';
import MuLink from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import { title, url } from '../../../siteConfig';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      Copyright ©
      <MuLink color="inherit" href={url}>
        {title}
      </MuLink>
      {new Date().getFullYear()}
      .
    </Typography>
  );
}
export default Copyright;
