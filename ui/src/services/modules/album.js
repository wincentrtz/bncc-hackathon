import { axiosInstance } from "../axios.config";
import MockAdapter from "axios-mock-adapter";

import album from "data/album.js";

import config from "configs";

const url = config.urls.album;

var mock = new MockAdapter(axiosInstance);

export default {
  getAlbums() {
    mock.onGet(url.getAlbums).reply(200, album.albumData);
    return axiosInstance.get(url.getAlbums);
  },
  getAlbumsDetail() {
    mock.onGet(url.getAlbumDetail).reply(200, album.albumDetail);
    return axiosInstance.get(url.getAlbumDetail);
  }
};
