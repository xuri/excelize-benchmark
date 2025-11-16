import xlsxwriter
import time
import psutil
import random
import string

def randomString(stringLength=6):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(stringLength))

p = psutil.Process()
print(p.memory_info())

start_time = time.time()

workbook = xlsxwriter.Workbook('xlsxWriter.xlsx', options={"constant_memory": True})
worksheet = workbook.add_worksheet()

for col in range (0, 50):
    for row in range(0, 102400):
        worksheet.write(row, col, randomString())

workbook.close()

print("--- %s seconds ---" % (time.time() - start_time))
print(p.memory_info())
