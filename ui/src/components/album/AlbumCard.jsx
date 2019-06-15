import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Typography from "@material-ui/core/Typography";
import Chip from "@material-ui/core/Chip";
import { Grid } from "@material-ui/core";

const useStyles = makeStyles({
  card: {
    marginBottom: "20px",
    marginRight: "20px"
  },
  cardTitle: {
    fontWeight: "bolder"
  },
  media: {
    height: 140
  },
  chip: {
    fontWeight: "bolder",
    color: "white"
  }
});

const renderColor = flag => (flag ? "#2196f3" : "#e10050");
const renderChipLabel = flag => (flag ? "Paid" : "Unpaid");

const AlbumCard = ({ album }) => {
  const classes = useStyles();
  return (
    <Card className={classes.card}>
      <CardMedia
        className={classes.media}
        image="assets/images/login-background.jpg"
      />
      <CardContent>
        <Grid container alignItems="center">
          <Grid item xs={10}>
            <Typography
              color="primary"
              className={classes.cardTitle}
              gutterBottom
              variant="h5"
              component="h2"
            >
              {album.name}
            </Typography>
            <Typography variant="body2" color="textSecondary" component="p">
              By Author
            </Typography>
          </Grid>
          <Grid item xs={2}>
            <Chip
              label={renderChipLabel(album.IsPaidOff)}
              className={classes.chip}
              style={{ backgroundColor: renderColor(album.IsPaidOff) }}
            />
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
};

export default AlbumCard;
