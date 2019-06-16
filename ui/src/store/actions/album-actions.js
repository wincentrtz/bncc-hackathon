import {
  FETCH_ALBUM,
  FETCH_ALBUM_DETAIL,
  FETCH_FLIGHTS
} from "../constant/album-constants";

import api from "services/modules/album";

export const fetchAlbums = () => async dispatch => {
  const { data } = await api.getAlbums();
  dispatch({
    type: FETCH_ALBUM,
    payload: data
  });
};

export const fetchAlbumDetail = () => async dispatch => {
  const { data } = await api.getAlbumsDetail();
  dispatch({
    type: FETCH_ALBUM_DETAIL,
    payload: data
  });
};

export const fetchFlights = () => async dispatch => {
  const { data } = await api.getFlights();
  dispatch({
    type: FETCH_FLIGHTS,
    payload: data
  });
};
