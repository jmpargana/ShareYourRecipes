import React from 'react';
import {Layer} from 'grommet';

export default function CreateRecipe({toggler}) {
  return (
    <Layer
      onClickOutside={() => toggler(false)}
      onEsc={() => toggler(false)}
    >
      <h1>Hey!</h1>
    </Layer>
  );
}
