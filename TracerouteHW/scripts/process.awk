BEGIN{
}

{
	for (i = 1; i <= 15; i++) {
		if ($i == "ms") {
			print $(i-1);
		};
	}
}

END{
}
