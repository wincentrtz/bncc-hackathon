import React from "react";
import Header from "./Header";
import Content from "./Content";

const Layout = Component => props => {
  return (
    <div>
      <Header />
      <Content render={() => <Component {...props} />} />
    </div>
  );
};

export default Layout;
