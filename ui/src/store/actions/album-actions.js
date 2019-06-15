import { FETCH_ALBUM } from "../constant/album-constants";

import api from "services/modules/album";

export const fetchAlbums = payload => async dispatch => {
  const { data } = await api.getAlbums(payload);
  dispatch({
    type: FETCH_ALBUM,
    payload: data
  });
};
