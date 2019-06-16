import { axiosInstance } from "../axios.config";

import config from "configs";

const url = config.urls.user;

export default {
  login(payload) {
    console.log(payload)
    return axiosInstance.post(url.login, payload);
  }
};
