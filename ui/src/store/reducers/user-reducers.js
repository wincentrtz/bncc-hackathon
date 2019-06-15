import { LOGIN_USER } from "../constant/user-constants";

const userState = {
  user: {}
};

const userReducers = (state = userState, action) => {
  switch (action.type) {
    case LOGIN_USER:
      return {
        ...state,
        user: action.payload
      };
    default:
      return state;
  }
};

export default userReducers;
