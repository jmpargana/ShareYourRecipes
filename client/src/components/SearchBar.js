import React, {useState} from 'react';
import {Box,Paper, Chip, TextField} from '@material-ui/core';
import {Autocomplete} from '@material-ui/lab';
import {makeStyles} from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'flex-end',
    flexWrap: 'wrap',
    listStyle: 'none',
    // padding: theme.spacing(0.5),
    margin: 0,
  },
  chip: {
    margin: theme.spacing(0.5),
  },
}));

export default function SearchBar() {
  const classes = useStyles();
  const [value, setValue] = useState(null)
  const [chipData, setChipData] = useState([
    { key: 0, label: 'Angular' },
    { key: 1, label: 'React' },
    { key: 2, label: 'Vue' },
    { key: 3, label: 'Long Chip Name' },
    { key: 4, label: 'Another Name' },
  ])

  const appendTag = tag => setChipData(chipData => [...chipData, tag]);

  const defaultProps = {
    options: recipes,
    getOptionLabel: (option) => option.title,
  }

  const handleDeleteTag = (chipToDelete) => () => {
    setChipData((chips) => chips.filter((chip) => chip.key !== chipToDelete.key));
  }

  const handleAppendTag = tag => {
    if (chipData.some(chip => tag === chip.label.toLowercase)) {
      appendTag(tag)
      setValue('')
    }
  }

  return (
    <Box>
      <Autocomplete
        {...defaultProps}
        autoHighlight
        value={value}
        onChange={(_, newValue) => handleAppendTag(newValue)}
        // onChange={(_, newValue) => setValue(newValue)}
        renderInput={
          (params) => 
            <TextField {...params} variant="outlined" margin="normal" />
        }
      />
      <Paper elevation={0} component='ul' className={classes.root}>
        {chipData.map((data) => (
          <li key={data.key}>
            <Chip 
              color="primary"
              label={data.label}
              onDelete={handleDeleteTag(data)}
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
]
