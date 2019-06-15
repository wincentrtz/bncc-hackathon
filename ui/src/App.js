import React from "react";
import { MuiThemeProvider, createMuiTheme } from "@material-ui/core/styles";
import { Switch, Route } from "react-router-dom";
import Layout from "./components/layout";
import Home from "./pages/Home";
import Login from "./pages/Login";
import configs from "./configs";

const theme = createMuiTheme({
  palette: configs.colors,
  typography: { useNextVariants: true }
});

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <Switch>
        <Route path="/login" component={Login} exact />
        <Route path="/" component={Layout(Home)} exact />
      </Switch>
    </MuiThemeProvider>
  );
}

export default App;
