import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Typography, Modal } from "@material-ui/core";

const useStyles = makeStyles(theme => ({
  paper: {
    width: 400,
    backgroundColor: theme.palette.background.paper,
    boxShadow: theme.shadows[5],
    padding: theme.spacing(4),
    outline: "none",
    margin: "30px auto"
  }
}));

const CreateAlbumModal = ({ onClose, open }) => {
  const classes = useStyles();
  return (
    <Modal
      aria-labelledby="simple-modal-title"
      aria-describedby="simple-modal-description"
      open={open}
      onClose={onClose}
    >
      <div className={classes.paper}>
        <Typography variant="h6" id="modal-title">
          Create Album
        </Typography>
      </div>
    </Modal>
  );
};

export default CreateAlbumModal;
