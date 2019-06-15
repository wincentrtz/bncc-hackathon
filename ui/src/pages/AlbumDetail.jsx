import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Typography, Grid, Chip } from "@material-ui/core";

import Layout from "../components/layout";

const useStyles = makeStyles(theme => ({
  background: {
    backgroundImage: 'url("../../assets/images/login-background.jpg")',
    backgroundSize: "cover",
    width: "100%",
    height: "400px"
  },
  albumTitle: {
    fontWeight: "bolder"
  },
  albumHeader: {
    margin: "20px 60px"
  },
  chip: {
    fontWeight: "bolder",
    color: "white"
  }
}));

const renderColor = flag => (flag ? "#2196f3" : "#e10050");
const renderChipLabel = flag => (flag ? "Paid" : "Unpaid");

const AlbumDetail = () => {
  const classes = useStyles();
  return (
    <div>
      <Grid container justify="center">
        <Grid item xs={10} className={classes.background} />
        <Grid item xs={10} className={classes.albumHeader}>
          <Grid container justify="space-between">
            <Grid item xs={5}>
              <Typography className={classes.albumTitle} variant="h4">
                Jalan - jalan bodoh
              </Typography>
            </Grid>
            <Grid item xs={5}>
              <Chip
                label={renderChipLabel(true)}
                className={classes.chip}
                style={{ backgroundColor: renderColor(true) }}
              />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
};

export default Layout(AlbumDetail);
