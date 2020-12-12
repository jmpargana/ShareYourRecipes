import React from 'react';
import {Box, Text} from 'grommet';
import {Paper, Chip} from '@material-ui/core'
import useStyles from './styles';

export default function RecipeBar({recipe}) {
  const classes = useStyles()

  return (
    <Box>
      <Box direction="row">
        <Text size="large">{recipe.title} </Text>
        <Box flex='grow' />
        <Text color="blue">{recipe.time} </Text>
      </Box>
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
