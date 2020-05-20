import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import Toolbar from '@material-ui/core/Toolbar';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import Copyright from 'src/components/Copyright';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { withTranslation } from 'src/i18n';
import Link from 'next/link';
import { isAuthenticated } from 'src/redux/auth/auth.selectors';
import { logout } from 'src/redux/auth/auth.actions';
import TypographyExample from 'src/example/components/TypographyExample/TypographyExample';
import ButtonExample from 'src/example/components/ButtonExample/ButtonExample';
import { title } from '../../../siteConfig';
import './home.scss';

const useStyles = makeStyles((theme) => ({
  toolbar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
  toolbarTitle: {
    flex: 1,
  },
  toolbarLink: {
    padding: theme.spacing(1),
    flexShrink: 0,
  },
  mainFeaturedPost: {
    position: 'relative',
    backgroundColor: theme.palette.grey[800],
    color: theme.palette.common.white,
    marginBottom: theme.spacing(4),
    backgroundImage: 'url(./static/example.jpg)',
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    backgroundPosition: 'center',
  },
  overlay: {
    position: 'absolute',
    top: 0,
    bottom: 0,
    right: 0,
    left: 0,
    backgroundColor: 'rgba(0,0,0,.3)',
  },
  mainFeaturedPostContent: {
    position: 'relative',
    padding: theme.spacing(3),
  },
  mainGrid: {
    marginTop: theme.spacing(3),
  },
  footer: {
    backgroundColor: theme.palette.background.paper,
    marginTop: theme.spacing(8),
    padding: theme.spacing(6, 0),
  },
  signup: {
    margin: '0 10px',
  },
}));


function HomeScreen({ t, isAuth, actions }) {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Container maxWidth="lg">
        <Toolbar className={classes.toolbar}>
          <Typography
            component="h2"
            variant="h5"
            color="inherit"
            align="center"
            noWrap
            className={classes.toolbarTitle}
          >
            {title}
          </Typography>
          {
            isAuth
              ? (
                <Button
                  variant="outlined"
                  size="small"
                  component="a"
                  onClick={actions.logout}
                >
                  LOGOUT
                </Button>
              )
              : (
                <React.Fragment>
                  <Link href="signup">
                    <Button variant="outlined" size="small" component="a" className={classes.signup}>
                      Sign up
                    </Button>
                  </Link>
                  <Link href="signin">
                    <Button variant="outlined" size="small" component="a">
                      Sign in
                    </Button>
                  </Link>
                </React.Fragment>
              )
          }
        </Toolbar>
        <main>
          <Paper className={classes.mainFeaturedPost}>
            <div className={classes.overlay} />
            <Grid container>
              <Grid item md={6}>
                <div className={classes.mainFeaturedPostContent}>
                  <Typography component="h1" variant="h3" color="inherit" gutterBottom>
                    Welcome to <br />{title}
                  </Typography>
                </div>
              </Grid>
            </Grid>
          </Paper>
          <p>{t('screenName')}</p>
          <br />
          <img src="/static/lamp.png" alt="Example of media from static." />
          <br />
          <Link href="/example/public"><a>Example Public screen</a></Link>
          <br />
          <Link href="/example/private"><a>Example Private screen</a></Link>
          <br />
          <Link href="/example/dashboard/dashboard-posts"><a>Example Dashboard screen</a></Link>
          <br />
          <Link href="/example/nextjs-material-kit-master"><a>Example nextjs-material-kit-master</a></Link>
        </main>
        <TypographyExample />
        <ButtonExample />
      </Container>
      {/* Footer */}
      <footer className={classes.footer}>
        <Container maxWidth="lg">
          <Copyright />
        </Container>
      </footer>
      {/* End footer */}
    </React.Fragment>
  );
}

const Extended = withTranslation('home')(HomeScreen);

function mapStateToProps(store) {
  return {
    isAuth: isAuthenticated(store),
  };
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ logout }, dispatch),
  };
}
HomeScreen.defaultProps = {
  isAuth: null,
};
HomeScreen.propTypes = {
  isAuth: PropTypes.bool,
  actions: PropTypes.shape({
    logout: PropTypes.func.isRequired,
  }).isRequired,
  t: PropTypes.func.isRequired,
};
export default connect(mapStateToProps, mapDispatchToProps)(Extended);
