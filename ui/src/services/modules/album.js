import { axiosInstance } from "../axios.config";
import MockAdapter from "axios-mock-adapter";

import album from "data/album.js";

import config from "configs";

const url = config.urls.album;

export default {
  getFlights() {
    return axiosInstance.get(url.getFlight);
  },
  getAlbums() {
    return axiosInstance.get(url.getAlbums);
  },
  getAlbumsDetail() {
    return axiosInstance.get(url.getAlbumDetail);
  }
};
