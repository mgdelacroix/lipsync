# LipSync

LipSync is a small program to generate a RSS feed for a set of audio
files, so it can be uploaded to a server and used in a podcast
app. The intent is to be able to listen to video shows as if they were
podcasts, with all the advantages that a podcast apps tend to have
(remembering position, turning off the phone screen, sync, etc) versus
video apps, that tend not to be as comfortable for this kind of
content.

```sh
$ lipsync --help
NAME:
   lipsync - create podcast RSS files on the fly!

USAGE:
   lipsync [global options] command [command options] [arguments...]

COMMANDS:
   generate  generates a podcast RSS file
   serve     starts a HTTP server with the podcast information
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)


# copy the sample configuration file
$ cp lipsync-sample.yaml lipsync.yaml

# after tuning the file, generate the RSS file
$ lipsync generate

# or start the webserver
$ lipsync serve
```
