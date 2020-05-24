import Container from "@material-ui/core/Container";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import React from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import Copyright from "src/components/Copyright";
import Navbar from "src/containers/Navbar";
import { withTranslation } from "src/i18n";
import { logout } from "src/redux/auth/auth.actions";
import { isAuthenticated } from "src/redux/auth/auth.selectors";
import { title } from "../../../siteConfig";

const useStyles = makeStyles((theme) => ({
    footer: {
        backgroundColor: theme.palette.background.paper,
        marginTop: theme.spacing(8),
        padding: theme.spacing(6, 0),
    },
    signup: {
        margin: "0 10px",
    },
}));

function HomeScreen() {
    const classes = useStyles();
    return (
        <React.Fragment>
            <Navbar />
            <main>
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
            </main>
            <footer className={classes.footer}>
                <Container maxWidth="lg">
                    <Copyright />
                </Container>
            </footer>
        </React.Fragment>
    );
}

const Extended = withTranslation("home")(HomeScreen);

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
HomeScreen.defaultProps = {};
HomeScreen.propTypes = {};
export default connect(mapStateToProps, mapDispatchToProps)(Extended);
