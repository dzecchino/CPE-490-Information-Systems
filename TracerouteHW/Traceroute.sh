i=0
while [ $i -le 287 ]
do
	echo "Traceroute #: $i"
	traceroute www.camera.it >> raw_data
	sleep 5m
	i=$((i+1))
done

#for (( i=0;i<288;i++ ))
#do
#	echo "Hi $i"
#	sleep 3
#done