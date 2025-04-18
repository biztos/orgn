# The Plan

## First get static-site service working.

- Use a Zola build as an example, have two live sites already.
- Very basic config, no templates yet.
- Live reload?  Maybe but don't do a watcher, that's crappo.

Config is maybe:

```toml
static_dir = "/var/mysite/pub"
update_path = "/update_XYZ"
update_token = "SOME_TOKEN_FOR_GITHUB_TO_USE"
```

`orgn init static /var/mysite` should do this for us.

Then update catches a POST with the token for auth and an URL for getting a
.tgz file which is then unpacked with no external binaries, and copied over
pub.

Standard-ish behavior like: don't list dirs unless told to; do look for
index.html unless overridden; etc.  Watch for scope creep!

(Uh-oh, guess we do have some need of optional templates if we are going to
allow directory listings!  Also for error pages.  Also config to use error
pages from the pub directory, easiest that way.)

## Second do a dumb Markdown server with decent defaults.

Always a good check on assumptions around URLs and so on.  Can use this for
something like AICS until I get around to doing a design.

Also gonna need the feed logic here.  And some site info in the markup.

## Third do something with proper templates, for a new site.

TBD but this is kinda the core thing I want to support here, the others are
just for transitions, stopgaps, other people's stuff.

## Fourth add deeper customization features with examples.

A lot of them from GHD in terms of app customization.

