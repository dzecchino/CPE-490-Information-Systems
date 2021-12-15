BEGIN{
	count = 0
	
}

{
	count++
	print count, $1
}

END{
	
}
