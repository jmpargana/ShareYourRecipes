import React from 'react';
import {Box} from 'grommet';
import SearchBar from "./SearchBar";
import Recipes from "./Recipes";
import Title from "./Title";

export default function Home() {
  return (
    <Box>
      <Title />
      <SearchBar />
      <Recipes />
    </Box>
  );
}
