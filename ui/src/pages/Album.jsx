import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { connect } from "react-redux";
import { Grid, Typography, Button } from "@material-ui/core";

import AlbumCard from "../components/album/AlbumCard";

import { fetchAlbums } from "store/actions/album-actions";

const useStyles = makeStyles(theme => ({
  root: {
    backgroundImage: "linear-gradient(300deg, #e4007c, #090088 )",
    height: "100vh"
  },
  albumContainer: {
    overflowX: "scroll"
  },
  albumScroll: {
    flexWrap: "nowrap"
  },
  albumList: {
    minWidth: "400px",
    marginLeft: "40px"
  },
  button: {
    color: "#fff",
    borderColor: "#fff",
    fontWeight: "bolder",
    padding: "10px 20px"
  }
}));

const renderAlbums = (classes, albums) => {
  return albums.map(() => (
    <Grid item xs={8} className={classes.albumList}>
      <AlbumCard />
    </Grid>
  ));
};

const Album = props => {
  const classes = useStyles();
  function handleLogin() {
    props.doGetAlbums();
  }
  useEffect(() => handleLogin(), []);
  return (
    <Grid
      container
      className={classes.root}
      alignItems="center"
      justify="center"
    >
      <Grid item>
        <Typography variant="h3" color="secondary">
          Albums
        </Typography>
      </Grid>
      <Grid item className={classes.albumContainer}>
        <Grid
          className={classes.albumScroll}
          container
          direction="row"
          alignItems="center"
        >
          {renderAlbums(classes, props.albums)}
        </Grid>
      </Grid>
      <Grid item>
        <Button variant="outlined" className={classes.button}>
          Create Holiday
        </Button>
      </Grid>
    </Grid>
  );
};

const mapStateToProps = ({ albumReducers }) => {
  return {
    albums: albumReducers.albums
  };
};

const mapDispatchToProps = dispatch => {
  return {
    doGetAlbums: () => dispatch(fetchAlbums())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Album);
