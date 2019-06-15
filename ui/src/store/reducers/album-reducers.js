import { FETCH_ALBUM, FETCH_ALBUM_DETAIL } from "../constant/album-constants";

const albumState = {
  albums: [],
  album: {}
};

const albumReducers = (state = albumState, action) => {
  switch (action.type) {
    case FETCH_ALBUM:
      return {
        ...state,
        albums: action.payload
      };
    case FETCH_ALBUM_DETAIL:
      return {
        ...state,
        album: action.payload
      };
    default:
      return state;
  }
};

export default albumReducers;
