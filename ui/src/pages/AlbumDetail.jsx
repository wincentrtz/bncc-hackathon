import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { connect } from "react-redux";
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

import { fetchAlbumDetail } from "store/actions/album-actions";
import { fetchFlights } from "../store/actions/album-actions";
import { axiosInstance } from "../services/axios.config";

import axios from "axios";

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

const AlbumDetail = props => {
  const classes = useStyles();
  const [price, setPrice] = React.useState(0);

  function handleGetAlbumDetail() {
    props.doGetAlbumDetail();
  }
  useEffect(() => handleGetAlbumDetail(), []);

  function searchFlights() {
    props.doGetFlights();
  }


  function handleBuy(newPrice) {
    setPrice(newPrice * 3)
  }
  
  function checkOut() {
    const data = axios.post('http://localhost:8000/otp/token');
    console.log(data)
  }

  return (
    <div className={classes.root}>
      <Paper className={classes.albumDetailPaper}>
        <Grid container justify="space-around" alignItems="center">
          <Grid item xs={6}>
            <Typography className={classes.albumTitle} variant="h4">
              Jalan - jalan keluarga
            </Typography>
            <Grid container>
              {[1,2,3].map(() => (
                  <Grid item xs={2} style={{ textAlign: "center" }}>
                    <AccountCircle className={classes.iconUser} />
                  </Grid>
                ))}
            </Grid>
          </Grid>
          <Grid item xs={6}>
            <Typography className={classes.albumTitle} variant="h4">
              Rp. {price}
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
            <Button onClick={searchFlights} variant="outlined">
              Search
            </Button>
          </Grid>
        </Grid>
        {props.flights.map(flight => (
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
                    <br /> {flight.destination}
                  </Grid>
                  <Grid item xs={3}>
                    14.40
                    <br /> {flight.departure}
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
                    <br /> {flight.departure}
                  </Grid>
                  <Grid item xs={3}>
                    14.40
                    <br /> {flight.destination}
                  </Grid>
                  <Grid item xs={3}>
                    1j 40m
                  </Grid>
                </Grid>
              </Grid>
              <Grid item xs={2}>
                <Grid container>
                  <Grid item xs={12}>
                    {"Rp. "+flight.price}
                  </Grid>
                  <Grid item xs={12}>
                    <Button onClick={() => handleBuy(flight.price)} variant="outlined">BUY</Button>
                  </Grid>
                </Grid>
              </Grid>
            </Grid>
          </Paper>
        ))}
      </Paper>
    </div>
  );
};

const mapStateToProps = ({ albumReducers }) => {
  return {
    album: albumReducers.album,
    flights: albumReducers.flights
  };
};

const mapDispatchToProps = dispatch => {
  return {
    doGetFlights: () => dispatch(fetchFlights()),
    doGetAlbumDetail: () => dispatch(fetchAlbumDetail())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AlbumDetail);
