import xlsxwriter
import time
import sys
from guppy import hpy
import random
import string

def randomString(stringLength=6):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(stringLength))

hp = hpy()
print hp.heap()

start_time = time.time()

workbook = xlsxwriter.Workbook('xlsxWriter.xlsx')
worksheet = workbook.add_worksheet()

for col in range (0, 50):
    for row in range(0, 102400):
        worksheet.write(row, col, randomString())

workbook.close()

print("--- %s seconds ---" % (time.time() - start_time))
print hp.heap()