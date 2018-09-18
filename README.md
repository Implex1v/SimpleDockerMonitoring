# SimpleDockerMonitoring
Using simple-dm you can simply monitor your running docker containers. You can run `simple-dm` as a cli tool and configure
it via a simple yml file. You simply have to define of docker container names which are expected to run. If these docker
containers are not running a e-mail will be send to a specified e-mail address with further information. To automate the
monitoring simply create a cron job and run the `simple-dm` command.

## Install & Usage
1) Clone this repo, cd into the directory and run `go install`
2) Copy the `config.yml.dist` to `config.yml` and modify it to your will
3) Run `simple-dm` and you are good to go

You can specify a custom configuration file by using the `-config` argument e.g. `simple-dm -config path/to/config.yml`.

## Configuration
```yaml
enable: true                          # enables SimpleDockerMonitoring
email:
  enable: true                        # enables e-mail support
  username: "example@gmail.com"       # your username 
  passwordenv: "SDM_MAIL_PASSWORD"    # use the env var instead of email.password
  password: "p@ssw0rd"                # your password
  hostname : "smtp.gmail.com"         # the hostname of the smtp server
  url: "smtp.gmail.com:587"           # the url to the smtp server with port
  sender: "example@gmail.com"         # the sender of the e-mail
  recipient: "example@gmail.com"      # the recipient of the e-mail

containers:                           # list of containers to monitor
  - mysql
  - nginx
```
An example configuration is provided above. 

If you want to use Gmail as your e-mail provider you need to enable
support for "[less secure apps](https://support.google.com/accounts/answer/6010255)". **Hint**: you might create a 
new Gmail account and use it only for _SimpleDockerMonitoring_ or use another e-mail provider.

For security aspects you can specify an environment variable which contains the e-mail password. The name of the 
environment variable can be defined in `email.passwordenv`. If `email.passwordenv` is not null (`""`) the specified 
environment variable will be used.

## Other
* This is my first go project and I am thankful for improvements :)
* On error `simple-dm` will write an error message to the console and exit with a unique code > 0.
* `simple-dm` might need sudo privileges to communicate with the docker host