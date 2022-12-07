#!/usr/bin/perl

use strict;

my $lineno = 1;
while (<>) {
	print "$lineno) ";
	if (/Perl is (so,? ?)+ cool!/) {
		print;
	} else {
		print "\n";
	}
	$lineno += 1;
}
