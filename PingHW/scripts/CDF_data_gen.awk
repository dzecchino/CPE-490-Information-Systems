BEGIN{
	count = 0
	MAX = 86400
}

{
	count++
	cumulative_cdf = count/MAX
	print cumulative_cdf, $1
}

END{
	
}