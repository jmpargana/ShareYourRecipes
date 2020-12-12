import React from 'react';
import {Box, Text} from 'grommet';
import {Paper, Chip} from '@material-ui/core'
import useStyles from './styles';

export default function RecipeBar({recipe}) {
  const classes = useStyles()

  return (
    <Box>
      <Text>{recipe.title} </Text>
      <Text>{recipe.time} </Text>
      <Paper elevation={0} component='ul' className={classes.root}>
        {recipe.tags.map((tag, index) => (
          <li key={index}>
            <Chip 
              color="primary"
              label={tag}
              className={classes.chip}
            />
          </li>
        ))}
      </Paper>
    </Box>
  );
}
