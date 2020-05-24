import React from 'react';
import useForm from 'react-hook-form';
import * as yup from 'yup';
import PropTypes from 'prop-types';
import Logo from 'src/components/Logo';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Link from 'next/link';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { sendVerifyAccount } from 'src/redux/auth/auth.actions';
import Loader, { LoaderTypes } from 'src/components/Loader';
import Error, { ErrorTypes } from 'src/components/Error';
import Copyright from 'src/components/Copyright';
import { withTranslation } from 'src/i18n';
import { useRouter } from 'next/router';
import useStyles from './styles';

const validationSchema = yup.object().shape({
  email: yup.string().email().required(),
});

function verifyAccount({ actions, t }) {
  const classes = useStyles();
  const router = useRouter();
  const { register, handleSubmit, errors } = useForm({
    validationSchema,
  });


  const onSubmit = (values, e) => {
    e.preventDefault();
    const actionPayload = {
      values,
      notificationMessage: t('notificationMessageOnSuccess'),
    };
    actions.sendVerifyAccount(actionPayload);
  };

  const valueFromParams = router.query.email || '';

  return (
    <Container component="main" maxWidth="xs">
      <div className={classes.paper}>
        <Logo />
        <Typography component="h1" variant="h5">
          {t('title')}
        </Typography>
        <br />
        <Typography component="p">
          {t('subTitle')}
        </Typography>
        <form className={classes.form} noValidate onSubmit={handleSubmit(onSubmit)}>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label={t('emailFieldLabel')}
            name="email"
            defaultValue={valueFromParams}
            autoComplete="email"
            autoFocus
            inputRef={register}
            error={errors.email}
            helperText={errors.email && errors.email.message}
          />
          <Error errorType={ErrorTypes.VERIFY_ACCOUNT} />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            {t('submitButton')}
          </Button>
          <Grid container>
            <Grid item>
              <Link href="/signin">
                <a><Typography component="span" variant="body2">{t('signinButtom')}</Typography></a>
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
      <Loader fullScreen loaderType={LoaderTypes.VERIFY_ACCOUNT} />
    </Container>
  );
}


function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ sendVerifyAccount }, dispatch),
  };
}
verifyAccount.propTypes = {
  t: PropTypes.func.isRequired,
  actions: PropTypes.shape({
    login: PropTypes.func.isRequired,
    sendVerifyAccount: PropTypes.func.isRequired,
  }).isRequired,
};

const Extend = withTranslation('verifyAccount')(verifyAccount);

export default connect(null, mapDispatchToProps)(Extend);
