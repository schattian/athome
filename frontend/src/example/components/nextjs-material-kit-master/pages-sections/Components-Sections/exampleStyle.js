import imagesStyle from './imagesStyles';
import { conatinerFluid } from './nextjs-material-kit';


const exampleStyle = {
  section: {
    padding: '70px 0',
  },
  container: {
    ...conatinerFluid,
    textAlign: 'center !important',
  },
  ...imagesStyle,
  link: {
    textDecoration: 'none',
  },
};

export default exampleStyle;
