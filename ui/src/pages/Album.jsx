import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { connect } from "react-redux";
import {
  Grid,
  Typography,
  Button,
  IconButton,
  Drawer,
  Modal
} from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";

import AlbumCard from "../components/album/AlbumCard";

import { fetchAlbums } from "store/actions/album-actions";
import Sidelist from "../components/common/Sidelist";
import CreateAlbumModal from "../components/album/CreateAlbumModal";

const useStyles = makeStyles(theme => ({
  root: {
    backgroundImage: "linear-gradient(300deg, #e4007c, #090088 )",
    height: "100vh"
  },
  list: {
    width: 250
  },
  headerAlbum: {
    width: "100%"
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
  },
  paper: {
    width: 400,
    backgroundColor: theme.palette.background.paper,
    boxShadow: theme.shadows[5],
    padding: theme.spacing(4),
    outline: "none",
    margin: "30px auto"
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
  const [open, setOpen] = React.useState(false);
  const [state, setState] = React.useState({
    top: false
  });

  const toggleDrawer = (side, open) => event => {
    if (
      event.type === "keydown" &&
      (event.key === "Tab" || event.key === "Shift")
    ) {
      return;
    }

    setState({ ...state, [side]: open });
  };

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <Drawer open={state.left} onClose={toggleDrawer("left", false)}>
        <Sidelist onToggle={toggleDrawer} side={"list"} />
      </Drawer>
      <Grid
        container
        className={classes.root}
        alignItems="center"
        justify="center"
      >
        <Grid item className={classes.headerAlbum}>
          <Grid container justify="center">
            <Grid item xs={1} />
            <Grid item xs={3}>
              <IconButton
                onClick={toggleDrawer("left", true)}
                color="secondary"
                aria-label="Open drawer"
                edge="start"
              >
                <MenuIcon />
              </IconButton>
            </Grid>
            <Grid item xs={4}>
              <Typography
                variant="h3"
                color="secondary"
                style={{ textAlign: "center" }}
              >
                Albums
              </Typography>
            </Grid>
            <Grid item xs={4} />
          </Grid>
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
          <Button
            onClick={handleOpen}
            variant="outlined"
            className={classes.button}
          >
            Create Album
          </Button>
        </Grid>
      </Grid>
      <CreateAlbumModal open={open} onClose={handleClose} />
    </div>
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
