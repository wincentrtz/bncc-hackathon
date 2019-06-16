import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import {
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Divider,
  Typography
} from "@material-ui/core";
import InboxIcon from "@material-ui/icons/MoveToInbox";
import MailIcon from "@material-ui/icons/Mail";
import AccountCircle from "@material-ui/icons/AccountCircle";

const useStyles = makeStyles(theme => ({
  list: {
    width: 250
  }
}));

const Sidelist = ({ side, onToggle }) => {
  const classes = useStyles();
  return (
    <div
      className={classes.list}
      role="presentation"
      onClick={onToggle(side, false)}
      onKeyDown={onToggle(side, false)}
    >
      <List>
        <ListItem
          button
          color="primary"
          style={{ flexDirection: "column", alignItems: "flex-start" }}
        >
          <h1 style={{ marginBottom: "5px" }}>Trippy</h1>
          <h2 style={{ margin: 0, color: "gray" }}>Hi, Wincent</h2>
        </ListItem>
      </List>
      <List>
        {["Notifications"].map((text, index) => (
          <ListItem button key={text}>
            <ListItemIcon>
              {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
            </ListItemIcon>
            <ListItemText primary={text} />
          </ListItem>
        ))}
      </List>
      <Divider />
      <List>
        {["Profile", "Settings", "Logout"].map((text, index) => (
          <ListItem button key={text}>
            <ListItemIcon>
              {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
            </ListItemIcon>
            <ListItemText primary={text} />
          </ListItem>
        ))}
      </List>
    </div>
  );
};

export default Sidelist;
