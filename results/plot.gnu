set datafile separator ","
set terminal svg size 728, 600 fname 'Helvetica Neue, Helvetica, Segoe UI, Arial, freesans, sans-serif'

set key left top

# define axis
# remove border on top and right and set color to gray
set style line 11 lc rgb '#808080' lt 1
set border 3 back ls 11
set tics nomirror

# define grid
set style line 12 lc rgb '#808080' lt 0 lw 1
set grid back ls 12

# Use logarithmic scale
set logscale xy

set xlabel 'Number of elements'

set ylabel 'Average insert time (ns)'
set output 'insert_times.svg'
plot for [col=2:6] 'insert_times.csv' using 1:col with lines title columnhead

set ylabel 'Average lookup time (ns)'
set output 'lookup_times.svg'
plot for [col=2:6] 'lookup_times.csv' using 1:col with lines title columnhead

set ylabel 'Average memory usage (Bytes)'
unset logscale y
set output 'memory_usage.svg'
plot for [col=2:6] 'memory_usage.csv' using 1:col with lines title columnhead