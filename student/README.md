# guess-it-1

The program takes numbers from standard input and makes a guess what is the interval the next number should be in.
The interval is printed in that way: lower_limit upper_limit.

## For auditor

**Short summary of audit's [README](https://github.com/01-edu/public/tree/master/subjects/guess-it-1/audit) and instruction of testing ** 

- you have to have node. It can be downloaded from their [site](https://nodejs.org/en/download/) or you can use brew to set up it (`brew install node`).
- download [test file](https://assets.01-edu.org/guess-the-number.zip) (link from audit's page);
- unzip it in a folder `guess-the-name` (or any name which you want);
- put the `student` folder to the folder where you have unpacked the test file (`guess-the-name`);
- from this folder (`guess-the-name`):
```console 
cd ai
chmod +x *
cd ..
npm install
node server.js
``` 
- in your brouther open http://localhost:3000/?guesser=big-range and try to test with data1, data2 and data3
- then open http://localhost:3000/?guesser=average and also try to test with data1, data2 and data3
- finally open http://localhost:3000/?guesser=median and also try to test with data1, data2 and data3

- after clicking the data button please wait for numbers to appear and then you can click `Quick` and can see scores;
- it is recommended to click `Clean` button after each test.



## Authors

- Olena Budarahina

