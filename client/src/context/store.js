import React, {createContext, useReducer} from 'react';

// const initialState = {};
const initialState = {
  recipes: [
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
}

const Context = createContext(initialState)

const Reducer = (state, action) => {
  switch (action.type) {
    case 'UPLOAD_RECIPES':
      return {
        ...state,
        recipes: [
          ...state.recipes,
          action.recipes
        ]
      };
    case 'UPDATE_RECIPE':
      // TODO:
      return {}
    case 'DELTE_RECIPE':
      // TODO:
      return {}
    default:
      throw new Error();
  };
}

const StateProvider = ( { children } ) => {
  const [state, dispatch] = useReducer(Reducer, initialState)

  // TODO: useEffect with localStorage

  return (
    <Context.Provider value={{state, dispatch}}>
      {children}
    </Context.Provider>
  );
}


export {Context, StateProvider};
