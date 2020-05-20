---
to: '<% if(scss) { %>src/components/<%= h.changeCase.pascal(name) %>/<%= h.changeCase.lcFirst(name) %>.scss%><% } %>'
---
.<%= h.changeCase.lcFirst(name) %>_component {
  padding: 15px;
  border: 1px solid green;
}