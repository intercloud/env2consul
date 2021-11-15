# env2consul

This is a tool to put environment in given dotenv file in consul key/value store. To do so, run:

```sh
$ env2consul -prefix foo/bar .env
```

Where *.env* is the dotenv file to load and *foo/bar* is the prefix for keys. You may pass more than one file to load.

*Enjoy!*
