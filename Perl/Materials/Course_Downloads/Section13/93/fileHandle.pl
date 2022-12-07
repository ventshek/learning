#!/usr/bin/perl

use strict;

my $i;

$i++ if (-e "testfile$i");
open FILE, ">testfile$i" or die "Could not open testfile $!";

print FILE "This is a test.\n";
