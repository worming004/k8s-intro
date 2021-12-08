import os as os

for x in os.walk('.'):
    if './.' not in x[0]:
        print(x)