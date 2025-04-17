# orgn

Simple but flexible Origin Server for your CND-fronted web sites.

(Work in progress.  Currently low priority but shooting for May 2025.)

## Motivation

I host some web sites that use Cloudflare as a CDN/caching layer.

I am not happy with the alternatives for simple origin hosting on a VPS.

So, here we have a new solution: `orgn`.

## Goals

* Relatively lightweight binary (subject to normal Golang bloat).
* Easy to deploy static sites, e.g. those built with Zola.
* Easy to deploy sites built around `html/template`.
* Easy to build custom binaries with just the new routes defined.
* Easy to update templates and content from e.g. a GitHub pipeline.
* Useful defaults for some templates, so you can kickstart with just `*.md`.
* Front-matter equivalent for Markdown, available in the templates.
	* Probably use this: github.com/adrg/frontmatter
* XML/Atom support out of the box.
* Sensible config defaults.

## Non-Goals

* High performance (you're using a CDN!).

## Maybe Later

* Support for multiple sites in one runner (not my use-case but useful).
* Public-facing, TLS support (meh, use the CDN).
* Simple database support (on which you could hang more complicated...).
* What else?


