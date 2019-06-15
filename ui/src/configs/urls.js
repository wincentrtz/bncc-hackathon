export default {
  user: {
    login: "/login"
  },
  album: {
    getAlbums: "/albums",
    getAlbumsDetail: id => "/albums/" + id
  }
};
