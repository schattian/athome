import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { makeStyles, createStyles } from '@material-ui/core';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { useTranslation } from 'src/i18n';
import Card from '@material-ui/core/Card';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import { isOnline, getWindowSize } from 'src/redux/global/global.selectors';
import { isAuthenticated as isAuthenticatedSelector } from 'src/redux/auth/auth.selectors';
import { logout } from 'src/redux/auth/auth.actions';
import { fetchFakeData } from 'src/example/redux/example.actions';
import { getFakeData } from 'src/example/redux/example.selectors';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Loader, { LoaderTypes } from 'src/components/Loader';
import Error, { ErrorTypes } from 'src/components/Error';
import { Query } from 'net-provider';
import { END_POINTS } from 'src/services/api';
import Link from 'next/link';
import { ROUTES } from 'src/enums';
import PrivateContent from 'src/components/PrivateContent';
import sendNotification from 'src/services/notification';

const useStyles = makeStyles((theme) => createStyles({
  root: {
    margin: 20,
    padding: 15,
  },
  header: {
    backgroundColor: theme.palette.grey[200],
    padding: 15,
    boxShadow: theme.shadows[1],
    margin: 20,
  },
  card: {
    padding: 15,
    marginBottom: 35,
    '& h3': {
      margin: '15px 0px',
    },
  },
}));

function ExampleScreen({
  networkOnline,
  windowHeight,
  windowWidth,
  actions,
  fakeData,
  isAuthenticated,
}) {
  const [count, setCount] = useState(0);
  const { t } = useTranslation('common');
  const classes = useStyles();
  return (
    <Paper className={`${classes.root} ExampleScreen_screen`}>
      {/*
      // --------------------------------------------------------------------
      // Example of using i18n
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">I18nNext</Typography>
        <p>{t('example')}</p>
      </Card>
      {/*
      // --------------------------------------------------------------------
      // Example of using react hook
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">react hook</Typography>
        <p>You clicked {count} times</p>
        <Button variant="contained" color="primary" type="button" onClick={() => setCount(count + 1)}>
          Click me
        </Button>
      </Card>
      {/*
      // --------------------------------------------------------------------
      // Example of sing data from redux, logout, nav to signin with next page
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">Data from redux state</Typography>
        <p>Network connect ? {networkOnline ? 'Yes' : 'No'}</p>
        <p>window Height ? {windowHeight}</p>
        <p>window Width ? {windowWidth}</p>
        <div>isAuthenticated ?
          <span>{isAuthenticated ? 'Yes   ' : 'No   '}</span>
          <span>{isAuthenticated
            ? <Button variant="contained" color="primary" type="button" onClick={actions.logout}>LOGOUT</Button>
            : <Link href={`${ROUTES.SIGNIN_ROUTE}?next=/example/private`}><a>SIGNIN</a></Link>}
          </span>
        </div>
      </Card>
      {/*
      // --------------------------------------------------------------------
      // Example of Fetch data with redux saga, using the Loader with loaderType
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">Fetch data with action</Typography>
        <Button variant="contained" color="primary" type="button" onClick={actions.fetchFakeData}>
          CLick to Fetch
        </Button>
        <Loader loaderType={LoaderTypes.FETCH_FAKE_DATA} />
        <Error errorType={ErrorTypes.FETCH_FAKE_DATA} />
        <List>
          {fakeData && fakeData.slice(0, 10).map((item) => (
            <ListItem key={item.id}>
              <ListItemText>{item.title}</ListItemText>
            </ListItem>
          ))}
        </List>
      </Card>
      {/*
      // --------------------------------------------------------------------
      // Example of Fetch data with the net-provider
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">Fetch data with net-provider</Typography>
        <Query
          query={{
            targetKey: 'fetchFakeData',
            url: END_POINTS.fakeData.url,
          }}
        >
          {
            (res) => (
              <div>
                <Button variant="contained" color="primary" type="button" onClick={res.crudActions.Refresh}>
                  CLick to Refresh
                </Button>
                <Loader isLoading={res.loading} />
                {res.data && res.data.slice(0, 4).map((item) => (
                  <ListItem key={item.id}>
                    <ListItemText>{item.title}</ListItemText>
                  </ListItem>
                ))}
              </div>
            )
          }
        </Query>
      </Card>
      {/*
      // --------------------------------------------------------------------
      // Example of using the <PrivateContent />
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">Example of using the PrivateContent component</Typography>
        <PrivateContent>
          <p>Only Authenticate user can see this message</p>
        </PrivateContent>
        <PrivateContent checkContext={(user) => user.isAdmin}>
          <p>Only Authenticate admin user can see this message</p>
        </PrivateContent>
      </Card>
      {/*
      // --------------------------------------------------------------------
      // Example of send notification
      // --------------------------------------------------------------------
      */}
      <Card className={classes.card}>
        <Typography variant="h3" component="h3">Example of send notification</Typography>
        <Typography variant="p" component="p">You can send notification from any place in your code</Typography>
        <Button variant="contained" color="primary" type="button" onClick={() => sendNotification(`current count ${count}`)}>
          Send Notification
        </Button>
      </Card>
    </Paper>
  );
}

function mapStateToProps(store) {
  const { windowWidth, windowHeight } = getWindowSize(store);
  return {
    networkOnline: isOnline(store),
    windowHeight,
    windowWidth,
    fakeData: getFakeData(store),
    isAuthenticated: isAuthenticatedSelector(store),
  };
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ fetchFakeData, logout }, dispatch),
  };
}

ExampleScreen.defaultProps = {
  networkOnline: null,
  windowHeight: 0,
  windowWidth: 0,
  fakeData: null,
};

ExampleScreen.propTypes = {
  networkOnline: PropTypes.bool,
  windowHeight: PropTypes.number,
  windowWidth: PropTypes.number,
  isAuthenticated: PropTypes.bool.isRequired,
  actions: PropTypes.shape({
    fetchFakeData: PropTypes.func.isRequired,
    logout: PropTypes.func.isRequired,
  }).isRequired,
  fakeData: PropTypes.arrayOf(PropTypes.object),
};
export default connect(mapStateToProps, mapDispatchToProps)(ExampleScreen);
