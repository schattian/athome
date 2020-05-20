---
to: '<% if(type === "Component") { %>src/components/<%= h.changeCase.pascal(name) %>/<%= h.changeCase.pascal(name) %>.js%><% } %>'
---
<% const comp = h.changeCase.pascal(name) -%>
<% const compClassName = h.changeCase.lcFirst(name) -%>
import React from 'react';
<% if(locals.withStyle) { -%>//<%}-%>import PropTypes from 'prop-types';
<% if(locals.withStyle) { -%>
import { withStyles, createStyles } from '@material-ui/core';
<%}-%>
<% if(locals.scss) { -%>
import './<%= h.changeCase.lcFirst(name) %>.scss';
<%}-%>
<% if(locals.redux) { -%>
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
<%}-%>
import { withTranslation } from 'src/i18n';

class <%=comp %> extends React.Component() {
  constructor(props) {
    super(props);
    this.state = {
      myName: '<%=comp %>',
    };
  }

  render() {
    const { myName } = this.state;
    const { <% if(locals.scss) { -%> classes, <%}-%> t } = this.props;
    return (
      <div <% if(locals.withStyle) { -%>className={`${classes.root} <%=compClassName %>_component`}<%}-%><% if(!locals.withStyle) { -%>className="<%=compClassName %>_component"<%}-%>>
        <h2>{`I am a ${myName} Component`}</h2>
        <p>{t('add_your_component_translates_to_common.json')}</p>
      </div>
    );
  }
}

const Extended = withTranslation('common')(<%=comp %>);

<% if(locals.withStyle) { -%>
const styles = () => createStyles({
  root: {
    backgroundColor: 'red',
  },
});
<%}-%>
<% if(locals.redux) { -%>
function mapStateToProps(store) {
  return {
    store,
  };
}
function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators({ }, dispatch)
  };
}
export default <% if(locals.withStyle) { -%>withStyles(styles, <%}-%>connect(mapStateToProps, mapDispatchToProps)(Extended)<% if(locals.withStyle) { -%>)<%}-%>;
<%}-%>
<% if(!locals.redux) { -%>
<% if(locals.withStyle) { -%>
export default withStyles(styles)(Extended);
<%}-%>
<% if(!locals.withStyle) { -%>
export default Extended;
<%}-%>
<%}-%>
<%=comp %>.propTypes = {
<% if(locals.withStyle) { -%>
  // eslint-disable-next-line
  classes: PropTypes.object,
<%}
-%>
};
