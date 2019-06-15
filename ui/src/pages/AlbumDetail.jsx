import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import {
  Typography,
  Grid,
  Chip,
  Paper,
  Button,
  TextField,
  Divider
} from "@material-ui/core";
import AccountCircle from "@material-ui/icons/AccountCircle";

const useStyles = makeStyles(theme => ({
  root: {
    backgroundImage: "linear-gradient(300deg, #e4007c, #090088 )",
    padding: "80px",
    height: "100vh"
  },
  albumDetailPaper: {
    padding: "20px",
    marginBottom: "40px"
  },
  albumTitle: {
    marginBottom: "10px",
    fontWeight: "bolder"
  },
  iconUser: {
    fontSize: "4.5rem"
  },
  formDate: {
    padding: "20px"
  },
  listFlight: {
    padding: "20px"
  }
}));

const renderColor = flag => (flag ? "#2196f3" : "#e10050");
const renderChipLabel = flag => (flag ? "Paid" : "Unpaid");

const AlbumDetail = () => {
  const classes = useStyles();

  function handleSelect(date) {
    console.log(date); // Momentjs object
  }
  return (
    <div className={classes.root}>
      <Paper className={classes.albumDetailPaper}>
        <Grid container justify="space-around" alignItems="center">
          <Grid item xs={6}>
            <Typography className={classes.albumTitle} variant="h4">
              Jakarta - Bandung
            </Typography>
            <Grid container>
              <Grid item xs={2} style={{ textAlign: "center" }}>
                <AccountCircle className={classes.iconUser} />
              </Grid>
              <Grid item xs={2} style={{ textAlign: "center" }}>
                <AccountCircle className={classes.iconUser} />
              </Grid>
              <Grid item xs={2} style={{ textAlign: "center" }}>
                <AccountCircle className={classes.iconUser} />
              </Grid>
            </Grid>
          </Grid>
          <Grid item xs={6}>
            <Typography className={classes.albumTitle} variant="h4">
              Rp. 0
            </Typography>
            <Button variant="outlined">Checkout</Button>
          </Grid>
        </Grid>
      </Paper>
      <Paper>
        <Grid container className={classes.formDate} alignItems="center">
          <Grid item xs={4}>
            <TextField type="date" margin="normal" variant="outlined" />
          </Grid>
          <Grid item xs={4}>
            <TextField type="date" margin="normal" variant="outlined" />
          </Grid>
          <Grid item xs={4}>
            <Button variant="outlined">Search</Button>
          </Grid>
        </Grid>
        <Paper className={classes.listFlight}>
          <Grid container alignItems="center">
            <Grid item xs={10}>
              <Grid
                style={{ margin: "10px 10px" }}
                container
                alignItems="center"
              >
                <Grid item xs={3}>
                  Lion Air
                </Grid>
                <Grid item xs={3}>
                  13.00
                  <br /> JKT
                </Grid>
                <Grid item xs={3}>
                  14.40
                  <br /> PKU
                </Grid>
                <Grid item xs={3}>
                  1j 40m
                </Grid>
              </Grid>
              <Divider />
              <Grid
                style={{ margin: "10px 10px" }}
                container
                alignItems="center"
              >
                <Grid item xs={3}>
                  Lion Air
                </Grid>
                <Grid item xs={3}>
                  13.00
                  <br /> PKU
                </Grid>
                <Grid item xs={3}>
                  14.40
                  <br /> JKT
                </Grid>
                <Grid item xs={3}>
                  1j 40m
                </Grid>
              </Grid>
            </Grid>
            <Grid item xs={2}>
              <Grid container>
                <Grid item xs={12}>
                  Rp 1.073.200/pax
                </Grid>
                <Grid item xs={12}>
                  <Button variant="outlined">BUY</Button>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        </Paper>
      </Paper>
    </div>
  );
};

export default AlbumDetail;
