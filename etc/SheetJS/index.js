var XLSX = require('xlsx')
var filename = "SheetJS_r102400xc50.xlsx";
var data = [];

while(data.length < 102400) {
  var col = [];
  for (var i = 50; i > 0; i--) {
    col.push(Math.random().toString(36).substring(6));
  }
  data[data.length] = col;
};

console.time('cost');

var ws_name = "SheetJS";
var wb = XLSX.utils.book_new(), ws = XLSX.utils.aoa_to_sheet(data);
XLSX.utils.book_append_sheet(wb, ws, ws_name);
XLSX.writeFile(wb, filename);

console.timeEnd('cost');
console.log(process.memoryUsage());
