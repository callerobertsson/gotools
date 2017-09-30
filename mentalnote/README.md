# mentalnote

MentalNote is a simple command line program for entering messages that will be stored in a Slack group. The messages of the group can be retrieved by using the -l option.

### Usage ###

```
$ mentalnote -h
```

MentalNote will show you the help text.

```
$ mentalnote
```

MentalNote will let you enter a multi line message. Input is terminated by Ctrl-D or an . on a single line.

```
$ mentalnote -m "TEXT MESSAGE"
```

The string TEXT MESSAGE will be stored.

```
$ mentalnote -l
```

A listing of (at the moment) all messages will be displayed.

### Set up ###

* Clone this repository
* Create a configuration file, `~/.mn.json`, and enter the following:
    * `api-token`, the token generated by Slack for your team.
	* `channel-id`, the identifier (not name) for the channel to use.
	* `username`, an optional name to display on the posts. Bot is the default.
	* `icon-url`, an optional URL to the avatar icon used for the messages.

The config file can be overridden by setting the environment variable
`MENTALNOTE_CONFIG`.

### Building ###

* Run "go build mn.go" to create the executable file
    * mn for Unix systems
    * mn.exe for Windows

### Running ###

Make sure there is a correct config.json file in the current working directory.

### Binary Versions ###

If you are lucky you might find a suitable binary in the download section.

*Note:* You need a config.json file in the same directory as the executable. Copy the config.json.example and fill in your Slack parameters.

### Future ###

Next up is to limit the content when running in list mode (-l). Now it returns all messages.

Later, searching would be a nice feature.

### Contact ###

Questions are welcome to [calle@upset.se](mailto:calle@upset.se)
