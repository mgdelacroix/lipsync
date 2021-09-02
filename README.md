# LipSync

LipSync is a small program to generate a RSS feed for a set of videos,
so it can be uploaded to a server and used in a podcast app. The
intent is to be able to listen to video shows as if they were
podcasts, with all the advantages that a podcast apps tend to have
(remembering position, turning off the phone screen, sync, etc) versus
video apps, that tend not to be as comfortable for this kind of
content.

```sh
$ lipsync generate --title "Sample podcast" --link "http://samplepodcast.com/files" --files ./files
```

## Roadmap

- [X] Add configuration file to avoid passing all details as flags.
- [X] Add a -serve option that will start a simple webserver.
- [X] Serve podcast files in a subpath as part of the webserver.
- [ ] Add a homepage in /
