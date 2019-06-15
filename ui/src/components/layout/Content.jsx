import React from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(theme => ({
  content: {
    padding: "20px 40px"
  }
}));

const Content = ({ render }) => {
  const classes = useStyles();
  return <div className={classes.content}>{render()}</div>;
};

export default Content;
