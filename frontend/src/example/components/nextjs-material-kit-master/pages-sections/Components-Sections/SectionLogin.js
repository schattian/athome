/*eslint-disable*/
import React from 'react';
// @material-ui/core components
import { makeStyles } from '@material-ui/core/styles';
import InputAdornment from '@material-ui/core/InputAdornment';
import Icon from '@material-ui/core/Icon';
// @material-ui/icons
import People from '@material-ui/icons/People';
import Email from '@material-ui/icons/Email';
// core components
import GridContainer from '../../Grid/GridContainer';
import GridItem from '../../Grid/GridItem';
import Card from '../../Card/Card';
import CardHeader from '../../Card/CardHeader';
import CardBody from '../../Card/CardBody';
import CardFooter from '../../Card/CardFooter';
import Button from '../../CustomButtons/Button';
import CustomInput from '../../CustomInput/CustomInput';

import styles from './loginStyle';

const useStyles = makeStyles(styles);

export default function SectionLogin() {
  const classes = useStyles();
  return (
    <div className={classes.section}>
      <div className={classes.container}>
        <GridContainer justify="center">
          <GridItem xs={12} sm={6} md={4}>
            <Card>
              <form className={classes.form}>
                <CardHeader color="primary" className={classes.cardHeader}>
                  <h4>Login</h4>
                  <div className={classes.socialLine}>
                    <Button
                      justIcon
                      href="#pablo"
                      target="_blank"
                      color="transparent"
                      onClick={(e) => e.preventDefault()}
                    >
                      <i className={`${classes.socialIcons} fab fa-twitter`} />
                    </Button>
                    <Button
                      justIcon
                      href="#pablo"
                      target="_blank"
                      color="transparent"
                      onClick={(e) => e.preventDefault()}
                    >
                      <i className={`${classes.socialIcons} fab fa-facebook`} />
                    </Button>
                    <Button
                      justIcon
                      href="#pablo"
                      target="_blank"
                      color="transparent"
                      onClick={(e) => e.preventDefault()}
                    >
                      <i
                        className={
                          `${classes.socialIcons} fab fa-google-plus-g`
                        }
                      />
                    </Button>
                  </div>
                </CardHeader>
                <p className={classes.divider}>Or Be Classical</p>
                <CardBody>
                  <CustomInput
                    labelText="First Name..."
                    id="first"
                    formControlProps={{
                      fullWidth: true,
                    }}
                    inputProps={{
                      type: 'text',
                      endAdornment: (
                        <InputAdornment position="end">
                          <People className={classes.inputIconsColor} />
                        </InputAdornment>
                      ),
                    }}
                  />
                  <CustomInput
                    labelText="Email..."
                    id="email"
                    formControlProps={{
                      fullWidth: true,
                    }}
                    inputProps={{
                      type: 'email',
                      endAdornment: (
                        <InputAdornment position="end">
                          <Email className={classes.inputIconsColor} />
                        </InputAdornment>
                      ),
                    }}
                  />
                  <CustomInput
                    labelText="Password"
                    id="pass"
                    formControlProps={{
                      fullWidth: true,
                    }}
                    inputProps={{
                      type: 'password',
                      endAdornment: (
                        <InputAdornment position="end">
                          <Icon className={classes.inputIconsColor}>
                            lock_outline
                          </Icon>
                        </InputAdornment>
                      ),
                      autoComplete: 'off',
                    }}
                  />
                </CardBody>
                <CardFooter className={classes.cardFooter}>
                  <Button simple color="primary" size="lg">
                    Get started
                  </Button>
                </CardFooter>
              </form>
            </Card>
          </GridItem>
        </GridContainer>
      </div>
    </div>
  );
}
