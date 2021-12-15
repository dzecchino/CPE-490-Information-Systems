BEGIN {
}

{
	if($4 == "ms" && $6 == "ms" && $8 == "ms") {
		print $3;
		print $5;
		print $7;
	}
	if($4 == "ms" && $8 == "ms" && $12 == "ms") {
		print $3;
		print $7;
		print $11;
	}
	if($5 == "ms"){
		print $4;
	}
}

END {	
}
