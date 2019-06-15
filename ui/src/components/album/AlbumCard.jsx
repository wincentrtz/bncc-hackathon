import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles({
  card: {
    maxWidth: 400,
    marginBottom: "20px",
    margin: "0 auto"
  },
  media: {
    height: 140
  }
});

const AlbumCard = () => {
  const classes = useStyles();

  return (
    <Card className={classes.card}>
      <CardMedia
        className={classes.media}
        image="assets/images/login-background.jpg"
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="h2">
          Bandung - Jakarta
        </Typography>
        <Typography variant="body2" color="textSecondary" component="p">
          By Author
        </Typography>
      </CardContent>
    </Card>
  );
};

export default AlbumCard;
