import React from "react";
import { MuiThemeProvider, createMuiTheme } from "@material-ui/core/styles";
import { Switch, Route } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/Login";
import configs from "./configs";
import Album from "./pages/Album";

const theme = createMuiTheme({
  palette: configs.colors,
  typography: { useNextVariants: true }
});

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <Switch>
        <Route path="/album" component={Album} />
        <Route path="/login" component={Login} />
        <Route path="/" component={Home} />
      </Switch>
    </MuiThemeProvider>
  );
}

export default App;
