#!/usr/bin/perl

print "Enter a time:\n";
my $time = <STDIN>;
chomp $time;
if ($time =~ /^([0-9]|10|11|12)(:[0-5][0-9])?(am|pm)|([0-9]|[01][0-9]|20|21|22|23):([0-5][0-9])$/) {
	print "Valid time\n";
}else {
	print "Invalid time\n";
}

