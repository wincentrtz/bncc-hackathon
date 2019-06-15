import React from "react";
import { connect } from "react-redux";

import LoginForm from "components/login/LoginForm";
import { loginUser } from "store/actions/user-actions";

const Login = ({ doLogin }) => {
  const submit = value => {
    doLogin(value);
  };

  return <LoginForm onSubmit={submit} />;
};

const mapDispatchToProps = dispatch => {
  return {
    doLogin: value => dispatch(loginUser(value))
  };
};

export default connect(
  null,
  mapDispatchToProps
)(Login);
