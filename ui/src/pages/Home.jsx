import React from "react";
import { makeStyles } from "@material-ui/core/styles";

import HomeTop from "../components/home/HomeTop";

const useStyles = makeStyles(theme => ({
  img: {
    width: "50%"
  }
}));

const Home = () => {
  const classes = useStyles();
  return (
    <div>
      <HomeTop />
    </div>
  );
};

export default Home;
