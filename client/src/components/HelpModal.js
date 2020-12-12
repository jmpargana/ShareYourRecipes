import React from 'react';
import {Layer} from 'grommet';

export default function HelpModal({toggler}) {
  return (
    <Layer
      onClickOutside={() => toggler(false)}
      onEsc={() => toggler(false)}
    >
      <h1>Help</h1>
    </Layer>
  );
}
