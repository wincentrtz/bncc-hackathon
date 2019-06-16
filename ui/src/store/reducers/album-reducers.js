import {
  FETCH_ALBUM,
  FETCH_ALBUM_DETAIL,
  FETCH_FLIGHTS
} from "../constant/album-constants";

const albumState = {
  albums: [],
  album: {},
  flights: []
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
    case FETCH_FLIGHTS:
      return {
        ...state,
        flights: action.payload
      };
    default:
      return state;
  }
};

export default albumReducers;
