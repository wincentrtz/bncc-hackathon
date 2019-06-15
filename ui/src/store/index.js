import { createStore, combineReducers, applyMiddleware } from "redux";
import { reducer as formReducer } from "redux-form";
import thunk from "redux-thunk";

import userReducers from "./reducers/user-reducers";
import albumReducers from "./reducers/album-reducers";

const rootReducer = combineReducers({
  form: formReducer,
  userReducers,
  albumReducers
});

const store = createStore(rootReducer, applyMiddleware(thunk));

export default store;
