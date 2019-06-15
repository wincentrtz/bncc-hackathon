import React from "react";
import Header from "./Header";
import Content from "./Content";

const Layout = props => () => {
  return (
    <div>
      <Header />
      <Content render={props} />
    </div>
  );
};

export default Layout;
