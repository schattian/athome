---
to: '<% if(type === "WithHook") { %>src/components/<%= h.changeCase.pascal(name) %>/<%= h.changeCase.pascal(name) %>.js%><% } %>'
---
<% const comp = h.changeCase.pascal(name) -%>
<% const compClassName = h.changeCase.lcFirst(name) -%>
import React, { useState } from 'react';
<% if(!locals.withStyle) { -%>//<%}-%>import PropTypes from 'prop-types';
<% if(locals.withStyle) { -%>
import { makeStyles, createStyles } from '@material-ui/core';
<%}-%>
<% if(locals.scss) { -%>
import './<%= h.changeCase.lcFirst(name) %>.scss';
<%}-%>
<% if(locals.redux) { -%>
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
<%}-%>
import { useTranslation } from 'src/i18n'

<% if(locals.withStyle) { -%>
const useStyles = makeStyles(() => createStyles({
  root: {
    backgroundColor: 'red',
  },
}));

<%}-%>
function <%=comp %>() {
  const [count, setCount] = useState(0);
  const { t } = useTranslation('common');
<% if(locals.withStyle) { -%>
  const classes = useStyles();
<%}-%>
  return (
    <div <% if(locals.withStyle) { -%>className={`${classes.root} <%=compClassName %>_component`}<%}-%><% if(!locals.withStyle) { -%>className="<%=compClassName %>_component"<%}-%>>
      <h2>I am a <%=comp %> Hook Component</h2>
      <p>{t('add_your_component_translates_to_common.json')}</p>
      <p>You clicked {count} times</p>
      <button type="button" onClick={() => setCount(count + 1)}>
        Click me
      </button>
    </div>
  );
}

<% if(locals.redux) { -%>
function mapStateToProps(store) {
  return {
    store,
  };
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ }, dispatch),
  };
}
export default connect(mapStateToProps, mapDispatchToProps)(<%=comp %>);
<%}-%>
<% if(!locals.redux) { -%>
<% if(locals.withStyle) { -%>
export default (<%=comp %>);
<%}-%>
<% if(!locals.withStyle) { -%>
export default <%=comp %>;
<%}-%>
<%}-%>
<%=comp %>.propTypes = {
};
