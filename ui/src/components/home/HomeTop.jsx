import React from "react";
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
    color: "#ffbd39",
    fontWeight: "bolder"
  },
  signature: {
    color: "#ffbd39"
  },
  button: {
    color: "#ffbd39",
    borderColor: "#ffbd39",
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
        <Button variant="outlined" className={classes.button}>
          Let's Holiday
        </Button>
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
