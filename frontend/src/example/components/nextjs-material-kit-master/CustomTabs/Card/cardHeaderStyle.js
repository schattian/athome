
const primaryBoxShadow = {
  boxShadow:
    '0 12px 20px -10px rgba(156, 39, 176, 0.28), 0 4px 20px 0px rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba(156, 39, 176, 0.2)',
};
const infoBoxShadow = {
  boxShadow:
    '0 12px 20px -10px rgba(0, 188, 212, 0.28), 0 4px 20px 0px rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba(0, 188, 212, 0.2)',
};
const successBoxShadow = {
  boxShadow:
    '0 12px 20px -10px rgba(76, 175, 80, 0.28), 0 4px 20px 0px rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba(76, 175, 80, 0.2)',
};
const warningBoxShadow = {
  boxShadow:
    '0 12px 20px -10px rgba(255, 152, 0, 0.28), 0 4px 20px 0px rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba(255, 152, 0, 0.2)',
};
const dangerBoxShadow = {
  boxShadow:
    '0 12px 20px -10px rgba(244, 67, 54, 0.28), 0 4px 20px 0px rgba(0, 0, 0, 0.12), 0 7px 8px -5px rgba(244, 67, 54, 0.2)',
};

const warningCardHeader = {
  color: '#fff',
  background: 'linear-gradient(60deg, #ffa726, #fb8c00)',
  ...warningBoxShadow,
};
const successCardHeader = {
  color: '#fff',
  background: 'linear-gradient(60deg, #66bb6a, #43a047)',
  ...successBoxShadow,
};
const dangerCardHeader = {
  color: '#fff',
  background: 'linear-gradient(60deg, #ef5350, #e53935)',
  ...dangerBoxShadow,
};
const infoCardHeader = {
  color: '#fff',
  background: 'linear-gradient(60deg, #26c6da, #00acc1)',
  ...infoBoxShadow,
};
const primaryCardHeader = {
  color: '#fff',
  background: 'linear-gradient(60deg, #ab47bc, #8e24aa)',
  ...primaryBoxShadow,
};

const cardHeaderStyle = {
  cardHeader: {
    borderRadius: '3px',
    padding: '1rem 15px',
    marginLeft: '15px',
    marginRight: '15px',
    marginTop: '-30px',
    border: '0',
    marginBottom: '0',
  },
  cardHeaderPlain: {
    marginLeft: '0px',
    marginRight: '0px',
  },
  warningCardHeader,
  successCardHeader,
  dangerCardHeader,
  infoCardHeader,
  primaryCardHeader,
};

export default cardHeaderStyle;
