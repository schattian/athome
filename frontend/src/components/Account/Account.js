import React, { useState, useEffect } from 'react';
import PropsType from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import { Grid } from '@material-ui/core';
import { getUser } from 'src/redux/user/user.selectors';
import { updateMe } from 'src/redux/user/user.actions';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import PrivateContent from 'src/components/PrivateContent';
import { AccountDetails } from './components';

const useStyles = makeStyles((theme) => ({
  root: {
    padding: theme.spacing(2),
  },
}));

const Account = ({ user, actions }) => {
  const classes = useStyles();
  const [updatingUser, updateUser] = useState({});

  useEffect(() => {
    updateUser({});
  }, [user]);

  const onUpdate = (data) => {
    updateUser({ ...updatingUser, ...data });
  };

  const onSave = () => {
    actions.updateMe(updatingUser);
  };
  const disabled = Object.keys(updatingUser).length === 0;

  const userToDisplay = { ...user, ...updatingUser };
  return (
    <PrivateContent>
      <div className={classes.root}>
        <Grid
          container
          spacing={4}
        >
          <Grid
            item
            lg={8}
            md={6}
            xl={8}
            xs={12}
          >
            <AccountDetails
              user={userToDisplay}
              onUpdate={onUpdate}
              disabled={disabled}
              onSave={onSave}
            />
          </Grid>
        </Grid>
      </div>
    </PrivateContent>
  );
};


function mapStateToProps(store) {
  return {
    user: getUser(store),

  };
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ updateMe }, dispatch),
  };
}

Account.defaultProps = {
  user: null,
};
Account.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  user: PropsType.object,
  actions: PropsType.shape({
    updateMe: PropsType.func.isRequired,
  }).isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(Account);
