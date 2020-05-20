/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import { useTranslation } from 'src/i18n';
import {
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Divider,
  Grid,
  Button,
  TextField,
} from '@material-ui/core';

const useStyles = makeStyles(() => ({
  root: {},
}));

const AccountDetails = (props) => {
  const {
    className, user, onUpdate, disabled, onSave,
    ...rest
  } = props;
  const { t } = useTranslation('common');
  const classes = useStyles();

  const values = user;

  const handleChange = (event) => {
    onUpdate({
      [event.target.name]: event.target.value,
    });
  };

  return (
    <Card
      {...rest}
      className={`${classes.root} ${className}`}
    >
      <form
        autoComplete="off"
        noValidate
      >
        <CardHeader
          subheader={t('account.subTitle')}
          title={t('account.title')}
        />
        <Divider />
        <CardContent>
          <Grid
            container
            spacing={3}
          >
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label={t('account.firstName')}
                margin="dense"
                name="firstName"
                onChange={handleChange}
                required
                value={values.firstName}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label={t('account.lastName')}
                margin="dense"
                name="lastName"
                onChange={handleChange}
                required
                value={values.lastName}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label={t('account.email')}
                margin="dense"
                name="email"
                disabled
                onChange={handleChange}
                required
                value={values.email}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label={t('account.mobile')}
                margin="dense"
                name="mobile"
                onChange={handleChange}
                type="number"
                value={values.mobile}
                variant="outlined"
              />
            </Grid>
          </Grid>
        </CardContent>
        <Divider />
        <CardActions>
          <Button
            color="primary"
            variant="contained"
            disabled={disabled}
            onClick={onSave}
          >
            {t('account.save')}
          </Button>
        </CardActions>
      </form>
    </Card>
  );
};

AccountDetails.defaultProps = {
  className: '',
  disabled: false,
};

AccountDetails.propTypes = {
  className: PropTypes.string,
  // eslint-disable-next-line react/forbid-prop-types
  user: PropTypes.object.isRequired,
  onUpdate: PropTypes.func.isRequired,
  disabled: PropTypes.bool,
  onSave: PropTypes.func.isRequired,

};

export default AccountDetails;
