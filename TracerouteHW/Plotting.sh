awk -f ./scripts/process.awk raw_data > timeValues

sort -n timeValues > sorted_time
rm timeValues

awk -f ./scripts/CDF_data_gen.awk sorted_time > refinedDATA
rm sorted_time

gnuplot gplot
