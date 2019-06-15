import React from "react";
import { TextField } from "@material-ui/core";

const EmailTextField = ({ input, ...custom }) => (
  <TextField
    style={{ width: "100%" }}
    label="Email"
    margin="normal"
    variant="outlined"
    {...input}
    {...custom}
  />
);

export default EmailTextField;
