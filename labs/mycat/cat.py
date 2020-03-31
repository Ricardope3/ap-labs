import os
import time
import matplotlib.pyplot as plt
time_arr = []
start_time = time.time()
os.system('cat test.txt > /dev/null')
end_time = time.time()
time_arr.append((end_time-start_time)*1000)
print("cat regular se tardo {} milisegundos").format(time_arr[0])



start_time = time.time()
os.system('./cat2 test.txt > /dev/null')
end_time = time.time()
time_arr.append((end_time-start_time)*1000)
print("cat2 se tardo {} milisegundos").format(time_arr[1])


plt.plot(time_arr, '*', color=[1, 0.5,1])
plt.show()