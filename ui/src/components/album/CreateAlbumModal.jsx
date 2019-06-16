import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import {
  Typography,
  Modal,
  TextField,
  Select,
  OutlinedInput,
  Button,
  FormControl,
  InputLabel
} from "@material-ui/core";

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

  const handleCreate = () => {
    window.location = "/album/1";
  };
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
        <TextField
          style={{ width: "100%" }}
          label="Album Name"
          margin="normal"
          variant="outlined"
        />
        <FormControl
          style={{ width: "100%", marginBottom: "10px" }}
          variant="outlined"
          className={classes.formControl}
        >
          <InputLabel htmlFor="outlined-age-native-simple">From</InputLabel>
          <Select
            native
            input={<OutlinedInput name="age" id="outlined-age-native-simple" />}
          >
            <option />
            <option value={"JKT"}>Jakarta</option>
            <option value={"BDG"}>Bandung</option>
            <option value={"BAL"}>Bali</option>
          </Select>
        </FormControl>
        <FormControl
          style={{ width: "100%", marginBottom: "10px" }}
          variant="outlined"
          className={classes.formControl}
        >
          <InputLabel htmlFor="outlined-age-native-simple">To</InputLabel>
          <Select
            native
            style={{ width: "100%", marginBottom: "10px" }}
            input={<OutlinedInput name="age" id="outlined-age-native-simple" />}
          >
            <option />
            <option value={"JKT"}>Jakarta</option>
            <option value={"BDG"}>Bandung</option>
            <option value={"BAL"}>Bali</option>
          </Select>
        </FormControl>
        <Button
          onClick={handleCreate}
          variant="outlined"
          style={{ width: "100%" }}
        >
          CREATE
        </Button>
      </div>
    </Modal>
  );
};

export default CreateAlbumModal;
