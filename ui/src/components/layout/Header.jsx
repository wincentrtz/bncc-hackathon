import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import AccountCircle from "@material-ui/icons/AccountCircle";
import { Grid } from "@material-ui/core";

import MainImage from "assets/images/main-image.svg";

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
            <Typography variant="h6" className={classes.title}>
              Ticket.com
            </Typography>
            <AccountCircle />
          </Toolbar>
        </AppBar>
      </div>
    </div>
  );
};

export default Header;
