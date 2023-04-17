# Description: #
This project receives a string as an argument and outputs the string in a graphic representation using ASCII. And prints the output via three banners: standard, shadow and thinkertoy. Also, the user can choose the color of the output text. The user can save the output text in a txt file.

# Usage: # 
-Run the program by typing "go run cmd/main.go" in the terminal

-Ctrl+CLICK on the link in the terminal or open the web browser of your choice and go to localhost:8080/


-Type the string inside the textbox and choose the banner and/or color of the output


-Click the Submit button to see the output on the website /OR Click the Download button to download the output in the data.txt file on your computer

# Implementation Details: #
-The input string is passed to the printWord function, which adds each letter(in ascii) to the output string and returns it to splitWord function
