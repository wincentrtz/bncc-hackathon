import { axiosInstance } from "../axios.config";

import config from "configs";

const url = config.urls.album;

export default {
  getAlbums() {
    return axiosInstance.get(url.getAlbums);
  }
};
