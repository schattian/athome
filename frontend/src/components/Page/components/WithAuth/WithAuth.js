/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable camelcase */
/* eslint-disable no-nested-ternary */
/* eslint-disable react/destructuring-assignment */
/* eslint-disable react/static-property-placement */
import React from 'react';
import PropTypes from 'prop-types';
import Router from 'next/router';
import { connect } from 'react-redux';
import {
  isLoading,
  getTokenValidateState,
  getLastAction,
} from 'src/redux/auth/auth.selectors';
import { reAuthenticate, actionsType, setLoading } from 'src/redux/auth/auth.actions';
import { isOnline } from 'src/redux/global/global.selectors';
import { getToken as getTokenFromStorage } from 'src/services/userToken';
import { ROUTES } from 'src/enums';
import logger from 'src/services/logger';
import Loader from 'src/components/Loader';


export default function withAuth(
  BaseComponent,
  { loginRequired, logoutRequired, adminRequired } = {},
) {
  class App extends React.PureComponent {
    static propTypes = {
      auth_isLoading: PropTypes.bool,
      auth_isTokenValid: PropTypes.bool,
      auth_lastAction: PropTypes.string,
      global_isOnline: PropTypes.bool,
      dispatch: PropTypes.func.isRequired,
    };

    static defaultProps = {
      auth_isLoading: null,
      auth_isTokenValid: null,
      auth_lastAction: null,
      global_isOnline: null,
    };

    static async getInitialProps(ctx) {
      // const isFromServer = !!ctx.req;
      const props = {};

      if (BaseComponent.getInitialProps) {
        Object.assign(props, (await BaseComponent.getInitialProps(ctx)) || {});
      }

      return props;
    }

    constructor(props) {
      super(props);
      this.state = {
        displayRedirect: false,
        checkAbilityEnd: (!loginRequired && !logoutRequired && !adminRequired)
          || (loginRequired && this.props.auth_isTokenValid)
          || (logoutRequired && this.props.auth_isTokenValid === false),
      };
    }


    componentDidMount() {
      const { auth_isLoading, dispatch } = this.props;
      const isTokenCheckNotStart = auth_isLoading === null;
      const isTokenCheckEnds = auth_isLoading === false;
      const tokenFromStorage = getTokenFromStorage();
      if (isTokenCheckNotStart) {
        if (tokenFromStorage) {
          reAuthenticate(dispatch, tokenFromStorage);
        } else {
          dispatch(setLoading(false));
          this.checkAbility();
        }
      } else if (isTokenCheckEnds || !tokenFromStorage) {
        this.checkAbility();
      }
    }

    componentDidUpdate() {
      const { auth_isLoading, auth_lastAction } = this.props;
      const isTokenCheckEnds = this.lastAuth_isLoading === null && auth_isLoading !== null;
      const isLogoutEnd = this.auth_lastAction !== actionsType.ON_LOGOUT_END
        && auth_lastAction === actionsType.ON_LOGOUT_END;
      if (isTokenCheckEnds) {
        this.checkAbility();
      } else if (isLogoutEnd && loginRequired) {
        this.redirect('/');
      }
      this.auth_lastAction = auth_lastAction;
      this.lastAuth_isLoading = auth_isLoading;
    }

    redirect = (route) => {
      this.setState({ displayRedirect: true });
      setTimeout(() => {
        Router.push(route);
      }, 1000);
      logger.debug(`
      route is ${loginRequired ? 'loginRequired' : (logoutRequired ? 'logoutRequired' : (adminRequired ? 'adminRequired' : ''))}
      and user is ${this.props.auth_isTokenValid ? 'Logged' : 'Not Logged'}, navigate to ${route}`);
    }


    checkAbility() {
      const { auth_isTokenValid } = this.props;
      const isLoggedIn = auth_isTokenValid === true;
      if (loginRequired && !isLoggedIn) {
        this.redirect(`${ROUTES.SIGNIN_ROUTE}?next=${Router.router.pathname}`);
      } else if (!this.checkAbilityRun && logoutRequired && isLoggedIn) {
        this.redirect('/');
      }
      this.checkAbilityRun = true;
      this.setState({
        checkAbilityEnd: true,
      });
    }

    render() {
      const { displayRedirect, checkAbilityEnd } = this.state;
      if (displayRedirect) return <Loader isLoading fullScreen />; // 'Redirect...';
      if (!checkAbilityEnd) return <Loader isLoading fullScreen />; // 'Loading...';
      return (
        <React.Fragment>
          <BaseComponent {...this.props} />
        </React.Fragment>
      );
    }
  }

  const mapStateToProps = (store) => ({
    auth_isLoading: isLoading(store),
    auth_isTokenValid: getTokenValidateState(store),
    auth_lastAction: getLastAction(store),
    global_isOnline: isOnline(store),
  });

  return connect(mapStateToProps)(App);
}
