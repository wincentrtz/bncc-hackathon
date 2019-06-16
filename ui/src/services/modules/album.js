import { axiosInstance } from "../axios.config";
import MockAdapter from "axios-mock-adapter";

import album from "data/album.js";

import config from "configs";

const url = config.urls.album;

export default {
  getFlights() {
    const mock = new MockAdapter(axiosInstance);
    mock.onGet(url.getFlight).reply(200, album.flightData);
    return axiosInstance.get(url.getFlight);
  },
  getAlbums() {
    const mock = new MockAdapter(axiosInstance);
    mock.onGet(url.getAlbums).reply(200, album.albumData);
    return axiosInstance.get(url.getAlbums);
  },
  getAlbumsDetail() {
    const mock = new MockAdapter(axiosInstance);
    mock.onGet(url.getAlbumDetail).reply(200, album.albumDetail);
    return axiosInstance.get(url.getAlbumDetail);
  }
};
