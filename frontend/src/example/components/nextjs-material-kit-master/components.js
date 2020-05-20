/*eslint-disable*/
import React from 'react';
// nodejs library that concatenates classes
import classNames from 'classnames';
// react components for routing our app without refresh
import Link from 'next/link';
// @material-ui/core components
import { makeStyles } from '@material-ui/core/styles';
// @material-ui/icons
// core components

import Header from './Header/Header';
import HeaderLinks from './Header/HeaderLinks';
import GridContainer from './Grid/GridContainer';
import GridItem from './Grid/GridItem';
import Button from './CustomButtons/Button';
import Parallax from './Parallax/Parallax';
// sections for this page
import SectionBasics from './pages-sections/Components-Sections/SectionBasics';
import SectionNavbars from './pages-sections/Components-Sections/SectionNavbars';
import SectionTabs from './pages-sections/Components-Sections/SectionTabs';
import SectionPills from './pages-sections/Components-Sections/SectionPills';
import SectionNotifications from './pages-sections/Components-Sections/SectionNotifications';
import SectionTypography from './pages-sections/Components-Sections/SectionTypography';
import SectionJavascript from './pages-sections/Components-Sections/SectionJavascript';
import SectionCompletedExamples from './pages-sections/Components-Sections/SectionCompletedExamples';
import SectionLogin from './pages-sections/Components-Sections/SectionLogin';


const componentsStyle = {
  container: {
    paddingRight: '15px',
    paddingLeft: '15px',
    marginRight: 'auto',
    marginLeft: 'auto',
    width: '100%',
    '@media (min-width: 576px)': {
      maxWidth: '540px',
    },
    '@media (min-width: 768px)': {
      maxWidth: '720px',
    },
    '@media (min-width: 992px)': {
      maxWidth: '960px',
    },
    '@media (min-width: 1200px)': {
      maxWidth: '1140px',
    },
  },
  brand: {
    color: '#FFFFFF',
    textAlign: 'left',
  },
  title: {
    fontSize: '4.2rem',
    fontWeight: '600',
    display: 'inline-block',
    position: 'relative',
  },
  subtitle: {
    fontSize: '1.313rem',
    maxWidth: '510px',
    margin: '10px 0 0',
  },
  main: {
    background: '#FFFFFF',
    position: 'relative',
    zIndex: '3',
  },
  mainRaised: {
    margin: '-60px 30px 0px',
    borderRadius: '6px',
    boxShadow:
      '0 16px 24px 2px rgba(0, 0, 0, 0.14), 0 6px 30px 5px rgba(0, 0, 0, 0.12), 0 8px 10px -5px rgba(0, 0, 0, 0.2)',
    '@media (max-width: 830px)': {
      marginLeft: '10px',
      marginRight: '10px',
    },
  },
  link: {
    textDecoration: 'none',
  },
  textCenter: {
    textAlign: 'center',
  },
};

const useStyles = makeStyles(componentsStyle);

export default function Components(props) {
  const classes = useStyles();
  const { ...rest } = props;
  return (
    <div>
      <Header
        brand="NextJS Material Kit"
        rightLinks={<HeaderLinks />}
        fixed
        color="transparent"
        changeColorOnScroll={{
          height: 400,
          color: 'white',
        }}
        {...rest}
      />
      <Parallax image={require('./nextjs_header.jpg')}>
        <div className={classes.container}>
          <GridContainer>
            <GridItem>
              <div className={classes.brand}>
                <h1 className={classes.title}>NextJS Material Kit.</h1>
                <h3 className={classes.subtitle}>
                  A Badass Material Kit based on Material-UI and NextJS.
                </h3>
              </div>
            </GridItem>
          </GridContainer>
        </div>
      </Parallax>

      <div className={classNames(classes.main, classes.mainRaised)}>
        <SectionBasics />
        <SectionNavbars />
        <SectionTabs />
        <SectionPills />
        <SectionNotifications />
        <SectionJavascript />
        <SectionCompletedExamples />
        <SectionLogin />
        <SectionTypography />
      </div>
    </div>
  );
}
