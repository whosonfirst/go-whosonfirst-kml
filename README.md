# go-whosonfirst-kml

Tools for working with KML files

## Install

You will need to have both `Go` and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Tools

### kmltocsv

_There are many KML to CSV tools. This one is ours._

Takes one or more KML files as inputs and outputs a CSV document (to `STDOUT`) with the following headers: `latitude,longitude,name,description`. For example:

```
./bin/kmltocsv ~/Downloads/icecream.kml
latitude,longitude,name,description
37.761845,-122.41195,Tartine Manufactory,"Of course there is much more than ice cream here, but folks are lining up for the buffalo-milk based ice cream. "
37.75773,-122.42097,Xanath Ice Cream,A favorite stop during Sunday Streets. 
37.744183,-122.422775,Mitchell's Ice Cream,"Buko, Buko..coconut and then dip it in chocolate. Yum. "
37.76159,-122.42572,Bi-Rite Creamery,"One person at ML does not like chocolate and can eat a whole pint of the Creamery's Mexican chocolate with salted peanuts. Here is an ode to that very ice cream. https://missionlocal.org/2014/05/perfect-or-not/  But one should get caught up on one flavor.  Justino, a palatero, preferred another. https://missionlocal.org/2012/06/ice-cream-social-at-bi-rite/"
37.758305,-122.42153,Smitten Ice Cream,Now open and a return to the owner's Mission roots. We still remember her pulling her cart around the Mission and serving up freshly churned ice cream. 
37.76505,-122.42222,CREAM,It is the ice cream sandwiches that draw customers. 
37.75877,-122.42045,Garden Creamery,It's vegan ice cream here. 
37.752815,-122.412155,Humphry Slocombe,Lots of odd flavors here - cornflakes and bourbon among them. 
37.765152,-122.42079,Nieves Cinco De Mayo,"Not Strictly ice cream, but refreshing as it is somewhere between a snow cone and a sorbet. https://missionlocal.org/2015/06/nieves-shop-reboots-on-16th-after-fire/"
37.755608,-122.41784,La Copa Loca,A one-man shop. https://missionlocal.org/2012/02/making-artisan-gelato-on-capp-street/
```

## Caveats

* This does not handle anything other than `Points` right now.
* All parsing errors are treated as existential and cause `kmltocsv` to stop processing.
