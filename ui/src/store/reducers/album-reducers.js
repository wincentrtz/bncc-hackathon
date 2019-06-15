import { FETCH_ALBUM } from "../constant/album-constants";

const albumState = {
  albums: [1, 2, 3, 4, 5]
};

const albumReducers = (state = albumState, action) => {
  switch (action.type) {
    case FETCH_ALBUM:
      return {
        ...state,
        albums: action.payload
      };
    default:
      return state;
  }
};

export default albumReducers;
