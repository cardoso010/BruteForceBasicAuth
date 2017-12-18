### Brute Force Basic Auth
Tool used to make brute force on basic auth heavily used in routers.

### Prerequisites
```
GO
git
```

## Getting Started
```
$mkdir -p $GOPATH/src/github.com/cardoso010 && cd $GOPATH/src/github.com/cardoso010
$git clone https://github.com/cardoso010/BruteForceBasicAuth
$go build
$./BruteForceBasicAuth -h=192.168.10.1 -u=username -p=/home/user/wordlist.txt
```
-h -- host where you to do attack
-u -- username
-p -- path where wordlist is

## Authors

* **Gabriel Cardoso Luiz** - *Initial work* - [cardoso010](https://github.com/cardoso010)




