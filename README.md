About
---

Just playing with setting up a development environment using [fig](http://www.fig.sh). Create a simple Golang server that connects to a linked Redis container to increment a counter. 

Use
---

```shell
fig build; fig up

curl -iX GET http://localdocker:8080
```
