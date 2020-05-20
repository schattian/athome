---
to: '<% if(scss) { %>src/screens/<%= h.changeCase.lcFirst(name) %>/<%= h.changeCase.lcFirst(name) %>.scss%><% } %>'
---
.<%= h.changeCase.lcFirst(name) %>_screen {
  padding: 15px;
  border: 1px solid red;
}