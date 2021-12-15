#Step 1: Collect the data
echo "Collecting data..."
#ping -c 86400 www.destination.com >> RAWDATA

#Step 2: Clean up the data, extract just the delay values
awk -f ./scripts/cleanUp.awk raw_data_ping > just_delay
echo "Cleaning up data..."

#Cleaning up DATA to retrieve every Nth line
awk 'NR % 800 == 0' just_delay > DATA_SAMPLE

#Adding a time to column 3
awk -f ./scripts/Adding_Time_Column.awk DATA_SAMPLE > RTTvTime
#Estimated RTT
awk -f ./scripts/EWMA_Estimate.awk RTTvTime > EstimatevTime

#Step 5: Plot the data, clean up, and you're all set!
echo "Detecting all temporary files..."
#rm RAWDATA
rm just_delay
#rm sorted_delay

gnuplot gplot
echo "Plot is ready to go!"
