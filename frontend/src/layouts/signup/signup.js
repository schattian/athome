import React from 'react';
import useForm from 'react-hook-form';
import * as yup from 'yup';
import PropTypes from 'prop-types';
import Logo from 'src/components/Logo';
import Button from '@material-ui/core/Button';
import Checkbox from '@material-ui/core/Checkbox';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Link from 'next/link';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { register as registerAction } from 'src/redux/auth/auth.actions';
import Loader, { LoaderTypes } from 'src/components/Loader';
import Error, { ErrorTypes } from 'src/components/Error';
import Copyright from 'src/components/Copyright';
import { withTranslation } from 'src/i18n';
import { useRouter } from 'next/router';
import useStyles from './styles';

const validationSchema = yup.object().shape({
  email: yup.string().email().required(),
  password: yup.string().required(),
});

function SignUp({ actions, t }) {
  const classes = useStyles();
  const router = useRouter();

  const { register, handleSubmit, errors } = useForm({
    validationSchema,
  });

  const onSubmit = (values, e) => {
    e.preventDefault();
    const actionPayload = {
      values,
      nextRoute: router.query.next || '/',
    };
    actions.register(actionPayload);
  };

  return (
    <Container component="main" maxWidth="xs">
      <div className={classes.paper}>
        <Logo />
        <Typography component="h1" variant="h5">
          {t('title')}
        </Typography>
        <form className={classes.form} noValidate onSubmit={handleSubmit(onSubmit)}>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
              <TextField
                name="firstName"
                variant="outlined"
                fullWidth
                id="firstName"
                label={t('firstNameFieldLabel')}
                autoFocus
                inputRef={register}
                error={!!errors.firstName}
                helperText={errors.firstName && errors.firstName.message}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                variant="outlined"
                fullWidth
                id="lastName"
                label={t('lastNameFieldLabel')}
                name="lastName"
                inputRef={register}
                error={!!errors.lastName}
                helperText={errors.lastName && errors.lastName.message}
              />
            </Grid>
          </Grid>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label={t('emailFieldLabel')}
            name="email"
            autoComplete="email"
            inputRef={register}
            error={!!errors.email}
            helperText={errors.email && errors.email.message}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label={t('passwordFieldLabel')}
            type="password"
            id="password"
            autoComplete="new-password"
            inputRef={register}
            error={!!errors.password}
            helperText={errors.password && errors.password.message}
          />
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label={t('signupConfirm')}
          />
          <Error errorType={ErrorTypes.REGISTER} />
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
                <a><Typography component="span" variant="body2">{t('haveAnAccount')}</Typography></a>
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
      <Loader fullScreen loaderType={LoaderTypes.REGISTER} />
    </Container>
  );
}


function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ register: registerAction }, dispatch),
  };
}
SignUp.propTypes = {
  t: PropTypes.func.isRequired,
  actions: PropTypes.shape({
    register: PropTypes.func.isRequired,
  }).isRequired,
};

const Extend = withTranslation('signup')(SignUp);

export default connect(null, mapDispatchToProps)(Extend);
