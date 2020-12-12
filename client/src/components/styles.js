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
  fab_user: {
    position: 'fixed',
    top: theme.spacing(2),
    right: theme.spacing(2),
  },
  fab_add: {
    position: 'fixed',
    top: theme.spacing(10),
    right: theme.spacing(2),
  },
  fab_help: {
    position: 'fixed',
    top: theme.spacing(2),
    left: theme.spacing(2),
  },
}));

export default useStyles;
