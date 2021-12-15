#Step 1: Collect the data
echo "Collecting data..."
ping -c 86400 www.destination.com >> RAWDATA

#Step 2: Clean up the data, extract just the delay values
awk -f ./scripts/cleanUp.awk RAWDATA > just_delay
echo "Cleaning up data..."

#Step 3: Sort the values
sort -n just_delay > sorted_delay

#Step 4: Get data ready for CDF plot
awk -f ./scripts/CDF_data_gen.awk sorted_delay > DATA

#Step 5: Plot the data, clean up, and you're all set!
echo "Detecting all temporary files..."
rm RAWDATA
rm just_delay
rm sorted_delay

gnuplot gplot
echo "Plot is ready to go!"
