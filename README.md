Google App Engine Reverse Http Proxy
====================================

A very basic http reverse proxy for Google App Engine. It proxies all http
requests to whatever upstream URL you specified in your app.yaml. We use it to
circumvent Internet censorship in Russia.

How to deploy
-------------

If you don't have google-cloud-sdk w/ app-engine-go (these might be ArchLinux-specific):

* Get google-cloud-sdk. Install at home, not via package manager :/
* Do $WHATEVER/bin/gcloud init
* Do $WHATEVER/bin/gcloud components install app-engine-go

Then, here:

* Replace UPSTREAM env var in app.yaml with your upstream URL,
* $WHATEVER/bin/gcloud app deploy --project YOUR-GAE-PROJECT
* if it works on resulting url, do
  $WHATEVER/bin/gcloud app deploy --project YOUR-GAE-PROJECT --promote
