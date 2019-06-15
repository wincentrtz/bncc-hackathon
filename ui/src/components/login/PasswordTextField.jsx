import React from "react";
import { TextField } from "@material-ui/core";

const PasswordTextField = ({ input, ...custom }) => {
  return (
    <TextField
      label="Password"
      style={{ width: "100%" }}
      type="password"
      margin="normal"
      variant="outlined"
      {...input}
      {...custom}
    />
  );
};

export default PasswordTextField;
