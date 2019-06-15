import React, { Component } from "react";
import { connect } from "react-redux";
import { Grid } from "@material-ui/core";

import AlbumCard from "../components/album/AlbumCard";

import { fetchAlbums } from "store/actions/album-actions";

class Album extends Component {
  componentDidMount() {
    this.props.doGetAlbums();
  }

  renderAlbums = () => {
    const { albums } = this.props;
    console.log(albums);
    return albums.map(() => (
      <Grid item xs={4}>
        <AlbumCard />
      </Grid>
    ));
  };

  render() {
    return <Grid container>{this.renderAlbums()}</Grid>;
  }
}

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
