import os
user_input = input("Enter a command: ")
os.system(user_input)  # Vulnerable to command injection
eval("print('Hello, ' + user_input)")  # Vulnerable to arbitrary code execution
# sample changes
