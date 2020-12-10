import React, {useState, useEffect} from 'react';
import {Box,Paper, Chip, TextField} from '@material-ui/core';
import {Autocomplete} from '@material-ui/lab';
import {makeStyles} from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'flex-end',
    flexWrap: 'wrap',
    listStyle: 'none',
    padding: theme.spacing(0.5),
    margin: 0,
  },
  chip: {
    margin: theme.spacing(0.5),
  },
}));

export default function SearchBar() {
  const classes = useStyles();
  const [value, setValue] = useState(null)
  const [tags, setTags] = useState([{key: -1, label: "all"}])

  const recipesToTags = () => {
    const tags = []
    recipes.map(recipe => tags.push(...recipe.tags))
    const dedupTags = [...new Set(tags)]
    return dedupTags.map((tag, index) => ({key: index, label: tag}))
  }

  const [chipData, setChipData] = useState(recipesToTags)

  // useEffect(() => {
  // }, [chipData, tags])


  // Delete from tags and reappend to list of suggestions
  const handleDeleteTag = (chipToDelete) => () => {
    setChipData((chips) => chips.filter((chip) => chip.key !== chipToDelete.key));
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
      setValue('')
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



const recipes = [
  {
    id: 1,
    userId: 2,
    title: "Delicious Vegan Korma",
    private: true,
    ingridients: ["rice", "tofu", "cream", "tomatoes"],
    time: 30,
    method: "1. Prepare all ingridients...",
    tags: ['indian', 'curry', 'korma', 'tofu'],
  },
  {
    id: 2,
    userId: 2,
    title: "Chop Suoey",
    private: true,
    ingridients: ["rice", "tofu", "noodles", "Soy"],
    time: 30,
    method: "1. Prepare all ingridients...",
    tags: ['chinese', 'spicy', 'noodles', 'tofu'],
  },
  {
    id: 3,
    userId: 1,
    title: "Pasta Aglio al Olio",
    private: true,
    ingridients: ["olive oil", "pasta", "garlic", "tomatoes"],
    time: 30,
    method: "1. Prepare all ingridients...",
    tags: ['italian', 'pasta', 'vegan'],
  },
];

