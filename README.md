# env2consul

This is a tool to put environment in given dotenv file in consul key/value store. To do so, run:

```sh
$ env2consul -prefix foo/bar .env
```

Where *.env* is the dotenv file to load and *foo/bar* is the prefix for keys. You may pass more than one file to load.

Consul command line tool must have been installed in your *PATH*. You can download it from <https://www.consul.io/downloads>.

*Enjoy!*
