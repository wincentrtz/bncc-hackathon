import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid } from "@material-ui/core";

import AlbumCard from "../components/album/AlbumCard";
import Layout from "../components/layout";

import { fetchAlbums } from "store/actions/album-actions";

const renderAlbums = albums => {
  return albums.map(() => (
    <Grid item xs={4}>
      <AlbumCard />
    </Grid>
  ));
};

const Album = props => {
  useEffect(() => props.doGetAlbums(), []);
  return <Grid container>{renderAlbums(props.albums)}</Grid>;
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
)(Layout(Album));
