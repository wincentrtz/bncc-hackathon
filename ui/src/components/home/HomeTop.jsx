import React from "react";
import {Link} from 'react-router-dom';
import { makeStyles } from "@material-ui/core/styles";
import { Grid, Typography, Button } from "@material-ui/core";

import MainImage from "assets/images/main-image.svg";

const useStyles = makeStyles(theme => ({
  root: {
    backgroundImage: "linear-gradient(300deg, #e4007c, #090088 )",
    padding: "80px",
    height: "100vh"
  },
  homeTitle: {
    color: "#fff",
    fontWeight: "bolder"
  },
  signature: {
    color: "#fff"
  },
  button: {
    color: "#fff",
    borderColor: "#fff",
    fontWeight: "bolder",
    padding: "10px 20px"
  },
  img: {
    width: "80%"
  }
}));

const HomeTop = () => {
  const classes = useStyles();
  return (
    <Grid
      className={classes.root}
      container
      alignItems="center"
      justify="space-between"
    >
      <Grid item xs={6}>
        <Typography variant="h2" className={classes.homeTitle} gutterBottom>
          Trippy
        </Typography>
        <Typography variant="h3" className={classes.signature} gutterBottom>
          Simply, Elegant & Fast
        </Typography>
       <Link to="/login"> <Button variant="outlined" className={classes.button}>
          Let's Holiday
        </Button></Link>
      </Grid>
      <Grid item xs={6}>
        <Grid container>
          <Grid item>
            <img className={classes.img} src={MainImage} />
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default HomeTop;
