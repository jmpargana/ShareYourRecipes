import React, {useContext} from 'react';
import {Box, Divider} from '@material-ui/core';
import { Context } from '../context/store';
import RecipeBar from './RecipeBar';

export default function Recipes() {
  const { state } = useContext(Context);

  return (
    <Box>
      {state.recipes.map((recipe, index) => (
        <React.Fragment key={index}>
          <RecipeBar 
            key={index} 
            recipe={recipe} 
          />
           <Divider />
        </React.Fragment>
      ))}
    </Box>
  );
}
