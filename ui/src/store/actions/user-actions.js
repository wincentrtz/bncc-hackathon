import { LOGIN_USER } from "../constant/user-constants";

import api from "services/modules/user";

export const loginUser = payload => async dispatch => {
  const { data } = await api.login(payload);
  console.log(data)
  dispatch({
    type: LOGIN_USER,
    payload: data
  });
};
