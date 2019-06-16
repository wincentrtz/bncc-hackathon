import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import AccountCircle from "@material-ui/icons/AccountCircle";

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  title: {
    flexGrow: 1
  }
}));

const Header = () => {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <div className={classes.mainPage}>
        <AppBar className={classes.appBar} position="static">
          <Toolbar>
            <Typography
              color="secondary"
              variant="h6"
              className={classes.title}
            >
              Trippy
            </Typography>
            <AccountCircle />
          </Toolbar>
        </AppBar>
      </div>
    </div>
  );
};

export default Header;
