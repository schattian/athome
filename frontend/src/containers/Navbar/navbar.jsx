import Button from "@material-ui/core/Button";
import { makeStyles } from "@material-ui/core/styles";
import Toolbar from "@material-ui/core/Toolbar";
import Link from "next/link";
import React from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { withTranslation } from "src/i18n";
import { logout } from "src/redux/auth/auth.actions";
import { isAuthenticated } from "src/redux/auth/auth.selectors";

const useStyles = makeStyles((theme) => ({
    buttonContainer: { marginLeft: "auto" },
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
    signup: {
        margin: "0 10px",
    },
}));

function NavbarContainer({ t, isAuth, actions }) {
    const classes = useStyles();
    return (
        <Toolbar className={classes.toolbar}>
            {isAuth ? (
                <Button
                    variant="outlined"
                    size="small"
                    component="a"
                    onClick={actions.logout}
                >
                    {t("logout")}
                </Button>
            ) : (
                <>
                    <Link href="signup">
                        <Button
                            variant="outlined"
                            size="small"
                            component="a"
                            className={classes.signup}
                        >
                            {t("signup")}
                        </Button>
                    </Link>
                    <Link href="signin">
                        <Button variant="outlined" size="small" component="a">
                            {t("signin")}
                        </Button>
                    </Link>
                </>
            )}
        </Toolbar>
    );
}

const Container = withTranslation("navbar")(NavbarContainer);

const mapStateToProps = (state) => ({ isAuth: isAuthenticated(state) });

const mapDispatchToProps = (dispatch) => ({
    actions: bindActionCreators({ logout }, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Container);
