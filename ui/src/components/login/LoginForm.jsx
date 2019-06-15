import React from "react";
import { Form, Field, reduxForm } from "redux-form";
import { makeStyles } from "@material-ui/core/styles";
import {
  Paper,
  Grid,
  Button,
  Typography,
  FormControlLabel,
  Checkbox,
  Divider
} from "@material-ui/core";
import EmailTextField from "./EmailTextField";
import PasswordTextField from "./PasswordTextField";

const useStyles = makeStyles(theme => ({
  root: {
    height: "100vh"
  },
  container: {
    boxShadow: "0 0 1em black"
  },
  background: {
    backgroundImage: 'url("./assets/images/login-background.jpg")',
    backgroundSize: "cover"
  },
  loginTitle: {
    fontWeight: "bolder"
  },
  subtitle: {
    color: "gray"
  },
  form: {
    padding: "40px 60px"
  },
  forgetPassword: {
    margin: 0,
    fontSize: "0.9em"
  },
  formButton: {
    width: "100%",
    padding: "12px",
    marginTop: "5px"
  },
  divider: {
    marginTop: "15px",
    textAlign: "center"
  },
  registerContainer: {
    margin: "10px 0px"
  },
  footerContainer: {
    marginTop: "20px"
  },
  registerTag: {
    fontSize: "1em",
    color: "gray"
  },
  registerLink: {
    fontSize: "1em",
    cursor: "pointer",
    fontWeight: "bolder",
    marginLeft: "5px"
  }
}));

const Login = ({ handleSubmit }) => {
  const classes = useStyles();
  return (
    <Grid
      className={classes.root}
      container
      justify="center"
      alignItems="center"
    >
      <Grid item xs={9}>
        <Paper className={classes.container}>
          <Grid container>
            <Grid item xs={6} className={classes.form}>
              <Typography
                className={classes.loginTitle}
                variant="h4"
                color="primary"
                gutterBottom
              >
                Welcome Back
              </Typography>
              <Typography
                className={classes.subtitle}
                variant="h6"
                gutterBottom
              >
                Plan Your Trip Now
              </Typography>
              <Form onSubmit={handleSubmit}>
                <Field name="email" component={EmailTextField} />
                <Field name="password" component={PasswordTextField} />
                <Grid container alignItems="center" justify="space-between">
                  <Grid item>
                    <FormControlLabel
                      value="end"
                      control={<Checkbox color="primary" />}
                      label="Remember Me"
                      labelPlacement="end"
                    />
                  </Grid>
                  <Grid item>
                    <Typography
                      className={classes.forgetPassword}
                      variant="h6"
                      gutterBottom
                    >
                      Forget Password?
                    </Typography>
                  </Grid>
                </Grid>
                <Button
                  type="submit"
                  variant="contained"
                  color="primary"
                  className={classes.formButton}
                >
                  Login
                </Button>
              </Form>
              <Grid
                className={classes.divider}
                container
                alignItems="center"
                justify="center"
              >
                <Grid xs={5} item>
                  <Divider />
                </Grid>
                <Grid xs={2} item>
                  <Typography
                    className={classes.registerTag}
                    variant="h6"
                    gutterBottom
                  >
                    Or
                  </Typography>
                </Grid>
                <Grid xs={5} item>
                  <Divider />
                </Grid>
              </Grid>
              <Grid
                className={classes.footerContainer}
                container
                alignItems="center"
              >
                <Grid item>
                  <Typography
                    className={classes.registerTag}
                    variant="h6"
                    gutterBottom
                  >
                    Don't Have Any Account?
                  </Typography>
                </Grid>
                <Grid item>
                  <Typography
                    className={classes.registerLink}
                    variant="h6"
                    color="primary"
                    gutterBottom
                  >
                    Register Here
                  </Typography>
                </Grid>
              </Grid>
            </Grid>
            <Grid item xs={6} className={classes.background} />
          </Grid>
        </Paper>
      </Grid>
    </Grid>
  );
};

export default reduxForm({
  form: "login"
})(Login);
