
export default (theme) => ({
  defaultFontStyle: {
  },
  defaultHeaderMargins: {
    marginTop: '20px',
    marginBottom: '10px',
  },
  quote: {
    padding: '10px 20px',
    margin: '0 0 20px',
    fontSize: '17.5px',
    borderLeft: `5px solid ${theme.palette.grey[50]}`,
  },
  quoteText: {
    margin: '0 0 10px',
    fontStyle: 'italic',
  },
  quoteAuthor: {
    display: 'block',
    fontSize: '80%',
    lineHeight: '1.42857143',
    color: theme.palette.grey.A700,
  },
  mutedText: {
    color: theme.palette.grey[50],
  },
  primaryText: {
    color: theme.palette.primary.main,
  },
  infoText: {
    color: theme.palette.info.main,
  },
  successText: {
    color: theme.palette.success.main,
  },
  warningText: {
    color: theme.palette.warning.main,
  },
  dangerText: {
    color: theme.palette.error.main,
  },
});
