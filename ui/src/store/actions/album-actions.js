import { FETCH_ALBUM, FETCH_ALBUM_DETAIL } from "../constant/album-constants";

import api from "services/modules/album";

export const fetchAlbums = () => async dispatch => {
  const { data } = await api.getAlbums();
  console.log(data);
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
