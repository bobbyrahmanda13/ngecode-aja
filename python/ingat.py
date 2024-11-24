import sys
import time
# from colorama import init, Fore

# init()

text = "kamu terlalu berharga untuk disentuh tanpa akad,\njadi jangan mau yaa diajak pacaran, pegang-pegangan, peluk-pelukan.\nkarena kamu itu berharga"

def scrolling_text(text):
    for char in text:
        sys.stdout.write(char)
        sys.stdout.flush()
        time.sleep(0.05)
    print()

scrolling_text(text)

# from colorama import init, Fore
# init()
# print(Fore.BLUE + "Hello, world!")
