About
---

Just playing with setting up a development environment using [fig](http://www.fig.sh). Create a simple Golang server that connects to a linked Redis container to increment a counter.

Use
---

```shell
fig build; fig up

curl -iX GET http://localdocker:8080
```

Docker Setup
---

* Install [Virtualbox](https://www.virtualbox.org/wiki/Downloads)
* Install XCode command line tools
	* Open a terminal and type `gcc`, you'll be prompted to install the necessary toolset
* Install Homebrew
	* `ruby -e "$(curl -fsSL https://raw.github.com/Homebrew/homebrew/go/install)"`
* Install boot2docker
	* brew install boot2docker
* Enable `localdocker` DNS
	* `echo "$(boot2docker ip 2> /dev/null) localdocker" | sudo tee -a /etc/hosts
192.168.59.103 localdocker`