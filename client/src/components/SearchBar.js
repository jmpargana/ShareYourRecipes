import React, {useState, useContext} from 'react';
import {Box,Paper, Chip, TextField} from '@material-ui/core';
import {Autocomplete} from '@material-ui/lab';
import useStyles from './styles';
import { Context } from '../context/store';

export default function SearchBar() {
  const classes = useStyles();
  const [value, setValue] = useState(null)
  const [tags, setTags] = useState([])
  const { state, dispatch } = useContext(Context);

  const fetchRecipes = async () => {
    const url = 'http://localhost:8000/recipes';
    const options = {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json;charset=UTF0-8'
      },
    };

    fetch(url, options)
      .then(response => {
        console.log(response)
        dispatch({ type: "UPLOAD_RECIPES" });
      })
  }
  // fetchRecipes();

  const recipesToTags = () => {
    const tags = []
    state.recipes.map(recipe => tags.push(...recipe.tags))
    const dedupTags = [...new Set(tags)]
    return dedupTags.map((tag, index) => ({key: index, label: tag}))
  }

  const [chipData, setChipData] = useState(recipesToTags)

  // Delete from tags and reappend to list of suggestions
  const handleDeleteTag = (chipToDelete) => () => {
    setTags((chips) => chips.filter((chip) => chip.key !== chipToDelete.key));
  }

  // 1. Add to chips
  // 2. Delete from suggestions
  // 3. Reset value
  const handleAppendTag = tag => {
    if (chipData.some(chip => tag === chip)) {
      // Add to chips
      setTags(tags => [...tags, tag])

      // Delete from available list
      setChipData((chips) => chips.filter((chip) => chip.key !== tag.key));

      // Reset list
      setValue(null)
    }
  }

  return (
    <Box>
      <Autocomplete
        autoHighlight
        value={value}
        onChange={(_, newValue) => handleAppendTag(newValue)}
        options={chipData}
        getOptionLabel={option => option ? option.label : ""}
        renderInput={
          (params) => 
            <TextField {...params} variant="outlined" margin="normal" />
        }
      />
      <Paper elevation={0} component='ul' className={classes.root}>
        {tags.map((tag) => (
          <li key={tag.key}>
            <Chip 
              color="primary"
              label={tag.label}
              onDelete={handleDeleteTag(tag)}
              className={classes.chip}
            />
          </li>
        ))}
      </Paper>
    </Box>
  );
}



